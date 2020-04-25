package main

import (
	"fmt"
	"sort"
)

func main() {
	//create of slice of int
	//并不是数组而是slice
	a := []int{1, 7, 4, 7, 3, 8, 9, 345, 8, 9}
	sort.Ints(a)
	//i是序号
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Println("---------------------")
	for _, v := range a {
		fmt.Println(v)
	}
}
