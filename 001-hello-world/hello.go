package main

// "fmt" package which is already return and use using import statment
import (
	"fmt"
)

// Working with go fmt package
func main() {
	fmt1()
	fmt2()
	fmt3()
}

func fmt1() {
	// Simple hello world showing
	fmt.Println("Hello world!")
}

func fmt2() {
	n, err := fmt.Println("This is number: ", 12, "and boolean: ", true, "And I can take any number of parameters")

	// You can't have any unused variables in your code
	fmt.Println("Number of bytes written: ", n)
	fmt.Println("Any write error encountered: ", err)
}

func fmt3() {
	// Throwing away returns using "_" underscores characters
	n, _ := fmt.Println("Throw away return values using _ underscore charatcters")
	fmt.Println(n)
}
