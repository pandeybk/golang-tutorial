package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 1; j < 7; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n < 9; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	println(sumIt(3, 5))
	println(sumIt(5, 3))

}

func sumIt(a, b int) int {

	if a > b {
		a, b = b, a
	}

	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	return sum
}
