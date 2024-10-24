package main

import "fmt"

func pointer() {
	a := 100
	b := &a
	fmt.Println(b)
	*b++

	var c *int = &a
	fmt.Println(c)
	fmt.Println(a)
}

func zero(x *int) {
	*x = 100
}

func changer(n *int) {
	*n = 2 * *n
}

func test(x1 *int, x2 *int) {
	res := *x1 * *x2
	fmt.Println(res)
}

func swap(a *int, b *int) {
	*a += *b
	*b = *a - *b
	*a -= *b
	fmt.Println(*a, *b)
}
