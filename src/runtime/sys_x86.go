// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64 || 386

package runtime

import (
	"internal/goarch"
	"unsafe"
)

// 设置sched.pc等于目标函数的地址
// 调整newg的栈空间，把goexit函数的第二条指令的地址入栈，伪造成goexit函数调用了fn，
// 从而使fn执行完成后执行ret指令时返回到goexit继续执行完成最后的清理工作；
// 重新设置newg.buf.pc 为需要执行的函数的地址，即fn，我们这个场景为runtime.main函数的地址。
//
// adjust Gobuf as if it executed a call to fn with context ctxt
// and then stopped before the first instruction in fn.
func gostartcall(buf *gobuf, fn, ctxt unsafe.Pointer) {
	sp := buf.sp         // newg的栈顶，初始化时 newg 栈上只有fn函数的参数，sp指向的是fn的第一参数
	sp -= goarch.PtrSize //为返回地址预留空间，
	// 这里在伪装fn是被goexit函数调用的，使得fn执行完后返回到goexit继续执行，从而完成清理工作
	*(*uintptr)(unsafe.Pointer(sp)) = buf.pc //在栈上放入goexit+1的地址
	//重新设置newg的栈顶寄存器
	buf.sp = sp
	//这里才真正让newg的ip寄存器指向fn函数，注意，这里只是在设置newg的一些信息，newg还未执行，
	//等到newg被调度起来运行时，调度器会把buf.pc放入cpu的IP寄存器，
	//从而使newg得以在cpu上真正的运行起来
	buf.pc = uintptr(fn)
	buf.ctxt = ctxt
}
