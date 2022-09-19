// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

// OS memory management abstraction layer
// OS 内存管理抽象层
//
// Regions of the address space managed by the runtime may be in one of four
// states at any given time:
// 1) None - Unreserved and unmapped, the default state of any region.
// 2) Reserved - Owned by the runtime, but accessing it would cause a fault.
//               Does not count against the process' memory footprint.
// 3) Prepared - Reserved, intended not to be backed by physical memory (though
//               an OS may implement this lazily). Can transition efficiently to
//               Ready. Accessing memory in such a region is undefined (may
//               fault, may give back unexpected zeroes, etc.).
// 4) Ready - may be accessed safely.
//
// 1) None - 内存没有被保留或者映射，是地址空间的默认状态.
// 2) Reserved - 运行时持有该地址空间，但是访问该内存会导致错误.
// 3) Prepared - 内存被保留，一般没有对应的物理内存访问该片内存的行为是未定义的可以快速转换到 Ready 状态.
// 4) Ready - 可以被安全访问.
//
// This set of states is more than is strictly necessary（必要） to support all the
// currently（目前） supported platforms. One could get by with just None, Reserved, and
// Ready. However, the Prepared state gives us flexibility（灵活） for performance（执行）
// purposes（作用）. For example, on POSIX-y operating systems, Reserved is usually a
// private anonymous mmap'd region with PROT_NONE set, and to transition
// to Ready would require setting PROT_READ|PROT_WRITE. However the
// underspecification of Prepared lets us use just MADV_FREE to transition from
// Ready to Prepared. Thus with the Prepared state we can set the permission
// bits just once early on, we can efficiently tell the OS that it's free to
// take pages away from us when we don't strictly need them.
//
// This file defines a cross-OS interface for a common set of helpers
// that transition memory regions between these states. The helpers call into
// OS-specific implementations that handle errors, while the interface boundary
// implements cross-OS functionality, like updating runtime accounting.

// sysAlloc transitions an OS-chosen region of memory from None to Ready.
// More specifically, it obtains a large chunk of zeroed memory from the
// operating system, typically on the order of a hundred kilobytes
// or a megabyte. This memory is always immediately available for use.
//
// sysStat must be non-nil.
//
// Don't split the stack as this function may be invoked without a valid G,
// which prevents us from allocating more stack.
// 会从操作系统中获取一大块可用的内存空间，可能为几百 KB 或者几 MB；
//
//go:nosplit
func sysAlloc(n uintptr, sysStat *sysMemStat) unsafe.Pointer {
	sysStat.add(int64(n))
	gcController.mappedReady.Add(int64(n))
	return sysAllocOS(n)
}

// sysUnused transitions a memory region from Ready to Prepared. It notifies the
// operating system that the physical pages backing this memory region are no
// longer needed and can be reused for other purposes. The contents of a
// sysUnused memory region are considered forfeit and the region must not be
// accessed again until sysUsed is called.
//
// 通知操作系统虚拟内存对应的物理内存已经不再需要，可以重用物理内存；
func sysUnused(v unsafe.Pointer, n uintptr) {
	gcController.mappedReady.Add(-int64(n))
	sysUnusedOS(v, n)
}

// sysUsed transitions a memory region from Prepared to Ready. It notifies the
// operating system that the memory region is needed and ensures that the region
// may be safely accessed. This is typically a no-op on systems that don't have
// an explicit commit step and hard over-commit limits, but is critical on
// Windows, for example.
//
// This operation is idempotent for memory already in the Prepared state, so
// it is safe to refer, with v and n, to a range of memory that includes both
// Prepared and Ready memory. However, the caller must provide the exact amount
// of Prepared memory for accounting purposes.
//
// 通知操作系统应用程序需要使用该内存区域，保证内存区域可以安全访问；
func sysUsed(v unsafe.Pointer, n, prepared uintptr) {
	gcController.mappedReady.Add(int64(prepared))
	sysUsedOS(v, n)
}

// sysHugePage does not transition memory regions, but instead provides a
// hint to the OS that it would be more efficient to back this memory region
// with pages of a larger size transparently.
func sysHugePage(v unsafe.Pointer, n uintptr) {
	sysHugePageOS(v, n)
}

// sysFree transitions a memory region from any state to None. Therefore, it
// returns memory unconditionally（无条件）. It is used if an out-of-memory error has been
// detected midway through an allocation or to carve out an aligned section of
// the address space. It is okay if sysFree is a no-op only if sysReserve always
// returns a memory region aligned to the heap allocator's alignment
// restrictions.
//
// sysStat must be non-nil.
//
// Don't split the stack as this function may be invoked without a valid G,
// which prevents us from allocating more stack.
// 会在程序发生内存不足（Out-of Memory，OOM）时调用并无条件地返回内存；
//
//go:nosplit
func sysFree(v unsafe.Pointer, n uintptr, sysStat *sysMemStat) {
	sysStat.add(-int64(n))
	gcController.mappedReady.Add(-int64(n))
	sysFreeOS(v, n)
}

// sysFault transitions a memory region from Ready to Reserved. It
// marks a region such that it will always fault if accessed. Used only for
// debugging the runtime.
//
// TODO(mknyszek): Currently it's true that all uses of sysFault transition
// memory from Ready to Reserved, but this may not be true in the future
// since on every platform the operation is much more general than that.
// If a transition from Prepared is ever introduced, create a new function
// that elides the Ready state accounting.
//
// 将内存区域转换成保留状态，主要用于运行时的调试；
func sysFault(v unsafe.Pointer, n uintptr) {
	gcController.mappedReady.Add(-int64(n))
	sysFaultOS(v, n)
}

// sysReserve transitions a memory region from None to Reserved. It reserves
// address space in such a way that it would cause a fatal fault upon access
// (either via permissions or not committing the memory). Such a reservation is
// thus never backed by physical memory.
//
// If the pointer passed to it is non-nil, the caller wants the
// reservation there, but sysReserve can still choose another
// location if that one is unavailable.
//
// NOTE: sysReserve returns OS-aligned memory, but the heap allocator
// may use larger alignment, so the caller must be careful to realign the
// memory obtained by sysReserve.
//
// 会保留操作系统中的一片内存区域，访问这片内存会触发异常；
func sysReserve(v unsafe.Pointer, n uintptr) unsafe.Pointer {
	return sysReserveOS(v, n)
}

// sysMap transitions a memory region from Reserved to Prepared. It ensures the
// memory region can be efficiently transitioned to Ready.
//
// sysStat must be non-nil.
//
// 保证内存区域可以快速转换至就绪状态；
func sysMap(v unsafe.Pointer, n uintptr, sysStat *sysMemStat) {
	sysStat.add(int64(n))
	sysMapOS(v, n)
}
