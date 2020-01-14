package main

import "fmt"

func main() {
	fmt.Println("This is cool")
	a := 1
	if a == 0 {
		fmt.Println("I am zero")
	} else {
		fmt.Println("I am 1")
	}

	evenOrOdd(8)
	evenOrOdd(7)
	evenOrOdd(888)
}

func evenOrOdd(number int) {
	if number%2 == 0 {
		fmt.Println(number, " is even")
	} else {
		fmt.Println(number, " is odd")
	}
}
