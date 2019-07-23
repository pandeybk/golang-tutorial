// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

package main

import "fmt"

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}
func main() {
	vals := []string{"demo", "cool", "not cool"}

	valsinter := make([]interface{}, len(vals))

	for i, v := range vals {
		valsinter[i] = v
	}
	PrintAll(valsinter)
}
