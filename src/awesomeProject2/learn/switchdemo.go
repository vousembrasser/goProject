package main

import "fmt"

//switc会自动break 除非使用fallthrough
func switchdemo(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op) //panic 是用来报错
	}
	return result
}

func grade(score int) string {
	var g = ""
	//switch后面可以没有表达式 再case中声明表达式即可
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score:%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}
