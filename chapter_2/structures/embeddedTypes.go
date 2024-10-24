package main

import "fmt"

type Person struct {
	Name string
}
type Android struct {
	Person
	Model string
}

func (person Person) Talk() {
	fmt.Println("Hello my name is: " + person.Name)
}

//In this lesson, we tried to imagine variables and functions that are already familiar to us as objects from real life. To consolidate the result, we offer you a small creative task.
//You need to implement a structure with properties-fields On , Ammo and Power , with types bool, int, int respectively. This structure should have methods: Shoot and RideBike, which do not take arguments, but return a bool value.
//If the value On == false, then both methods will return false.
//Shoot can only be done if Ammo is present (then Ammo is reduced by one, and the method returns true), if it is not present, the method will return false. The RideBike method works the same way, but it only depends on the Power property.
//To check that you have done everything correctly, you should create a pointer to an instance of this structure named testStruct in the main function , and the program will then check the result.
//You can't see the closing curly brace at the end of main (), but it's there.
//You don't need to declare the main package!
//Good luck!
