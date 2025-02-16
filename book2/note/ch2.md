第二章 快速入门


2.2.2 hello world

```go
import "fmt"
func main(){
	fmt.Println("hello world")
}
```

2.2.3 变量

Go语言是静态强类型，两层含义

1. 不会以隐含的方式自动转换变量的类型，必须手动转换
2. 变量的类型会尝试在编译时确定

```go
var age = 16

var a int

var ab int = 16
```

Go语言中大写字母开头的变量是可导出的，小写字母开头的变量是不可导出的。

一次性初始化多个变量

```go
var (
	name string
	age int
)
```

没有逗号，赋值本身就充当间隔符。用const来创建全局常量。

```go
const name = "llm"
```
必须要初始化。

```go
int 32/64位平台上32/64位有符号整数
uint 32/64位平台上32/64位无符号整数
uintptr 无符号整数，用于存放一个指针
```

数组

```go
names0 := []string{"kk","km"}

var names []string
names = []string{"kk","km"}

names = append(names,"k²")

```

Go中的数组存储在连续的内存块中。在追加时一般扩展。但是后续内存被使用的话，为了容纳更大数组，操作系统会分配一个全新的缓冲区，并复制之前的元素，释放之前的缓冲区。
这样巨慢。因此为了更高效，可以使用make

```go
names := make([]string,3)
names[0] = "Tanmay Bakshi"
names[1] = "Baheer Kamal"
names[2] = "Kathy"
```

2.2.4 if语句和switch语句

```go
func main(){
	age:=15
	if(age >= 18){
		fmt.Println("成年")
	}
	else{
		fmt.Println("未成年")
	}
}
```

现代编译器其实switch和if的速度没啥区别。看个人喜好吧。

```go
switch name{
	case "kk":
		fmt.Println("kk")
	case "km":
		fmt.Println("km")
	default:
		fmt.Println("default")
			
}
```

2.2.5 循环

仅使用一种循环来完成了所有循环。

第一种for-in循环

```go
for i in names{
	fmt.Println(i)
}
```
哦这是不对的，不要和python杂交了。

```go
for i := range names{
	fmt.Println(i)
}
```

实际上，上述代码输出的并不是names中的内容，而是1，2，3...即索引。

如果想输出元素，则

```go
for i := range names{
	fmt.Println(names[i])
}
```
这也是不对的。因为 := range返回的不只有值。还有索引。所以应该这样

```go
for k,v := range names{
	fmt.Println(k)
	fmt.Println(v)
}
```
忽略索引的代码如下

```go
for _,v := range names{
	fmt.Println(v)
}
```

用for循环模拟while循环

```go
i := 1
for i<1000{
	fmt.Println(i)
	i++
}
```
忽略了第一个和最后一个表达式。

2.2.6 函数

go语言不要求main一定有返回值，下面代码也可以。
```go
func main(){}
```

```go
func valueOfPi()float64{
	return 3.14
}
func main(){
	fmt.Println(valueOfPi())
}
```

来一个强制类型转换的例子，pi乘以任意常数

```go
func multiPi(multiplier uint)float32{
	return 3.14 * float32(multiplier)
}
```
我们必须添加无符号整数到float32类型的强制转换。安全性是go高度重视的东西。


传递函数指针给其他函数

```go
func runMathOp(a,b int, op func(int,int)int)int{
    return op(a,b)
}
func add(a,b int)int{
    return a+b
}
func main(){
    fmt.Println(runMathOp(1,2,add))
}
```

defer函数，很强大。在计算返回表达式之后，实际返回调用程序发生之前。 

泛型

```go
func Print[T any](s []T){
    for _,v := range s{
		fmt.Println(v)
}
}
```

go不是面向对象的语言。没有类和对象的概念。
```go
type User struct{
    Name string
    Age int
}
myuser := User{Name:"km",Age:16}
```

通过传递给函数的指针修改结构体

```go
func increaseAge(u *User){
    u.Age++
	fmt.Println(u.Age)
}
```
调用之后的确增加了。不同于C语言。可以使用u.Age，不需要什么*，->之类的。

通过方法，修改结构体
下面是两个代码。微小的区别

```go
func (u User)increaseAge(){
    u.Age++
    fmt.Println(u.Age)
}
func main(){
	myuser := User{Name:"km",Age:16}
    myuser.increaseAge()
    fmt.Println(myuser.Age)
}
```

```go
func (u *User)increaseAge(){
    u.Age++
    fmt.Println(u.Age)
}
func main(){
    myuser := User{Name:"km",Age:16}
    myuser.increaseAge()
    fmt.Println(myuser.Age)
}
```

前者输出17，16.后者17，17。因为前者改变的是副本。而后者，用了指针，直接在原结构体上修改。












