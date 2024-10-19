package main

import "fmt"

func arrays() {
	var arr [3]int
	fmt.Println(arr)

	// Declare an array 'a' with a fixed size of 3 elements,
	// initialized with the values 1, 2, and 3.
	var a [3]int = [3]int{1, 2, 3}

	// Declare and initialize array 'b' using shorthand notation,
	// also with a fixed size of 3 elements initialized with values 1, 2, and 3.
	b := [3]int{1, 2, 3}

	// Declare and initialize array 'c' with a variable size
	// (using the '...' syntax) that automatically determines
	// the length based on the number of elements provided,
	// which is 6 elements: 1, 2, 3, 4, 5, and 6.
	c := [...]int{1, 2, 3, 4, 5, 6}

	// Declare and initialize array 'd' with a fixed size of 3 elements,
	// where the second element (index 1) is explicitly set to 12,
	// and the first and third elements are initialized to the zero value (0).
	d := [3]int{1: 12}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
