package main

import "fmt"

func main() {

	outerData := "This is from main \n"

	f := func(s string) error {
		fmt.Printf(s)
		// outerData = "Now its owned by 'f'\n"
		fmt.Printf(outerData)
		outerData = "I will not print from 'f'\n"
		return nil
	}

	err := runThis(f)

	if err != nil {
		fmt.Printf("Got error: %v\n", err)
	}

	outerData = "This is the second run with main data\n"
	runThis(f)
}

func runThis(f func(string) error) error {
	s := "data inside runThis \n"
	return f(s)
}
