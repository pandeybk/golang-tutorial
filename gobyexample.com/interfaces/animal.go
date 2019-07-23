// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

type Cat struct {
}

type Llama struct {
}

type JavaProgrammer struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

func (c *Cat) Speak() string {
	return "Meow!"
}

func (l Llama) Speak() string {
	return "????"
}

func (j JavaProgrammer) Speak() string {
	return "Design Patterns!"
}

func AnimalSounds(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	dog := Dog{}
	cat := new(Cat)
	llama := Llama{}
	javaProgrammer := JavaProgrammer{}

	AnimalSounds(dog)
	AnimalSounds(cat)
	AnimalSounds(llama)
	AnimalSounds(javaProgrammer)

	animals := []Animal{Dog{}, new(Cat), Llama{}, JavaProgrammer{}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
