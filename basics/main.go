package main

import "fmt"

func main() {
	fmt.Println("\\")
	fmt.Println("\"")
	fmt.Println("123\f456\f789")
	var a int
	fmt.Printf("\a Input: ")
	fmt.Scan(&a)
}
