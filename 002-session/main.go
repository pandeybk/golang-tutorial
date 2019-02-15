package main

import (
	"fmt"

	"github.com/pandeybk/speaking-golang/002-session/printer"
)

func main() {
	t1 := printer.NewThing()
	fmt.Printf("Thing 1: %v\n", t1)
	runThis(t1)
	putPrint(t1, "cool")
}

func runThis(input printer.Printer) {
	res := input.Print()
	fmt.Printf("Thing 2: %v\n", res)
}

func putPrint(input printer.PrinterPutter, s string) {
	input.Put(s)
	fmt.Printf("Thing 1: %v\n", input.Print())
}
