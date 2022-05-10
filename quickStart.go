package main

// q全局变量声明

import (
	"fmt"
	"time"
)

/*
	声明全局变量
*/
var myint int = 3   // 一般形式
var mystr = "hello" // 省略类型，自动推导

var ( // 声明全局变量 （因式分解关键字 ）
	br      = "\n"
	glb_it  = 3
	glb_str = "hello"
	glb_bl  = true
	glb_arr = []int{11, 22, 33, 44, 55, 66, 77, 88}
)

/*
	函数的基本格式
	func function_name( [parameter list] ) [return_types] {
   		函数体
	}
*/
func funcDemo_myswap(s1, s2 string) (string, string) {
	// 匿名函数
	func() {
		for i := 0; i < 3; i++ {
			fmt.Println("swap: ", i)
		}
	}() // 注意括号调用
	return s2, s1
}

/*
	变量，指针
	除 全局变量 和 常量 外, 声明的变量就要使用
*/
func variableDemo() {
	const cst_str = "this is const" // 定义常量
	// fmt.Println(cst_str)

	local_int := 0 // 函数内 快速定义，自动类型推导

	// 【定义指针】的方式
	var ptr *int = &local_int
	if ptr != nil {
		fmt.Println("pointer address: ", ptr)
		fmt.Println("pointer value: ", *ptr)
	}

	// 指针数组（存放指针的数组）
	var ptrarr [3]*int
	ptrarr[0] = &glb_it
	fmt.Print(ptrarr)

	// 通过指针交换两个值
	a, b := 1, 2
	fmt.Println("before swap ptr: ", a, b)
	// 匿名函数 交换指针
	func(p1, p2 *int) {
		tmp := p1
		p1 = p2
		p2 = tmp
	}(&a, &b)
	fmt.Println("after swap ptr: ", a, b)

	// 普通交换，有返回值
	s1, s2 := funcDemo_myswap("end", ">> golang")
	fmt.Println(s1, s2)

	multiStr :=
		`
	multiple
	lines
	string
	`
	fmt.Println(multiStr)
}

/*
	数组，切片
*/
func arrayDemo() {
	// 1. [size]
	var arr0 [3]int

	// 创建 变量大小的数组
	n := 3
	ans0 := make([]int, n)

	// 2. [...]
	arr1 := [...]int{4, 5, 6}

	// 3. 通过字面量在声明数组的同时快速初始化数组
	arr2 := []int{1, 2, 3}
	fmt.Println(ans0, arr0, arr1, arr2)

	// 4. 赋值指定位置元素
	arr3 := []string{0: "origin", 5: "endOfArr"}
	fmt.Println("sizeof, len arr3: ", cap(arr3), len(arr3), arr3)

	// 5. 切片
	slice1 := []int{2, 3, 4}
	// 使用make方法
	// make([]T, length, capacity)
	slice2 := make([]int, 2, 3)

	fmt.Println(slice1, slice2)

	// append 追加元素
	slc := []int{3, 4, 5}
	slc2 := slc[:2]
	slc3 := append(slc2, 5, 6, 7)

	fmt.Println(slc, slc2, slc3)
}

/*
	控制语句
	if, switch, for
*/
func controlDemo() {

	if t := glb_it; t > 0 {
		fmt.Println("glb_it > 0")
	} else {
		fmt.Println("glb_it <= 0")
	}

	// case 可以多写几个
	// 不同的 case 之间不使用 break 分隔，默认只会执行一个 case。
	fmt.Println("SWITCH: ")
	switch glb_it {
	case 0, 2, 4, 6, 8:
		fmt.Println("even number")
	case 1, 3, 5, 7, 9:
		fmt.Println("odd number")
		fallthrough
	case 999:
		fmt.Println("fallthrough anyway")
	default:
		fmt.Println("default")
	}

	// 常规 for，可以省略括号
	fmt.Println("FOR: ")
	for i := 0; i <= 10; i++ {
		fmt.Print(" no:", i)
		if i == 10 {
			fmt.Println(" ")
		}
	}

	// 死循环 for
	fmt.Println("FOR true: ")
	for true {
		glb_it++
		if glb_it >= 10 {
			break
		} else {
			fmt.Println(glb_it)
		}
	}

	// 遍历for，使用range
	fmt.Println("FOR RANGE: ")
	for idx, val := range glb_arr {
		fmt.Println("idx-val: ", idx, val)
	}

}

/*
	结构体
*/
type mystruct struct {
	name, addr string
	age, id    int
	checked    bool
}

func (r *mystruct) name_addr() string {
	fmt.Print("name, addr: ")
	return r.name + r.addr
}

func (r mystruct) age_id() (int, int) {
	fmt.Print("age, id: ")
	return r.age, r.id
}

func structDemo() {
	var stu1 mystruct
	stu1.name = "zhangsan"
	// 指针也可以使用 . 来访问
	stu11 := &stu1
	stu11.name = "lisi"

	fmt.Println(stu1, stu11)

	// 快速声明
	stu2 := mystruct{"zs", "佰京", 12, 12, true}
	// 可以指定 key : value
	stu3 := mystruct{addr: "长沙", name: "ls", id: 12, age: 12, checked: true}

	fmt.Println(stu3, stu2)

	// 调用结构体方法（类方法）
	fmt.Println(stu2.name_addr())
	fmt.Println(stu2.age_id())
}

/*
	字典Map
*/
func mapDemo() {
	map1 := map[string]int{"zs": 1, "ls": 2}
	fmt.Println(map1)

	delete(map1, "zs")
	fmt.Println(map1)
}

/*
	接口类型
	它把所有的具有共性的方法定义在一起，
	任何其他类型只要实现了这些方法就是实现了这个接口。
*/
type Phone interface {
	call()
	hello() string
}

