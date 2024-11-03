package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	durationExplanation()
	//practicalExamples()
}

func getMiddle(a, b, c float64) float64 {
	maximum := math.Max(a, math.Max(b, c))
	minimum := math.Min(a, math.Min(b, c))
	sum := a + b + c
	return sum - (maximum + minimum)
}

func pdpTask(n int) int {
	var result int
	isEven := n%2 == 0
	if isEven {
		result = 0
	} else {
		result = 1
	}
	for n > 0 {
		if isEven {
			result += n % 10
		} else {
			result *= n % 10
		}
		n /= 10
	}
	return result
}
func customTime() {
	customTime := time.Date(
		2024,         // year
		time.October, // month
		30,           // day
		23,           // hour
		20,           // minute
		12,           // second
		0,            // nanosecond
		time.UTC,     // location
	)
	fmt.Println("Time: ", customTime)
}

func unixTimeExample() {
	unixTime := time.Unix(150000, 0) // seconds, nanoseconds
	fmt.Println("Unix time: ", unixTime)
}

func formattingTime() {
	now := time.Now()
	println(now.Format("02 2006 06"))
	//	Common formatting examples
	//now.Format("2006-01-02 15:04:05") → 2023-10-25 14:30:15
	//now.Format("02-01-2006") → 25-10-2023
	//now.Format("3:04 PM") → 2:30 PM
	//now.Format("02 Jan 06") → 25 Oct 23
	currentTime := time.Date(2024, time.January, 01, 20, 19, 10, 0, time.UTC)
	fmt.Println(currentTime.Format("01-01-2002 03:04:05"))
	print(now.Format("02-01-2006 15:04:05"))
}

func stringsToTime() {

	//	time.Parse()
	timeString := "2024/09/16 18-17"
	parsedTime, err := time.Parse("2006/01/02 02-04", timeString)
	if err != nil {
		fmt.Println("Error parsing time: ", err)
	} else {
		fmt.Println("Parsed time: ", parsedTime)
	}

	// time.ParseInLocation()

	loc, err := time.LoadLocation("Asia/Samarkand")
	if err != nil {
		panic(err)
	}

	timeString = "May 15 20 05:45:10pm"
	parsedTime, err = time.ParseInLocation("Jan 2 06 03:04:05pm", timeString, loc)
	if err != nil {
		fmt.Println("Error parsing time: ", err)
	} else {
		fmt.Println("Parsed time with location: ", parsedTime)
	}
}

func practicalExamples() {
	//	1.Get current time and format it
	now := time.Now()
	fmt.Println("Current time: ", now.Format("2024-01-01 2024 13:04:05"))

	//	2. Parse a String into a Time object
	timeString := "2024-4-01 19:17:09"
	// The layout should match the exact format of timeString
	parsedTime, err := time.Parse("2006-1-02 15:04:05", timeString)
	if err != nil {
		_ = fmt.Errorf("error parsing time: %s", timeString)
		panic(err)
	}
	fmt.Println("Parsed time", parsedTime)
	//	3. Parse a string into Time object with a specific time zone
	timeString = "Jan 2025 10 18:09:10"
	loc, err := time.LoadLocation("Asia/Japan")
	parsedTime, err = time.ParseInLocation("Jan 2024 01 01:01:01", timeString, loc)
	if err != nil {
		fmt.Println("Error parsing time with specific location:", err)
	}
	fmt.Println("Parsed time: ", parsedTime)
}

// useful method for working with time
func currentPractice() {
	current := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)

	// func (t Time) Date() (year int, month Month, day int)
	fmt.Println(current.Date()) // 2020 May 15

	// func (t Time) Year() int
	fmt.Println(current.Year()) // 2020

	// func (t Time) Month() Month
	fmt.Println(current.Month()) // May

	// func (t Time) Day() int
	fmt.Println(current.Date()) // 15

	// func (t Time) Clock() Clock
	fmt.Println(current.Clock()) // 17 45 12

	// func (t Time) Hour() int
	fmt.Println(current.Hour()) // 17

	// func (t Time) Minute() int
	fmt.Println(current.Minute()) // 45

	// func (t Time) Second() int
	fmt.Println(current.Second()) // 12

	// func (t Time) Unix() int64
	fmt.Println(current.Unix()) // 1589546712 seconds from unix's start time in the early 1970s up to the specified time in 'current'

	// func (t Time) Weekday() Weekday
	fmt.Println(current.Weekday()) // Friday

	// func (t Time) YearDay() int
	fmt.Println(current.YearDay()) // 136
}

