package main

import "fmt"

func main() {

	var s = "some string variable" // var s of the type string
	var p = &s                     // var p is a pointer to a string

	fmt.Println(s) // Printing the value
	fmt.Println(p) // Printing the address where the value is stored
	fmt.Println(*p)
	fmt.Println(&p)
	s = "Aurora "
	fmt.Println(*p)
}
