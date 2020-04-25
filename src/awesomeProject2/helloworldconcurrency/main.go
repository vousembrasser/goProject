package main

import (
	"fmt"
)

//并发编程
func main() {

	//new 的作用是初始化一个指向类型的指针(*T)，make 的作用是为 slice，map 或 chan 初始化并返回引用(T)。
	ch := make(chan string)

	//:= 申明变量并赋值
	for i := 0; i < 5; i++ {
		/*go starts go goroutine
		goroutine 是并发的执行 不是等到main结束之后再去执行
		类似线性，优于线程
		goroutine 是编译器级别的多任务和线程不一样，很多个goroutine能同时映射到一个物理线程上去
		*/
		go printHelloWorld(i, ch)
	}

	for {
		msg := <-ch
		fmt.Println(msg)
	}

	//time.Sleep(5 * time.Millisecond)
	//fmt.Println("over!")
}

//chan string译位chan of string (自我理解 应该是sting类型的chan)
func printHelloWorld(i int, ch chan string) {
	//go语言没有while语句 用for{}表示while
	for {
		ch <- fmt.Sprintf("hello world %d!\n", i)
	}
}
