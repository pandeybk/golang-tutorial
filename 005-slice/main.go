package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Declare a list/slice of string pointers
	listOfNumberStrings := []*string{}

	// Forward declaration of the variable we will store the data in before adding to slice

	// Loop from 0 to 9
	for i := 0; i < 10; i++ {
		var numberString string
		// Format the string with '#' prefixed to the number
		numberString = fmt.Sprintf("#%s", strconv.Itoa(i))
		// Add number string to the slice
		listOfNumberStrings = append(listOfNumberStrings, &numberString)
	}

	for _, n := range listOfNumberStrings {
		fmt.Printf("%s\n", *n)
		fmt.Printf("%s\n", n)
	}
}
