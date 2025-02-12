第二章 指针

2.1 指针构成

指针是一个变量，其值为另一个变量的地址。指针变量的声明格式为：
```go
var p *int
```

其中，p是指针变量的名称，int是指针变量的类型。指针变量的值为另一个变量的地址，可以通过指针变量访问另一个变量的值。指针变量的值为nil时，表示指针变量没有指向任何变量。
有着不同元素类型的指针被视为不同的类型，不能相互赋值。指针变量的值可以通过&运算符获取，&运算符返回变量的地址。指针变量的值可以通过*运算符获取，*运算符返回指针变量指向的变量的值。

2.2.1 地址

地址在运行阶段用来在进程的内存空间中定位变量。地址是一个无符号整数，表示变量在内存中的位置。ch1介绍了x86的常见寻址方式，一般指针会用到基址寄存器和偏移量，基址寄存器存储变量的基地址，偏移量存储变量的偏移量。指针变量的值为另一个变量的地址，可以通过指针变量访问另一个变量的值。指针变量的值为nil时，表示指针变量没有指向任何变量。

在amd64架构下通过go build命令编译的一个例子如下：

```go
package main

import "fmt"

func main() {
	n := 10
	Println(read(&n))
}

func read(p *int) (v int) {
	v = *p
	return v
}
```
使用Go自带的objdump工具反编译上述函数，得到汇编代码如下
```asm
TEXT main.read(SB), NOSPLIT, $0
    MOVQ p+0(FP), AX
    MOVQ (AX), AX
    MOVQ AX, ret+8(FP)
    RET
``` 
在汇编代码中，`MOVQ p+0(FP), AX`表示将p的值传送到AX寄存器中，`MOVQ (AX), AX`表示将AX寄存器的值传送到AX寄存器中，`MOVQ AX, ret+8(FP)`表示将AX寄存器的值传送到ret寄存器中。

2.1.2 元素类型

将read()函数改为read32()函数，如下
```go
package main

import "fmt"

func main() {
	n := 10
	Println(read(&n))
}

func read32(p *int) (v int32) {
	v = *p
	return v
}
```

得到的汇编代码如下
```asm
TEXT main.read32(SB), NOSPLIT, $0
    MOVQ 0x8(SP), AX
    MOVQ (AX), AX
    MOVL AX,0x10(SP)
    RET
```
在汇编代码中，`MOVQ 0x8(SP), AX`表示将SP寄存器的值加上8传送到AX寄存器中，`MOVQ (AX), AX`表示将AX寄存器的值传送到AX寄存器中，`MOVL AX,0x10(SP)`表示将AX寄存器的值传送到SP寄存器的值加上16的地址中。
为什么要将SP寄存器的值加上8传送到AX寄存器中呢？因为在amd64架构下，指针的大小是8个字节，所以需要将SP寄存器的值加上8传送到AX寄存器中。

2.2 相关操作

2.2.1 取地址

在Go语言中，可以通过&运算符获取变量的地址。&运算符返回变量的地址。例如
```go
package main

var n int

func main(){
    println(addr(&n))
}

func addr(p *int) int {
    return *p
}
```
得到的反编译代码如下
```asm
TEXT main.addr(SB), NOSPLIT, $0
    LEAQ n(SB), AX
    MOVQ AX, 0x8(SP)
    RET
```
在汇编代码中，`LEAQ n(SB), AX`表示将n的地址传送到AX寄存器中，`MOVQ AX, 0x8(SP)`表示将AX寄存器的值传送到SP寄存器的值加上8的地址中。

在C语言中，不应该返回局部变量的地址，因为局部变量的地址在函数返回后会被释放，所以返回局部变量的地址是不安全的。在Go语言中，可以返回局部变量的地址，因为Go语言的垃圾回收器会自动回收不再使用的内存。

在Go语言中，通过逃逸分析机制，可以判断变量是否逃逸到堆上。如果变量逃逸到堆上，那么变量的地址就不会被释放，否则变量的地址会被释放。

2.2.2 解引用

通过指针中的地址取访问原来的变量，这就是解引用。在Go语言中，可以通过*运算符获取指针变量指向的变量的值。在编译层面，解引用就是把地址存入某个通用寄存器，然后用作基址进行寻址。