type IP struct{}

func (ip IP) call() {
	fmt.Println("iphone is calling")
}

func (ip IP) hello() string {
	return "hello iphone"
}

type HW struct{}

func (hw HW) call() {
	fmt.Println("huawei is calling")
}

func (hw HW) hello() string {
	return "hello huawei"
}

func interfaceDemo() {

	var phone Phone

	phone = new(HW)
	phone.call()
	phone = new(IP)
	phone.call()

	ph2 := new(HW)     // 隐式声明ph2
	hl1 := ph2.hello() // hl1 是返回值
	ph3 := new(IP)
	hl2 := ph3.hello()

	fmt.Println(hl1, hl2)

}

/*
	异常处理，抛出错误
	1. 内置error接口
	type error interface {
    	Error() string
	}
	2. errmsg := errors.New("...")
*/
func (s mystruct) Error() string {
	errStr := `
	age cannot lowwer than 0
	your age is: %d`
	return fmt.Sprintf(errStr, s.age)
}

func (s mystruct) info() (string, int) {
	if ag := s.age; ag < 0 {
		return s.Error(), 0
	} else {
		return "", ag
	}
}

func errDemo() {
	p := mystruct{name: "monica", age: -13}
	fmt.Println(p.info())
}

/*	defer, panic, recover

	panic 与 recover 是 Go 的两个内置函数，
	这两个内置函数用于处理 Go 运行时的错误，
	panic 用于主动抛出错误，recover 用来捕获 panic 抛出的错误。

	[defer] 类似于finally
	Go语言的 defer 语句会将其后面跟随的语句进行延迟处理，
	在 defer 归属的函数即将返回时，
	将延迟处理的语句按 defer 的逆序进行执行

	引发panic有两种情况，
	一是程序主动调用，二是程序产生运行时错误，由运行时检测并退出。

	发生panic后，程序会从调用panic的函数位置或发生panic的地方立即返回，
	逐层向上执行函数的defer语句，然后逐层打印函数调用堆栈，
	直到被recover捕获或运行到最外层函数。

	panic不但可以在函数正常流程中抛出，
	在defer逻辑里也可以再次调用panic或抛出panic。
	defer里面的panic能够被后续执行的defer捕获。

	!!
	recover用来捕获panic，阻止panic继续向上传递。
	[recover()和defer一起使用]
	但是defer只有在后面的函数体内直接被掉用才能捕获panic来终止异常，
	否则返回nil，异常继续向外传递。
*/
func except(c *int) {
	old := *c
	*c = 25
	fmt.Println(old, "改变c的值", *c)
	recover()
}

func exceptionDemo() {
	a, b, c := 3, 4, 27

	defer except(&c)

	if ans := (a*a + b*b); ans != c {
		// recover()
		panic("勾三股四弦非五")
	} else {
		fmt.Println("勾三股四弦五")
	}
}

/*
	并发
	goroutine
*/
func mygoroutine(pre string) {
	for i := 0; i < 5; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Println(pre, i, ", ")
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // 手动关闭通道
	// 关闭通道并不会丢失里面的数据，
	//只是让读取通道数据的时候不会读完之后一直阻塞等待新数据写入
}

// !! 控制通道权限
func chanControl() {
	chAuth := make(chan int)

	//读写均可的channel
	go func(c chan int) {
		ipt := 123
		c <- ipt
		v := <-c
		fmt.Println("读写均可的channel ", v)
	}(chAuth)

	//只写的Channel
	go func(c chan<- int) {
		fmt.Println("只写的Channel ")
		ipt := 123
		c <- ipt
		// v := <-c
		// fmt.Println(v)
	}(chAuth)

	//只读的Channel
	go func(c <-chan int) {
		// ipt := 123
		// c <- ipt
		v := <-c
		fmt.Println("只读的Channel ", v)
	}(chAuth)

	func() {
		fmt.Println("end")
	}()

	// fmt.Println(<-ch)
	// 已读的 chan 不可以再读一次
	time.Sleep(150 * time.Millisecond)
}

func goroutineDemo() {

	// 协程与普通调用
	mygoroutine("normal: ")
	go mygoroutine("goroutine: ")

	// chan
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int) // 声明通道
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收
	fmt.Println(x, y, x+y)

	// demo3 带缓冲的chan
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	c2 := make(chan int, 10)
	go fibonacci(cap(c2), c2)
	for i := range c2 {
		fmt.Println(i)
	}

	// demo4
}

/*	select

	select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。
	每个 case 必须是一个通信操作，要么是发送要么是接收。

	select 随机执行一个可运行的 case。如果没有 case 可运行，
	它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。

	1. 每个 case 都必须是一个通信
	2. 所有 channel 表达式都会被求值
	3. 所有被发送的表达式都会被求值
	4. 如果任意某个通信可以进行，它就执行，其他被忽略。
	5. 如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。
	其他不会执行。
		否则：
		5.1 如果有 default 子句，则执行该语句。
		5.2 如果没有 default 子句，select 将阻塞，直到某个通信可以运行；
			Go 不会重新对 channel 或值进行求值。
*/
func selectDemo() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Print("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Print("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Print("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}

func test() {
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 3; j++ {
			fmt.Println(i, j)
		}
	}
}

func main() {
	// chanControl()

	// goroutineDemo()

	// exceptionDemo()

	// errDemo()

	// interfaceDemo()

	// mapDemo()

	// structDemo()

	// selectDemo()

	// controlDemo()

	// arrayDemo()

	// variableDemo()

	// fmt.Print("hello world")

	test()

	fmt.Println(br, "结束")
}
