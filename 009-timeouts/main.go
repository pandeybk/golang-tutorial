// https://gobyexample.com/timeouts

package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		chan1 <- "result 1"
	}()

	select {
	case res := <-chan1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout 1")
	}

	chan2 := make(chan string, 2)

	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "result 2"
	}()

	select {
	case res := <-chan2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout 2")
	}
}
