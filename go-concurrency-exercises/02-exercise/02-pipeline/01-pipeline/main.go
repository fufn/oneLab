package main

import "fmt"

// TODO: Build a Pipeline
// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int{
	send := make (chan int)

	go func() {
		for _,n := range (nums){
			send <- n
		}
		close(send)
	}()
	return send
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <- chan int) <-chan int {

	out := make (chan int)
	go func() {
		for n :=range in{
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// set up the pipeline
	for n := range square(square(generator(2, 3))) {
		fmt.Println(n)
	}
	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.

}
