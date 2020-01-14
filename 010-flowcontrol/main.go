package main

import "fmt"

func forloop() {
	var n int
	for n < 3 {
		fmt.Println("I am number: ", n)
		n++
	}
}

func main() {
	forloop()
}
