package main

// package main is a special package
// it allows Go to create an executable file

/*
This is a multi-line comment.
import keyword makes another package availablec for this .go "file".
import "fmt" lets you access fmt package's functionality here in this file.
*/

import "fmt"

// "func main" is special.
// Go has to know where to start
// func main() creates a starting point for Go
//
// When code is compiled, Go runtime will first run main() function
func main() {

	// after: import "fmt"
	// Println function of "fmt" package becomes available

	// To see 'Println' docs, type below command in shell:
	// go doc -src fmt Println

	// Println is just an exported function from "fmt" package

	// Exported = First Letter is uppercase

	fmt.Println("Hello World")
}
