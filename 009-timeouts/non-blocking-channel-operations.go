package main

import "fmt"

// https://gobyexample.com/non-blocking-channel-operations

func main() {
	messages := make(chan string, 1)
	signals := make(chan string)

	select {
	case msg := <-messages:
		fmt.Println("Received", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("received:", msg)
	default:
		fmt.Println("No message sent")
	}

	select {
	case messages <- msg:
		fmt.Println("Received", msg)
	case sig := <-signals:
		fmt.Println("Signals: ", sig)
	default:
		fmt.Println("Nothing yet")
	}
}
