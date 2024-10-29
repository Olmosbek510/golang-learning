package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

func main() {
	passingAsArguments()
	var i interface{} = "1"
	fmt.Printf("%T\n", i)
	_, isCastable := i.(int)
	fmt.Println("IsCastable: ", isCastable)
	if v, ok := i.(int); ok {
		fmt.Println(v + 12)
	}

	typeSwitch("Olmos knows")
	typeSwitch(21)

}

func passingAsArguments() {
	src := "olmosbek"
	result := strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return r - 'a' + 'A'
		} else if r >= 'A' && r <= 'Z' {
			return r - 'A' + 'a'
		}
		return r
	}, src)
	fmt.Println(result)
	sumOfDigits := func(num int) int {
		sum := 0
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		return sum
	}
	number := sumOfDigits(123)
	fmt.Println(number)
}
func invert(r rune) rune {
	if unicode.IsLower(r) {
		return unicode.ToUpper(r)
	}
	return unicode.ToLower(r)
}

func getInverter() func(rune) rune {
	return invert
}

func typeSwitch(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Println("Multiplication to 2: ", v*2)
	case string:
		fmt.Println(v + " golang")
	default:
		fmt.Println("I dont know the type of ", v)
	}
}

func task() {
	var value1, value2, operation interface{}
	fmt.Scan(&value1, &value2, &operation)
	if _, ok := value1.(float64); !ok {
		fmt.Printf("value=%v: %s\n", value1, reflect.TypeOf(value1).Kind())
		return
	}
	if _, ok := value2.(float64); !ok {
		fmt.Printf("value=%v: %s\n", value2, reflect.TypeOf(value2).Kind())
		return
	}
	var isOper = func(operation string) bool {
		if operation == "+" || operation == "-" || operation == "/" || operation == "*" {
			return true
		}
		return false
	}
	if value, ok := operation.(string); !ok || !isOper(value) {
		fmt.Println("неизвестная операция")
		return
	}
	oper, _ := operation.(string)
	val1 := value1.(float64)
	val2 := value2.(float64)
	var res float64
	switch oper {
	case "+":
		res = val1 + val2
	case "-":
		res = val1 - val2
	case "*":
		res = val1 * val2
	case "/":
		res = val1 / val2
	}
	fmt.Printf("%.4f", res)
}
