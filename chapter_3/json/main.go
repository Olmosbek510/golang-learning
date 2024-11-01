package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type person struct {
	Name   string
	Age    int
	Status bool
	Values []int
}

type user struct {
	Name     string
	Email    string
	Status   bool
	Language []byte
}

func main() {
	lastTask()
}

//	Marshalling data(converting to json)

func typeToJsonTest() {
	s := person{
		Name:   "Olmosbek",
		Age:    20,
		Status: true,
		Values: []int{14, 54, 34},
	}

	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
	fmt.Printf("%s", data)
}

func typeToJsonPrettyPrinting() {
	p := person{
		Name:   "John",
		Age:    10,
		Status: false,
		Values: []int{23, 32, 45, 23},
	}

	data, err := json.MarshalIndent(p, "", "\t")
	if err == nil {
		fmt.Println("Error encoding JSON: ", err)
	}
	fmt.Printf("%s\n", data)
}

// UnMarshalling data(converting json data into go structs)
func jsonToType() {
	var p person
	s := "{\n        \"Name\": \"John\",\n        \"Age\": 10,\n        \"Status\": false,\n        \"Values\": [\n                23,\n                32,\n                45,\n                23\n        ]\n}\n"
	if err := json.Unmarshal([]byte(s), &p); err != nil {
		fmt.Println("Error while decoding JSON int person: ", err)
	}
	fmt.Printf("%v", p)
}

// Checking json validity
func checkToValidity() {
	m := user{
		Name:  "John Connor",
		Email: "test email",
	}

	data, _ := json.Marshal(m)
	data = data[1:]
	if !json.Valid(data) {
		fmt.Println(string(data))
		fmt.Println("Invalid json")
	} else {
		fmt.Println("Valid Json")
	}
}

// interfaces for unknown json structure
func interfaceForUnknownType() {
	data := []byte(`
	{
  "Name": "Sarah Connor",
  "Age": 29,
  "Attributes": {
    "Status": true,
    "Values": [5, 10, 15]
  }
}
`)
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("Error decoding JSON: ", err)
		return
	}
	name := result["Name"].(string)
	age := result["Age"].(float64)
	attributes := result["Attributes"].(map[string]interface{})
	status := attributes["Status"].(bool)
	values := attributes["Values"].([]interface{})

	fmt.Printf(`
	Name: %s,
	Age: %d
	Attributes:
		Status: %v
		Values: %v
	`, name, int(age), status, values)
}

// struct tags for customizing JSON
type myStruct struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	IsActive bool   `json:"isActive,omitempty"`
}

type group struct {
	ID       int
	Number   string
	Year     int
	Students []student
}
type student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}
type Result struct {
	AverageNumberOfGrades float64 `json:"AverageNumberOfGrades"`
}

func customizingJsonFields() {
	s := myStruct{
		Name:     "Olmosbek",
		Age:      19,
		IsActive: true,
	}
	data, err := json.MarshalIndent(s, "", "\t")
	if err == nil {
		fmt.Println(string(data))
	}
}

func task() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		// ...
	}
	var g group
	if err := json.Unmarshal(data, &g); err != nil {
		fmt.Println("Error while unmarshalling")
	}
	var count float64 = 0
	var temp float64 = 0
	for _, elem := range g.Students {
		for _, rate := range elem.Rating {
			temp += float64(rate)
			count++
		}
	}
	average := temp / count
	s, err := json.MarshalIndent(Result{AverageNumberOfGrades: average}, " ", "\t")
	fmt.Println(string(s))
}

type Report struct {
	GlobalId uint64 `json:"global_id"`
}

func lastTask() {
	path := "/Users/orazboyevolmosbek/Downloads/data-20190514T0100.json"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error while opening file")
		return
	}
	var reports []Report
	data, _ := io.ReadAll(file)
	if err := json.Unmarshal(data, &reports); err != nil {
		fmt.Println("Error")
		return
	}
	var sum uint64 = 0
	for _, elem := range reports {
		sum += elem.GlobalId
	}
	fmt.Println(sum)
}