接下来介绍C语言中与指针解引用相关的几个常见问题，以及这些问题在Go语言中如何解决的。

1. 空指针异常

空指针是地址值为0的指针，空指针异常是指针指向的地址没有被分配，访问该地址会导致异常。按照操作系统的内存管理机制，操作系统会将未分配的内存映射到一个特殊的地址，访问这个地址会导致异常。
所以对空指针进行解引用会导致异常。在Go语言中解引用会导致panic异常，宕机。

2. 野指针异常

野指针是指针指向的地址已经被释放，访问该地址会导致异常。C语言对于未初始化的指针，指针的值是随机的，这就是野指针。
Go语言中声明的变量默认都会初始化为零值，指针类型变量都会初始化为nil，所以不会出现未初始化的指针。另外，Go语言的垃圾回收器会自动回收不再使用的内存，所以不会出现野指针。

3. 悬挂指针异常

在C语言中，需要手动分配和释放内存，如果释放了一个指针，但是没有将指针置为nil，那么这个指针就是悬挂指针。悬挂指针是指针指向的地址已经被释放，但是指针的值没有被置为nil，访问该地址会导致异常。
程序过早地释放内存，但是仍然保留了指向该内存的指针，这就是悬挂指针。Go语言实现了自动内存管理，由GC负责内存的分配和释放，基于标记-清除算法，进行对象的存活判断，自动回收不再使用的内存。只有明确不可达的对象才会被回收，所以不会出现悬挂指针。

2.2.3 强制类型转换

出于安全性考虑，Go语言不建议频繁指针之间的强制类型转换。在Go语言中，指针的类型是不同的，不同元素类型的指针被视为不同的类型，不能相互赋值。两种不同类型的指针的转换需要通过unsafe包实现。

在Go语言中，可以通过unsafe.Pointer将指针转换为通用指针，然后通过类型断言将通用指针转换为指定类型的指针。例如
```go
func convert(p *int) *int32{
    return (*int32)(unsafe.Pointer(p))
}
```
得到的汇编代码如下
```asm
TEXT main.convert(SB), NOSPLIT, $0
    MOVQ 0x8(SP),AX
    MOVL $ 0x0,0(AX)
    RET
```
在汇编代码中，`MOVQ 0x8(SP),AX`表示将SP寄存器的值加上8传送到AX寄存器中，`MOVL $ 0x0,0(AX)`表示将0传送到AX寄存器的地址中。
区别就在于MOVQ变成了MOVL，表示传送的是32位的数据。

2.2.4 指针运算

Go语言中的数组必须指定长度，数组的长度是数组类型的一部分，而且是值类型，因此与指针不再等价。指针运算也是不允许的，因为指针运算会导致内存越界，Go语言的内存管理机制不允许内存越界。
这和C语言是不同的，C语言中数组是指针，数组的长度不是数组类型的一部分，数组是引用类型，数组的长度是数组的属性，数组的长度是数组的一部分。指针运算是允许的，指针加上一个整数，指针会移动整数个元素的位置。

因此Go语言的slice集成了数组的优点和指针的优点。既可以像指针关联一个可以动态增长的Buffer，又可以像数组让编译器生成下标越界检查的代码。

假如要实现C语言中的指针运算，可以通过切片来实现。例如
```C
++p
```
在Go语言中可以通过切片来实现
```go
p = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1))
```
uintptr是一个整数类型，用于存放指针的整数值。unsafe.Pointer是一个通用指针类型，可以存放任意类型的指针.uintptr(unsafe.Pointer(p))将指针p转换为整数，uintptr(unsafe.Pointer(p)) + 1将整数加1，unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)将整数转换为指针，
(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Sizeof(*p)))将指针转换为int类型的指针。
unsafe.Pointer返回的是一个指针，uintptr返回的是一个整数，uintptr(unsafe.Pointer(p))将指针转换为整数，unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)将整数转换为指针。

2.3 unsafe包

unsafe.Pointer进行指针的强制类型转换，unsafe.Sizeof返回变量的大小，unsafe.Offsetof返回结构体成员的偏移量，unsafe.Alignof返回变量的对齐方式。unsafe包提供了一些不安全的操作。实际上是人为地干预编译器对Go语言的内存管理机制。
代码中用好unsafe包，可以提高代码的性能。

