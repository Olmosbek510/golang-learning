package main

import "fmt"

func main() {
	specialPrint(1, 233, 45)
}

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func makeItSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func describe(i interface{}) {
	fmt.Printf("Type = %T, Value: %v\n", i, i)
}

func specialPrint(nums ...int) error {
	if _, err := fmt.Println(nums); err != nil {
		return err
	}
	return nil
}
