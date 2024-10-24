package basics

import (
	"fmt"
	"math"
)

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

	// Comparing arrays
	e := [3]int{1, 2, 3}
	f := [3]int{1, 2, 3}

	fmt.Println(e == f)

	// Indexing in array

	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("arr[%d] = %d\n", i, arr1[i])
	}

	for indx, elem := range arr1 {
		fmt.Printf("index: %d, element: %d\n", indx, elem)
	}

	for indx := range arr1 {
		fmt.Printf("%d\n", arr1[indx])
	}
}

func workingArrayTask() {
	var workingArray [10]uint8 = [10]uint8{}
	var indexes [6]uint8 = [6]uint8{}

	for indx := 0; indx < len(workingArray); indx++ {
		fmt.Scan(&workingArray[indx])
	}

	for indx := 0; indx < len(indexes); indx++ {
		fmt.Scan(&indexes[indx])
	}

	for i := 0; i < len(indexes); i += 2 {
		temp := workingArray[indexes[i]]
		workingArray[indexes[i]] = workingArray[indexes[i+1]]
		workingArray[indexes[i+1]] = temp
	}

	for indx := 0; indx < len(workingArray); indx++ {
		fmt.Print(workingArray[indx])
		fmt.Print(" ")
	}
	fmt.Println("type ok")
}

func slices() {
	var a []int
	var b []int = []int{1, 2, 345, 4, 5}
	c := []int{1, 2, 3, 4, 5, 6, 7}
	d := []int{1, 12}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	slice := make([]int, 10, 19)
	fmt.Println(slice)
}

func sliceOperator() {
	var users = [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}

	users1 := users[2:6] //from 3rd to 6th
	users2 := users[:4]  // from 1 to 4
	users3 := users[3:]  // from 4 to the end

	fmt.Println(users1) // Kate, Sam, Tom, Paul
	fmt.Println(users2) // Bob, Alice, Kate, Sam
	fmt.Println(users3) // Sam, Tom, Paul, Mike, Robert
}

func buildInFunctionsSlices() {
	// func append() adds elements to slice
	a := []int{1, 2, 3, 4}
	a = append(a, 5, 6, 7, 8)
	fmt.Println(a)

	baseArray := [4]int{1, 4, 5, 6}
	baseSlice := baseArray[1:3]
	baseSlice = append(baseSlice, 3, 4, 5, 6)
	fmt.Printf("Base array length: %d, elements: %v\n", len(baseArray), baseArray)
	fmt.Printf("Slice length: %d, capacity: %d, elements: %v\n", len(baseSlice), cap(baseSlice), baseSlice)

	//	removing element from the slice
	a = append(a[0:1], a[2:]...)
	fmt.Printf("Slice after deletion of a single element: %v", a)

	//	function copy    syntax: func copy(dst, src []Type) int
	a = []int{1, 2, 4, 5, 6}
	b := make([]int, 5, 9)

	copy(b, a)
	fmt.Printf("\na: %v\nb: %v", a, b)
}

func task() {
	var size int
	fmt.Scan(&size)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
		if i%2 == 0 {
			fmt.Print(arr[i])
			if i != size-1 {
				fmt.Print(" ")
			}
		}
	}
}

func minCount() {
	count := 1
	min := math.MaxInt
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var temp int
		fmt.Scan(&temp)
		if min > temp {
			min = temp
			count = 1
		}
		if temp == min {
			count++
		}
	}
	fmt.Println(count)
}

func ignoreDigit() {
	var num, miss string
	fmt.Scan(&num, &miss)

	// Ensure that 'miss' is a single character
	if len(miss) != 1 {
		fmt.Println("Error: 'miss' should be a single character")
		return
	}

	res := ""
	for i := 0; i < len(num); i++ {
		if string(num[i]) == miss {
			continue
		}
		res += string(num[i])
	}
	fmt.Println(res)
}
