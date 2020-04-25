package main

import (
	"fmt"
	"time"
)

func main() {
	//http.HandleFunc("/", func(
	//	writer http.ResponseWriter,
	//	request *http.Request) {
	//	_, _ = fmt.Fprintf(writer, "<h1>hello world! %s</h1>", request.FormValue("name"))
	//})
	//_ = http.ListenAndServe(":8888", nil)

	var a = []int{1, 2, 3, 4, 5}

	var out = make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()

	go func() {
		//var out1= make(chan int)
		var v1, ok1 = <-out
		fmt.Println(ok1)
		for ok1 {
			fmt.Println(ok1)
		}
	}()

	time.Sleep(time.Second)
}
