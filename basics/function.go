package basics

func Add(a, b int, firstname, lastname string) (int, string) {
	z := a + b
	var fullname string = firstname + " " + lastname
	return z, fullname
}

func sumInt(nums ...int) (int, int) {
	count := 0
	sum := 0
	for _, elem := range nums {
		sum += elem
		count++
	}
	return count, sum
}
