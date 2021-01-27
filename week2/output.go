package main

import "fmt"

func main() {
	name := "Maria"
	fmt.Println(name)

	//Data Types
	//Example of a String
	//Go have to be explicit about the data type
	var ourString string
	ourString = "THis is 5 more"
	fmt.Printf("Type: %T\n", ourString)

	//Example of an int
	var ourInt int
	ourInt = 12
	fmt.Printf("%d is a %T!\n", ourInt, ourInt)

	var ourBool bool
	ourBool = false
	fmt.Printf("OurBool a %T\n", ourBool)

	var ourFloat float32
	ourFloat = 32.32
	fmt.Printf("ourFloat is a %T\n", ourFloat)

	var ourSlice []string
	ourSlice = []string{"hello", "how", "why", "Goodbye"}
	fmt.Println(ourSlice)

	var ourRune rune
	ourRune = 'h'
	//Single quote is important
	fmt.Printf("rune type: %T\n", ourRune)
	fmt.Printf("rune value: %v\n", ourRune)

	// Coerce   55 + int("55")

}
