package main

import "fmt"

func foo(a int) {
    if (a > 0) {
        fmt.Println("I am feeling lucky")
    } else {
        fmt.Println("3" + 5)
    }
}

func main() {
    foo(1)
}
