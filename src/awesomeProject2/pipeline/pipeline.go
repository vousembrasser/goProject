package pipeline

//可变长参数 int
//<- 使用者只能取
func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	//channel 是goroutine 和 goroutine 之间的通讯途径
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		//Read into memory
		var a []int
		for v := range in {
			a = append(a, v)
			out <- v
		}

		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

//func Merge(in1, in2 <-chan int) <-chan int{
//	out := make(chan int)
//	go func(){
//		v1 ,ok1 := <- in1
//		v2, ok2 := <- in2
//		for ok1 || ok2{
//			if !ok2 || (ok1 && v1 <= v2){
//				out <- v1
//			}
//		}
//	}
//}
