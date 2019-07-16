package main

import "fmt"

func main() {
	message := make(chan string)

	go func() {
		message <- "ping"
		message <- "cool"
		message <- "not"
	}()

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)

	buffered()
}

// Channel Buffering
func buffered() {
	message := make(chan string, 3)
	message <- "Hello"
	message <- "World"
	message <- "!!"

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}