// converting a time structure to String
func convertingTimeToString() {
	current := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
	fmt.Println(current.Format("02-01-2006 15:04:05")) // 15-05-2020 17:45:12
}

// Comparison of Tome structures

func comparisonTimeStructures() {
	/* There are three main comparison operations in time :
	-After(u Time) bool
	-Before(u Time) bool
	-Equal(u Time) bool
	*/
	firstTime := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
	secondTime := time.Date(2020, time.May, 15, 16, 45, 12, 0, time.Local)

	fmt.Println(firstTime.After(secondTime)) // true

	fmt.Println(firstTime.Before(secondTime)) // false

	fmt.Println(firstTime.Equal(secondTime)) // false
}

// Methods for changing the time structure

func methodsForChangeTimeStructure() {
	now := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
	// func (t Time) Add(d Duration) Time
	// изменяет дату в соответствии с параметром - "продолжительностью"
	future := now.Add(time.Hour * 12) // move forward 12 hours
	// func (t Time) AddDate(years int, months int, day int) Time
	// changes the date in accordance with the parameters - the number of years, months and days
	past := now.AddDate(-1, -2, -3) // move 1 year two months and 3 days back
	// func (t Time) Sub(u Time) Duration
	// calculates the time elapsed between two dates

	fmt.Println(future.Sub(past)) // // 10332h0m0s

	/*
		Note that negative values can also be used in the Add and AddDate methods,
		which allows you to no only 'add' time (as can bee seen from the method names), but also 'subtract' it.
	*/
}

func task() {
	var timeString string
	_, err := fmt.Scan(&timeString)
	if err != nil {
		return
	}
	parsedTime, _ := time.Parse("2006-01-02T15:04:05-07:00", timeString)
	fmt.Println(parsedTime.Format(time.UnixDate))
}

func task1() {
	bf := bufio.NewReader(os.Stdin)
	stringTime, _ := bf.ReadString('\n')
	stringTime = strings.TrimSuffix(stringTime, "\n")
	layout := "2006-01-02 15:04:05"
	timeVal, _ := time.Parse(layout, stringTime)
	testTime := time.Date(timeVal.Year(), timeVal.Month(), timeVal.Day(), 13, timeVal.Minute(), timeVal.Second(), timeVal.Nanosecond(), time.Local)
	if timeVal.After(testTime) {
		timeVal = timeVal.AddDate(0, 0, 1)
	}
	fmt.Println(timeVal.Format(layout))
}

const now = 1589570165

// type Duration explanation
func durationExplanation() {
	now := time.Now()
	past := now.AddDate(0, 0, -1)
	future := now.AddDate(0, 0, 1)

	// func Since(t Time) Duration
	// calculates the period between the current moment and a given time in the past
	fmt.Println(time.Since(past).Round(time.Second))

	// func Until(t Time) Duration
	// calculates the period between the current moment and a given time in the future
	fmt.Println(time.Until(future).Round(time.Second))

	// func ParseDuration(s string) (Duration, error)
	// converts the string to Duration using annotations:
	// "ns" - nanoseconds,
	// "us" - microseconds,
	// "ms" - milliseconds,
	// "s" - seconds,
	// "m" - minutes,
	// "h" - hours.

	dur, err := time.ParseDuration("1h12m3s90ns")
	if err != nil {
		panic(err)
	}
	fmt.Println(dur.Round(time.Hour).Hours())
}

func task2() {
	bf := bufio.NewReader(os.Stdin)
	strTime, _ := bf.ReadString('\n')
	layout := "02.01.2006 15:04:05"
	strings.TrimPrefix(strTime, "\n")
	sliceTimes := strings.Split(strTime, ",")
	smaller, _ := time.Parse(layout, strings.TrimSpace(sliceTimes[0]))
	larger, _ := time.Parse(layout, strings.TrimSpace(sliceTimes[1]))

	if smaller.After(larger) {
		temp := smaller
		smaller = larger
		larger = temp
	}
	fmt.Println(larger.Sub(smaller))
}
