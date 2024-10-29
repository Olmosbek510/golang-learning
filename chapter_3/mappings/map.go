package main

import "fmt"

func main() {
	isExists()
}

func isExists() {
	fruits := map[string]int{
		"apple":  10,
		"banana": 11,
		"peach":  8,
	}
	for key, val := range fruits {
		fmt.Printf("Key: %s, Value: %v\n", key, val)
	}

	val, ok := fruits["apple"]
	fmt.Println(ok)
	fmt.Println(val)
	fmt.Println(len(fruits))
}
