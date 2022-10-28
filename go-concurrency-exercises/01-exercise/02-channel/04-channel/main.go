package main

import "fmt"

// TODO: Implement relaying of message with Channel Direction

func genMsg(send chan<-string) {
	// send message on ch1
	send <- "msg"
}

func relayMsg(recv <- chan string, send chan<-string) {
	// recv message on ch1
	msg := <-recv
	msg = msg + " relay"
	// send it on ch2
	send <- msg

}

func main() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)
	// spine goroutine genMsg and relayMsg
	go relayMsg(ch1, ch2)
	go genMsg(ch1)
	// recv message on ch2
	msg := <- ch2
	fmt.Println(msg)
}
