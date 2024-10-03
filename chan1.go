package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	//ch <- 1
	//close(ch)

	select {
	case v, ok := <-ch:
		fmt.Println(ok)
		fmt.Println(v)
	default:
		fmt.Println("GZ")
	}
}