经典的类型转换代码如下：
```go
func convert(s []byte) string{
    return *(*string)(unsafe.Pointer(&s))
}
```
它的意思是将s转换为string类型，但是这样的转换是不安全的，因为string类型是只读的，不能修改，但是s是可写的，可以修改。相当于Slice中内嵌了一个string类型。
这导致string和原来的Slice共享底层的内存，如果修改了Slice，string也会被修改，这是不安全的。

unsafe包包含的操作绕过了安全机制。

2.3.1 标准库与keyword

unsafe的本质是标准库还是keyword，这个问题在Go语言中是一个有争议的问题。

从unsafe源码入手
```go
//一个任意类型的定义
type ArbitraryType int
//一个任意类型的指针
type Pointer *ArbitraryType
//3个工具函数原型
func Alignof(variable ArbitraryType) uintptr
func Offsetof(selector ArbitraryType) uintptr
func Sizeof(variable ArbitraryType) uintptr
```
ArbitraryType不属于unsafe包，可以表达任意Go表达式类型。Sizeof用来返回任意类型的大小，Offsetof用来返回任意类型的偏移量，Alignof用来返回任意类型的对齐方式。三个返回值都是uintptr类型。

所以unsafe不是包，而是一个keyword，是一个关键字。

指针强制类型转换是在编译阶段实现，而Sizeof、Offsetof、Alignof是在运行阶段实现的。也就是要求返回值必须在编译阶段确定。

代码验证一下
```go
func size()(o unitptr){
    o = unsafe.Sizeof(o)
    return o
}
```
得到的汇编代码如下
```asm
TEXT main.size(SB), NOSPLIT, $0
    MOVQ $8,AX
    MOVQ AX,0(SP)
    RET
```
在汇编代码中，`MOVQ $8,AX`表示将8传送到AX寄存器中，`MOVQ AX,0(SP)`表示将AX寄存器的值传送到SP寄存器的地址中。
这条汇编代码表示Sizeof返回的是8。说明了Sizeof是在编译阶段实现的。unsafe这个名字就是想告诉你，这个操作是不安全的，可能会导致程序崩溃。

2.3.2 关于uintptr

它不是一个指针，而是一个整数类型。不要用uintptr来存储堆上对象的地址。因为uintptr是一个整数类型，它不会被GC扫描，如果一个对象的地址被存储在uintptr中，会被GC忽略，导致对象被回收。所以不要用uintptr来存储堆上对象的地址。

2.3.3 内存对齐

complex类型由实部和虚部组成，实部和虚部都是float64类型，所以complex类型的大小是16个字节。complex类型的对齐方式是8个字节。complex类型的偏移量是0个字节。
map大多数情况下被分配到堆上，map的大小是8个字节。这是因为map类型是一个指针类型，指向一个runtime.hmap结构体。runtime.hmap结构体的大小是8个字节是因为runtime.hmap结构体只有一个指针类型的字段，指向一个runtime.bmap结构体。runtime.bmap结构体的大小是8个字节，所以map的大小是8个字节。map的对齐方式是8个字节。map的偏移量是0个字节。

string和slice的结构定义如下：
```go
type StringHeader struct {
    Data uintptr
    Len int
}
type SliceHeader struct {
    Data uintptr
    Len int
    Cap int
}
```
它们的对齐边界与uintptr相同，即8字节。所以string和slice的对齐方式是8个字节。string和slice的偏移量是0个字节。
对于struct而言，每个成员都会以结构题的起始地址为基准，按照成员的大小对齐，如果成员的大小不是对齐边界的整数倍，会在成员之后填充一些字节，这也被称为padding。

来看一个实例：
```go
type T struct {
    a int8
    b int16
	
    c int32
    d int64
	e int8
}
```
这个T在amd64架构上占用了32个字节，因为int8占用1个字节，int16占用2个字节，int32占用4个字节，int64占用8个字节。T的对齐方式是8个字节，T的偏移量是0个字节。
a后面的padding为7个字节，b后面的padding为6个字节，c后面的padding为4个字节，d后面的padding为0个字节，e后面的padding为7个字节。所以T的大小是32个字节。
而我们可以人为地调整结构体成员的顺序，使得结构体的大小变小，这就是内存对齐。

代码如下：
```go
type T struct {
    a int8
    e int8
    b int16
    c int32
    d int64
}
```
padding为0.



