package main

//main.main STEXT size=197 args=0x0 locals=0x98 funcid=0x0 align=0x0
// 0x0000: 当前指令相对于当前函数的偏移量；
// TEXT:由于程序代码在运行期会放在内存的 .text 段中，所以TEXT 是一个指令，用来定义一个函数；
// "".main(SB): 表示的是包名.函数名，这里省略了包名。SB是一个虚拟寄存器，保存了静态基地址(static-base) 指针，
//  			即我们程序地址空间的开始地址；
// $24-0:$24表即将分配的栈帧大小；0指定了调用方传入的参数大小。
//0x0000 00000 (main.go:76)   TEXT    main.main(SB), ABIInternal, $152-0
//0x0000 00000 (main.go:76)   LEAQ    -24(SP), R12
//0x0005 00005 (main.go:76)   CMPQ    R12, 16(R14) ; 栈溢出检测
//0x0009 00009 (main.go:76)   PCDATA  $0, $-2      ; GC 相关
//0x0009 00009 (main.go:76)   JLS     185
//0x000f 00015 (main.go:76)   PCDATA  $0, $-1      ; GC 相关
// 在执行栈上调用的时候由于栈是从内存地址高位向低位增长的，所以会根据当前的栈帧大小
// 调用 SUBQ $32, SP 表示分配 32bytes 的栈内存
//0x000f 00015 (main.go:76)   SUBQ    $152, SP     ; 分配了 152bytes 的栈地址
//0x0016 00022 (main.go:76)   MOVQ    BP, 144(SP)  ; 将 BP 的值存储到栈上
//0x001e 00030 (main.go:76)   LEAQ    144(SP), BP  ; 将刚分配的栈空间 8bytes 的地址赋值给BP
//0x0026 00038 (main.go:76)   FUNCDATA        $0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
//0x0026 00038 (main.go:76)   FUNCDATA        $1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
//0x0026 00038 (main.go:77)   MOVQ    $4, (SP)     ; 将给add函数的第一个参数1，写到SP
//0x002e 00046 (main.go:77)   MOVQ    $4, 8(SP)    ; 将给add函数的第二个参数2，写到SP
//0x0037 00055 (main.go:77)   MOVQ    $5, 16(SP)
//0x0040 00064 (main.go:77)   MOVQ    $5, 24(SP)
//0x0049 00073 (main.go:77)   MOVQ    $6, 32(SP)
//0x0052 00082 (main.go:77)   MOVQ    $7, 40(SP)
//0x005b 00091 (main.go:77)   MOVQ    $7, 48(SP)
//0x0064 00100 (main.go:77)   MOVQ    $88, 56(SP)
//0x006d 00109 (main.go:77)   MOVQ    $8, 64(SP)
//0x0076 00118 (main.go:77)   MOVL    $1, AX
//0x007b 00123 (main.go:77)   MOVL    $2, BX
//0x0080 00128 (main.go:77)   MOVL    $4, CX
//0x0085 00133 (main.go:77)   MOVL    $5, DI
//0x008a 00138 (main.go:77)   MOVL    $6, SI
//0x008f 00143 (main.go:77)   MOVL    $7, R8
//0x0095 00149 (main.go:77)   MOVQ    R8, R9
//0x0098 00152 (main.go:77)   MOVL    $88, R10
//0x009e 00158 (main.go:77)   MOVL    $8, R11
//0x00a4 00164 (main.go:77)   PCDATA  $1, $0
//0x00a4 00164 (main.go:77)   CALL    main.add(SB)   ; 调用 add 函数
//0x00a9 00169 (main.go:78)   MOVQ    144(SP), BP    ; 将栈上储存的值恢复BP
//0x00b1 00177 (main.go:78)   ADDQ    $152, SP       ; 增加SP的值，栈收缩，收回 32 bytes的栈空间
//0x00b8 00184 (main.go:78)   RET
//0x00b9 00185 (main.go:78)   NOP
//0x00b9 00185 (main.go:76)   PCDATA  $1, $-1
//0x00b9 00185 (main.go:76)   PCDATA  $0, $-2
//0x00b9 00185 (main.go:76)   CALL    runtime.morestack_noctxt(SB)
//0x00be 00190 (main.go:76)   PCDATA  $0, $-1
//0x00be 00190 (main.go:76)   NOP
//0x00c0 00192 (main.go:76)   JMP     0
//0x0000 4c 8d 64 24 e8 4d 3b 66 10 0f 86 aa 00 00 00 48  L.d$.M;f.......H
//0x0010 81 ec 98 00 00 00 48 89 ac 24 90 00 00 00 48 8d  ......H..$....H.
//0x0020 ac 24 90 00 00 00 48 c7 04 24 04 00 00 00 48 c7  .$....H..$....H.
//0x0030 44 24 08 04 00 00 00 48 c7 44 24 10 05 00 00 00  D$.....H.D$.....
//0x0040 48 c7 44 24 18 05 00 00 00 48 c7 44 24 20 06 00  H.D$.....H.D$ ..
//0x0050 00 00 48 c7 44 24 28 07 00 00 00 48 c7 44 24 30  ..H.D$(....H.D$0
//0x0060 07 00 00 00 48 c7 44 24 38 58 00 00 00 48 c7 44  ....H.D$8X...H.D
//0x0070 24 40 08 00 00 00 b8 01 00 00 00 bb 02 00 00 00  $@..............
//0x0080 b9 04 00 00 00 bf 05 00 00 00 be 06 00 00 00 41  ...............A
//0x0090 b8 07 00 00 00 4d 89 c1 41 ba 58 00 00 00 41 bb  .....M..A.X...A.
//0x00a0 08 00 00 00 e8 00 00 00 00 48 8b ac 24 90 00 00  .........H..$...
//0x00b0 00 48 81 c4 98 00 00 00 c3 e8 00 00 00 00 66 90  .H............f.
//0x00c0 e9 3b ff ff ff                                   .;...
//rel 165+4 t=7 main.add+0
//rel 186+4 t=7 runtime.morestack_noctxt+0
//main.add STEXT nosplit size=106 args=0x90 locals=0x10 funcid=0x0 align=0x0
//0x0000 00000 (main.go:80)   TEXT    main.add(SB), NOSPLIT|ABIInternal, $16-144
//0x0000 00000 (main.go:80)   SUBQ    $16, SP
//0x0004 00004 (main.go:80)   MOVQ    BP, 8(SP)
//0x0009 00009 (main.go:80)   LEAQ    8(SP), BP
//0x000e 00014 (main.go:80)   FUNCDATA        $0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
//0x000e 00014 (main.go:80)   FUNCDATA        $1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
//0x000e 00014 (main.go:80)   FUNCDATA        $5, main.add.arginfo1(SB)
//0x000e 00014 (main.go:80)   MOVQ    AX, main.a+96(SP)
//0x0013 00019 (main.go:80)   MOVQ    BX, main.b+104(SP)
//0x0018 00024 (main.go:80)   MOVQ    CX, main.i+112(SP)
//0x001d 00029 (main.go:80)   MOVQ    DI, main.i2+120(SP)
//0x0022 00034 (main.go:80)   MOVQ    SI, main.i3+128(SP)
//0x002a 00042 (main.go:80)   MOVQ    R8, main.i4+136(SP)
//0x0032 00050 (main.go:80)   MOVQ    R9, main.i5+144(SP)
//0x003a 00058 (main.go:80)   MOVQ    R10, main.i6+152(SP)
//0x0042 00066 (main.go:80)   MOVQ    R11, main.i7+160(SP)
//0x004a 00074 (main.go:80)   MOVQ    $0, main.~r0(SP)    ; 初始化返回值
//0x0052 00082 (main.go:81)   MOVQ    main.a+96(SP), AX   ; AX = 1
//0x0057 00087 (main.go:81)   ADDQ    main.b+104(SP), AX  ; AX = AX + 2
//0x005c 00092 (main.go:81)   MOVQ    AX, main.~r0(SP)    ; (24)SP = AX = 3
//0x0060 00096 (main.go:81)   MOVQ    8(SP), BP
//0x0065 00101 (main.go:81)   ADDQ    $16, SP
//0x0069 00105 (main.go:81)   RET
//0x0000 48 83 ec 10 48 89 6c 24 08 48 8d 6c 24 08 48 89  H...H.l$.H.l$.H.
//0x0010 44 24 60 48 89 5c 24 68 48 89 4c 24 70 48 89 7c  D$`H.\$hH.L$pH.|
//        0x0020 24 78 48 89 b4 24 80 00 00 00 4c 89 84 24 88 00  $xH..$....L..$..
//        0x0030 00 00 4c 89 8c 24 90 00 00 00 4c 89 94 24 98 00  ..L..$....L..$..
//        0x0040 00 00 4c 89 9c 24 a0 00 00 00 48 c7 04 24 00 00  ..L..$....H..$..
//        0x0050 00 00 48 8b 44 24 60 48 03 44 24 68 48 89 04 24  ..H.D$`H.D$hH..$
//0x0060 48 8b 6c 24 08 48 83 c4 10 c3                    H.l$.H....
//go:cuinfo.producer.<unlinkable> SDWARFCUINFO dupok size=0
//0x0000 2d 4e 20 2d 6c 20 72 65 67 61 62 69              -N -l regabi
//go:cuinfo.packagename.main SDWARFCUINFO dupok size=0
//0x0000 6d 61 69 6e                                      main
//main..inittask SNOPTRDATA size=24
//0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
//0x0010 00 00 00 00 00 00 00 00                          ........
//gclocals·g2BeySu+wFnoycgXfElmcg== SRODATA dupok size=8
//0x0000 01 00 00 00 00 00 00 00                          ........
//main.add.arginfo1 SRODATA static dupok size=22
//0x0000 48 08 50 08 58 08 60 08 68 08 70 08 78 08 80 08  H.P.X.`.h.p.x...
//        0x0010 88 08 00 08 fc ff                                ......

func main() {
	add(1, 2, 4, 5, 6, 7, 7, 88, 8, 4, 4, 5, 5, 6, 7, 7, 88, 8)
}

func add(a, b, i, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16 int) int {
	return a + b
}
