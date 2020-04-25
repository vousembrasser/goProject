package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//return a / b
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operator: %s" + op) //panic 是用来报错
	}
}

//带余数除法 13/2 =6·····1
func div(a, b int) (quotient int, remainder int) {
	//return a / b, a % b
	quotient = a / b
	remainder = a % b
	return quotient, remainder
}

func div1(a, b int) (quotient int) {
	return a / b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling funtion %s with args"+
		"(%d, %d)", opName, a, b)
	return op(a, b)
}
