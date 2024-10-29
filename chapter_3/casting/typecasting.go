package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	//parseFloat()
	println(adding("%^80 hhhhh20&&&&nd\n"))
}

func bytesToString() {
	a := "str"
	b := []byte(a)
	c := string(b)
	fmt.Println(a)

	fmt.Println(b)

	fmt.Println(c)
}

func runesToString() {
	str := "Olmosbek"
	r := []rune(str)
	b := []byte(str)
	fmt.Printf("rune: %v\nbyte: %v\nstr: %s", r, b, string(r))
}

func stringsToOtherTypes() {
	foo := "10"
	bar := "15"
	err := "olmos"
	_, error := strconv.Atoi(err)
	if error != nil {
		fmt.Println(error)
	}
	//baz := foo - bar invalid operation: foo - bar (operator - not defined on string)
	fmt.Printf("%T", foo)
	barInt, errBar := strconv.Atoi(bar)
	fooInt, errFoo := strconv.Atoi(foo)
	if errFoo == nil && errBar == nil {
		fmt.Println(barInt - fooInt)
	}
}

func parseFloat() {
	s := "23.23456"
	// как и в прошлом шаге, здесь 2 параметр - bitSize
	// bitSize - 32 или 64 (32 для float32, 64 для float64)
	// но нужно понимать что метод все равно вернет float64
	result, err := strconv.ParseFloat(s, 39)

	if err != nil {
		panic(err)
	}
	fmt.Println(result)         // 23.23456k,m
	fmt.Printf("%T \n", result) // float64

	// Конкретный пример для разных bitSize:
	s = "1.0000000012345678"
	//  не будем обрабатывать ошибки в примерах, но на практике так не делайте ;)
	result32, _ := strconv.ParseFloat(s, 32)
	result64, _ := strconv.ParseFloat(s, 64)
	fmt.Println("bitSize32:", result32) // вывод 1 (не уместились)
	fmt.Println("bitSize64:", result64) // вывод  1.0000000012345678
}

func adding(str string) int64 {
	temp := ""
	res := 0
	for _, elem := range str {
		if unicode.IsDigit(elem) {
			temp += string(elem)
		} else {
			val, _ := strconv.Atoi(temp)
			res += val
			temp = ""
		}
	}
	return int64(res)
}
