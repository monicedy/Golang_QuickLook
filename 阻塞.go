/*
	无缓冲是同步的，例如 make(chan int)，
	就是一个送信人去你家门口送信，你不在家他不走，
	你一定要接下信，他才会走，
	【无缓冲保证信能到你手上】

	有缓冲是异步的，例如 make(chan int, 1)，
	就是一个送信人去你家仍到你家的信箱，转身就走，
	除非你的信箱满了，他必须等信箱空下来，
	【有缓冲的保证信能进你家的邮箱】

	阻塞即停下，后面的代码都它结束

	更多的看评论区
	https://www.runoob.com/go/go-concurrent.html
*/

package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Printf("sum:")
	fmt.Printf("%#v\n", sum)
	c <- sum // 把 sum 发送到通道 c
	fmt.Println("after channel pro")
}

// 通道不带缓冲，表示是同步的，只能向通道 c 发送一个数据，只要这个数据没被接收然后所有的发送就被阻塞
func main() {
	s := []int{7, 2, 8, -9, 4, 0, 6, 1}
	// c := make(chan int)
	c := make(chan int, 2)

	fmt.Println("go [0,2]")
	go sum(s[0:2], c) //a

	//这里开启一个新的运行期线程，这个是需要时间的，本程序继续往下走

	fmt.Println("go [2,4]")
	go sum(s[2:4], c) //b
	fmt.Println("go2 [4,6]")
	go sum(s[4:6], c) //c
	fmt.Println("go2 [6,8]")
	go sum(s[6:8], c) //d

	/*
	   a b c d和main一起争夺cpu的，他们的执行顺序完全无序，甚至里面不同的语句都相互穿插
	   但无缓冲的等待是同步的，所以接下来a b c d都会执行，一直执行到c <- sum后，开始同步阻塞
	   因此after channel pro是打印不出来的, 等要打印after channel pro的时候，main就结束了
	*/

	fmt.Println("go3 start waiting...")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("go3 waited 1000 ms")

	//因为a b c d都在管道门口等着，这里读一个，a b c d就继续一个，这个结果的顺序可能是acbd
	aa := <-c
	bb := <-c
	fmt.Println(aa)
	fmt.Println(bb)
	x, y := <-c, <-c
	fmt.Println(x, y)
}
