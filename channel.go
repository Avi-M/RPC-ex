package main

import "fmt"

func myfunc(ch chan int) {
	fmt.Println(230 + <-ch)

}
func main() {
	fmt.Println("start main method")
	ch := make(chan int)

	go myfunc(ch)
	ch <- 23
	fmt.Println("close method")
}
