// https://gobyexample.com/interfaces

package main

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures

// Basic interface for geometric shapes which is holding area() and perim() method.
type geometry interface {
	area() float64
	perim() float64
}

// we have 2 types of geometry reactangle and circle
type rectangle struct {
	weidth, height float64
}

type circle struct {
	radius float64
}

// To implement and interface in go, we just need to implement all the methods in the interface.
// Here we implement geometry and reacts.
func (rectangle rectangle) area() float64 {
	return rectangle.height * rectangle.weidth
}

func (rectangle rectangle) perim() float64 {
	return 2*rectangle.height + 2*rectangle.weidth
}

// The implementes for circles
func (circle circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (circle circle) perim() float64 {
	return 2 * math.Pi * circle.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	c := circle{radius: 5}
	r := rectangle{weidth: 5, height: 10}

	measure(c)
	measure(r)
}
