package main

import "fmt"

func main() {
	fmt.Printf("%v\n", returnSomething())
}

func returnSomething() string {
	return "Something1"
}

func add(a int, b int) int {
	return a + b
}
