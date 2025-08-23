package main

import "fmt"

// middleName := "Cane"
var middleName = "Cane"

func main() {
	// var age int
	// var name1 string = "John"
	// var name2 = "Jane"

	// count := 15

	fmt.Println("Test scope", &middleName, middleName)

	lastName := "Smith"
	middleName := "Pit"
	fmt.Println("Test scope", &middleName, middleName)

	printName(lastName)
	printName(middleName)
}

func printName(s string) {
	myPrint := "Print string -"
	fmt.Println(myPrint, s)
}
