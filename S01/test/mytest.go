package mytest

import (
	simle_math "S01/math"
	"fmt"
	"math"
	"unsafe"
)

/*
包结构
*/
func Test01() {
	a := 3
	b := 5
	fmt.Printf("this is %d\n", a)
	fmt.Printf("this is %d\n", b)

	sum := simle_math.Add(a, b)
	fmt.Printf("a + b = %d\n", sum)
	sub := simle_math.Sub(a, b)
	fmt.Printf("a - b = %d\n", sub)
}

/*
格式化字符串
*/
func Test02() {
	// Sprint()，可用来做拼接输出
	concat_str := fmt.Sprintln("hello,", 1+2, "abcd")
	fmt.Print(concat_str) // hello, 3 abcd

	// Sprintf()，返回一个格式化的字符串
	url := fmt.Sprintf("the url to github is %s\n", "www.github.com")
	fmt.Print(url) // the url to github is www.github.com

	// Printf()，用来格式化输出
	fmt.Printf("today is %s\n", "2024/6/3") // today is 2024/6/3
}

var ( // 这种因式分解关键字的写法一般用于声明全局变量，全局变量是允许声明但不使用的。
	q int
	w bool
)

/*
变量
*/
func Test03() {
	// 变量的声明
	var a int
	a = 99
	println("a = ", a)

	// 如果未手动赋值，对于三种基本类型（数字、布尔型、字符串）默认赋值为0或false或空串
	// 对于其他类型是nil
	var b int
	var flag bool
	var s string
	fmt.Printf("b = %d, flag = %t, s = %s\n", b, flag, s) // b = 0, flag = false, s =

	// 可以不手动指定类型，自动类型判断。只能被用在函数体内，而不可以用于全局变量的声明与赋值
	what := 3.14 // 如果变量已经生命过，再用 := 会报错
	var who = "Tomcat"
	fmt.Printf("what = %f, and who = %s\n", what, who) // what = 3.140000, and who = Tomcat

	// 尝试使用全局变量
	q = 13
	w = true
	fmt.Printf("q = %d, w = %t\n", q, w) // q = 13, w = true

	// 同类型的多个变量可以声明在同一行，在最后指定类型
	var i, j, k int
	var str string
	// 即使是不同类型的变量，也可以在同一行赋值
	i, j, k, str = 12, 23, 34, "this is a normal str"
	println(i, j, k, str) // 12 23 34 this is a normal str
	// 也可以不同类型的变量在同一行自动推断类型并赋值（并行赋值）
	n, m := 8848, "钛金手机"
	println("n = ", n, ", m = ", m) // n =  8848 , m =  钛金手机

	// 交换两个相同类型变量的值
	i, j = j, i
	println(i, j) // 23 12

	// 空白标识符 _可用于抛弃不需要使用的值。它是一个只读变量，不能获取到它的值。
	// 可以用它来接收不想使用的函数返回值（一个函数可以返回多个值）
	_ = a
}

/*
常量
*/
func Test04() {
	// 常量用const来声明
	const width = 3
	const length = 4
	var arc = width * length
	println("面积为：", arc)

	// 常量可以作为枚举
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)
	println(Unknown, Female, Male)

	// 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过
	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	fmt.Printf("a = %s, b = %d, c = %d\n", a, b, c) // a = abc, b = 3, c = 16

	// iota 特殊常量
	// iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次
	// (iota 可理解为 const 语句块中的行索引)。
	const (
		q = iota //0
		w        //1
		e        //2
		r = "ha" //独立值，iota += 1
		t        //"ha"   iota += 1
		y = 100  //iota +=1
		u        //100  iota +=1
		i = iota //7,恢复计数
		o        //8
	)
	fmt.Println(q, w, e, r, t, y, u, i, o) // 0 1 2 ha ha 100 100 7 8
	// 再次出现const将使iota重新计数
	const (
		k = iota
	)
	fmt.Printf("now iota = %d", k) // now iota = 0
}

/*
Go语言函数可以返回多个值，返回几个值就在返回值列表中写几个对应的类型
*/
func getLargerAndSmaller(x, y int) (int, int) {
	return (x + y + int(math.Abs(float64(x-y)))) / 2, (x + y - int(math.Abs(float64(x-y)))) / 2
}

/*
函数
*/
func Test05() {
	a, b := 3, 5
	larger, smaller := getLargerAndSmaller(a, b)
	fmt.Printf("the bigger and smaller num between %d and %d is : %d and %d\n", a, b, larger, smaller)
}

func printArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		print(arr[i], ", ")
	}
	println()
}

/*
数组
*/
func Test06() {
	// 声明一个大小为5的未初始化的数组，默认初始化元素为0
	var arr1 [5]int
	printArr(arr1[:]) // 以切片方式传参，可以节省内存，也可以避免容量不同的数组之间类型不匹配 0, 0, 0, 0, 0,

	// 也可以不指定数组长度，用...
	arr2 := [...]int{1, 2, 3, 4}
	printArr(arr2[:]) // 1, 2, 3, 4,
	arr3 := []int{4, 3, 2}
	printArr(arr3[:]) // 4, 3, 2,

	// 如果设置了数组的长度，我们还可以通过指定下标来初始化元素
	arr4 := [5]int{1: 9, 3: 5}
	printArr(arr4[:]) // 0, 9, 0, 5, 0,
}

/*
指针
*/
func Test07() {
	a := 999
	// 声明指针类型
	var ptr_a *int
	ptr_a = &a
	ptr_b := &a

	*ptr_a = 10

	fmt.Printf("a = %d\n", a)            // a = 10
	fmt.Printf("a的内存地址：%x\n", &a)        // a的内存地址：c00000a0a8
	fmt.Printf("ptr_a = %x\n", ptr_a)    // ptr_a = c00000a0a8
	fmt.Printf("ptr_a指向的值：%d\n", *ptr_a) // ptr_a指向的值：10
	fmt.Printf("ptr_b = %x\n", ptr_b)    // ptr_b = c00000a0a8
	fmt.Printf("ptr_b指向的值：%d\n", *ptr_b) // ptr_b指向的值：10
}
