package main

import (
	"awesomeProject2/pipeline"
	"fmt"
)

func main() {
	p := pipeline.ArraySource(3, 2, 6, 7, 345, 23, 6, 8, 3, 7)
	for {
		/*
			读取完成之后 会一直打印0
		*/
		//num := <-p
		//fmt.Println(num)

		//方法1
		if num, ok := <-p; ok {
			fmt.Println(num)
		} else {
			break
		}
	}

	fmt.Println("---------------------")
	//方法2
	for v := range p {
		fmt.Println(v)
	}
	fmt.Println("over")
}
