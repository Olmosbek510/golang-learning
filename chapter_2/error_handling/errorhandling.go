package main

import (
	"errors"
	"fmt"
)

func main() {
	deferExample()
}

func divide() {
	var input, num int
	_, err := fmt.Scan(&input, &num)
	if err != nil {
		fmt.Println("Please check the format of the input. Int is expected")
	} else {
		if num == 0 {
			panic("division by zero!")
		}
		fmt.Println(input / num)
	}

}

func customError(a, b int) (int, error) {
	//You must call the divide function  which divides two numbers, but it returns not only the division result, but also error information.
	//In case of any error, you must output " error ", if there is no error - the result of the function .
	//The divide (a int, b int) (int, error) function receives two numbers to be divided and returns the result (int) and an error (error).
	//The main package has already been declared, and the necessary packages have already been imported!
	//Don't forget to count the two numbers to be divided (pass them to the function)!
	if b == 0 {
		return -1, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func deferExample() {
	defer finish()
	fmt.Println("Program has been started")
	fmt.Println("Program is executing")
}
func finish() {
	fmt.Println("Program has been finished")
}
