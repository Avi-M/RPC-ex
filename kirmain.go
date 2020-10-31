package main

import "fmt"

func main() {
	mychnl := make(chan string)
	go func() {
		mychnl <- "GFG"
		mychnl <- "gfg"
		mychnl <- "Geeks"
		mychnl <- "GeeksforGeeks"
		close(mychnl)
	}()
	for res := range mychnl {
		fmt.Println(res)
	}
}
