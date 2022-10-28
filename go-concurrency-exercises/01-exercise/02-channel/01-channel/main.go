package main

import "fmt"

func main() {
	ch := make(chan int)

	go func(a, b int) {
		c := a + b
		ch <- c
		close(ch)
	}(1, 2)
	// TODO: get the value computed from goroutine
	c := <-ch
	fmt.Printf("computed value %v\n", c)
}
