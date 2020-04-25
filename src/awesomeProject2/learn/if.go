package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	const filename = "abc.txt"

	var contents, err = ioutil.ReadFile(filename)

	//------写法1
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println("---------------------")

	//------写法2   if的条件里可以赋值 变量只在if中生效 出了if不能使用
	if contents1, err1 := ioutil.ReadFile(filename); err1 != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%s\n", contents1)
	}

	fmt.Println(switchdemo(1, 2, "+"))

	fmt.Println(
		grade(1),
		grade(61),
		grade(80),
		grade(100),
	)

	fmt.Println(
		convertToBin(13),
	)

	printFile("abc.txt")

	fmt.Println(apply(div1, 13, 4))

	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(strconv.Itoa(a) + "  " + strconv.Itoa(b))

	a, b = swap1(a, b)
	fmt.Println(strconv.Itoa(a) + "  " + strconv.Itoa(b))

	s := []int{1, 2}
	fmt.Println(sliceDemo(s))
	s[0] = 100
	fmt.Println(s)

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6] //[2,3,4,5]
	s2 := s1[3:5]
	//fmt.Println(s1[3])
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1))

	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2))

}
