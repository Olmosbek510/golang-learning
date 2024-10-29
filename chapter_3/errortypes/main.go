package main

import (
	"fmt"
	"unicode"
)

type customError uint

func (c customError) Error() string {
	return fmt.Sprintf("digit, index %d", c)
}

func errorInString(str string) error {
	// Useful string work ignored
	for i, s := range str {
		if unicode.IsDigit(s) {
			return customError(i)
		}
	}
	return nil
}

type Battery struct {
	charge string
}

func (b Battery) String() string {
	var res []rune
	count := 0
	res = append(res, '[')
	for _, ch := range b.charge {
		if ch == '1' {
			count++
		}
	}
	start := len(b.charge) - count - 1
	for indx := range b.charge {
		if indx > start {
			res = append(res, 'X')
		} else {
			res = append(res, ' ')
		}
	}
	res = append(res, ']')
	return string(res)
}

func main() {
	var input string
	fmt.Scan(&input)
	batteryForText := Battery{input}
	fmt.Println(batteryForText)
	//err := errorInString("string1string")
	//if err != nil {
	//	fmt.Printf("Error processed: %v\n", err)
	//}
	//var cError customError
	//if errors.As(err, &cError) {
	//	fmt.Printf("Context: %d\n", cError)
	//}
	// Output:
	// Error processed: digit, index 6
	// Context: 6
}
