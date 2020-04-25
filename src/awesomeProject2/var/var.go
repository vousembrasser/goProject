package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/**
	内建变量类型
	bool, string
 	(u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
		加u为无法符号整数，go语言中没有long类型，
		int不指定长度会更具操作系统来默认给长度 32位系统就是32，64位系统就是64
		ptr为指针
	byte，rune
 		rune为go语言中的char类型，但不完全相同 runes长度32位
	float32, float64, complex64, complex128
  		complex为复数 数学中的i虚数
*/

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	// 这里的1i 为 虚数i 直接写i无法识别成虚数i 所以写成1i的形式
	// e^pi * i + 1 = 0 的写法第一种
	fmt.Println(cmplx.Pow(math.E, math.Pi*1i) + 1) //这里输出的不是0 是应为浮点数的原因
	// e^pi * i + 1 = 0 的写法第二种
	fmt.Println(cmplx.Exp(math.Pi*1i) + 1)
}

//枚举的定义
func enums() {
	const (
		cpp = iota
		_
		java
		python
		goLong
		javascript
	)

	const (
		k = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, javascript, python, goLong)
	fmt.Println(k, kb, mb, gb, tb, pb)
}

func main() {
	euler()
	enums()
}
