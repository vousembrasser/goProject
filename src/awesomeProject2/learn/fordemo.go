package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//for 的条件里不需要括号
//for的条件里可以省略除事条件，结束条件，递增表达式
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2 //lsb最底位
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
