package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	//isCorrect("Быть или не быть.")
	//isPalindrome()
	firstSub()
}

func ExampleString() {
	// Let's create a string literal s whose value is "This is a string."
	// The line consists of 17 characters.
	var s string = "Это строка"

	// However, the length of the string len(s) will be 19 bytes, because Cyrillic characters used
	// take up 2 bytes, and space takes up 1 byte.
	fmt.Printf("Длина строки: %d байт\n", len(s))

	// Посмотрим как взять подстроку
	fmt.Printf("Напечатаем только второе слово в кавычках: \"%v\"\n", s[7:])

	/*
		Let's try to change something in the line:
		s[3] = 12
		Compilation error: cannot assign to s[3] because strings are immutable sequences.
	*/

	// change the line by creating a new one
	s = s + " Новая строка"
	fmt.Printf("%v\n", s)

	// Now let's iterate through this line
	for _, b := range s {
		fmt.Printf("%v ", b)
	}
	fmt.Print("\n")

	// Output:
	// String length: 19 bytes
	// Print only the second word in quotes: "string"
	// This is a line New line
	// 1069 1090 1086 32 1089 1090 1088 1086 1082 1072 32 1053 1086 1074 1072 1103 32 1089 1090 1088 1086 1082 1072
}

func BuildInMethods() {
	fmt.Println(
		// Is the substring contained in the string
		strings.Contains("test", "es"),
		// result: true

		// Number of substrings in a line
		strings.Count("test", "t"),
		// result: 2

		// Does the line start with a prefix
		strings.HasPrefix("test", "te"),
		// result: true

		// Does the string end with a suffix
		strings.HasSuffix("test", "st"),
		// result: true

		// Returns the starting index of a substring in a string, and if there is no occurrence, returns -1
		strings.Index("test", "e"),
		// result: 1

		// concatenates an array of strings via a symbol
		strings.Join([]string{"hello", "world"}, "-"),
		// result: "hello-world"

		// Repeats a line n times in a row
		strings.Repeat("a", 5),
		// result: "aaaaa"

		// The Replace function replaces any occurrence of old in your string with new
		// If n is -1, then all occurrences will be replaced.
		// General view: func Replace(s, old, new string, n int) string
		// Example:
		strings.Replace("blanotblanot", "not", "***", -1),
		// result: "bla***bla***"

		// Splits the string according to the delimiter
		strings.Split("a-b-c-d-e", "-"),
		// result: []string{"a","b","c","d","e"}

		// Returns a lowercase string
		strings.ToLower("TEST"),
		// result: "test"

		// Returns a string with upper case
		strings.ToUpper("test"),
		// result: "TEST"

		// Returns a string with the cut set
		strings.Trim("tetstet", "te"),
		// result: s
	)
}

func ExampleByteSlice() {
	bs := []byte("This is byte slice")
	fmt.Println(bs)

	fmt.Printf("Так байтовый срез выглядит внутри: %v\n", bs)

	// Продемонстрируем, что байтовый срез можно изменить,
	// а затем напечатаем его в виде строки
	for i := range bs {
		if bs[i]%2 == 0 {
			bs[i] = bs[i] + 1
			continue
		}
		bs[i] = bs[i] - 1
	}
	fmt.Printf("Измененный байтовый срез в виде строки: %s", bs)
}

func ExampleRune() {
	// Поступим также, как в работе с типом []byte
	rs := []rune("Это срез рун")

	var str = "Hello there"
	for _, elem := range []rune(str) {
		fmt.Printf("%s\n", string(elem))
	}
	// Итерируясь мы будем заменять символ 'р' на '*'
	for i := range rs {
		if rs[i] == 'р' {
			rs[i] = '*'
		}
	}
	//fmt.Printf("Измененнный срез в виде строки: %s\n", string(rs))

	// Output:
	// Измененнный срез в виде строки: Это с*ез *ун
}
func ExampleUnicode() {
	// функции ниже принимают на вход тип rune

	// проверка символа на цифру
	fmt.Println(unicode.IsDigit('1')) // true
	// проверка символа на букву
	fmt.Println(unicode.IsLetter('a')) // true
	// проверка символа на нижний регистр
	fmt.Println(unicode.IsLower('A')) // false
	// проверка символа на верхний регистр
	fmt.Println(unicode.IsUpper('A')) // true
	// проверка символа на пробел
	// пробел это не только ' ', но и:
	//  '\t', '\n', '\v', '\f', '\r' - подробнее читайте в документации
	fmt.Println(unicode.IsSpace('\t'))            // true
	fmt.Println("Is space", unicode.IsSpace(' ')) // true

	// С помощью функции Is можно проверять на кастомный RangeTable:
	// например, проверка на латиницу:
	fmt.Println(unicode.Is(unicode.Latin, 'ы')) // false

	// функции преобразований
	fmt.Println(string(unicode.ToLower('F'))) // f
	fmt.Println(string(unicode.ToUpper('f'))) // F
}

func isCorrect(s string) bool {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	for indx, elem := range text {
		if (indx == 0 && !unicode.IsUpper(elem)) || (indx == len(text)-1 && elem != '.') {
			return false
		}
	}
	return true
}

func isPalindrome() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	slice := []rune(text)
	right := len(text) - 1
	for i := 0; i < len(slice)/2; i++ {
		if text[i] != text[right] {
			fmt.Println("Wrong")
			return
		}
		right--
	}
	fmt.Println("Right")
}

func firstSub() {
	var X string
	var S string
	X, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	S, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	xrune := []rune(X)
	srune := []rune(S)
	sruneSize := len(srune)
	for i := 0; i < len(xrune); i++ {
		if i+sruneSize < len(xrune) {
			temp := srune[i : i+sruneSize-1]
			fmt.Println(temp)
		}
	}
	X = strings.TrimSpace(X)
	S = strings.TrimSpace(S)
	index := strings.Index(X, S)
	fmt.Println(index)
}
