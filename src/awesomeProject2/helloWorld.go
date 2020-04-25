package main

import "fmt"

//函数外 不可使用 := 来定义变量  aa := 3
//包内部变量
//var aa = 3
//var ss = "kk"
//多个变量 定义时可以采用以下形式定义
var (
	aa = 3
	ss = "kk"
)

func main() {
	fmt.Println("hello world !")
	variableZeroValue()
}

//变量赋值操作
func variableInitialValue() {
	var a, b int = 3, 4
	var c, d = true, "def"
	var s string = "abc"
	fmt.Println(a, b, s, c, d)
}

//去类型，go语言会自动识别
func variableTypeDeduction() {
	var a, b = 1, 2
	var s = "abc"
	fmt.Println(a, b, s)
}

//去除var 使用:=替代
func variableShorter() {
	a, b := 1, 2
	b = 5
	fmt.Println(a, b)
}

func variableZeroValue() {
	var a int
	var s string
	fmt.Println(a, s)
	fmt.Printf("%d, %q\n", a, s)
	fmt.Println(aa, ss)
}
