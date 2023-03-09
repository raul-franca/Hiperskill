package main

import (
	"fmt"
)

func main() {

	fmt.Printf("kitty1\nkitty2")
	fmt.Printf("\n")
	fmt.Printf("kitty1\tkitty2\n")
	// kitty1
	// kitty2
	// kitty1  kitty2

	// Using %f to print a float value with default width and precision
	fmt.Printf("%f\n", 20.66) // 20.660000

	// Using %f with the precision length 1
	fmt.Printf("%.2f \n", 20.33) // 20.3

	// The %c formatter is used to format a single character
	letter := 't'
	fmt.Printf("%c", letter) // t
	secret := '🤫'
	fmt.Printf("%c", secret) // 🤫

	// The %s formatter is great to format a string
	fmt.Printf("%s", "This is a random string\n") // This is a random string

	// The %t formatter is suitable for Boolean values
	fmt.Printf("%t \n", false) // false

	// Formatting the width of a string
	fmt.Printf("|%10s| \n", "pikachu") // | pikachu| (this action added one additional space where the string begins)

	a := "First variable"
	b := "Second variable"
	fmt.Printf("%[1]s | %[1]s \n", a) // First variable | First variable
	fmt.Printf("%[2]s | %[1]s", a, b) // Second variable | First variable
	fmt.Println(" ")
	fmt.Printf("%v", '😄') // 128516 (Unicode number)
	fmt.Println(" ")
	fmt.Printf("%v", "世界") // 世界 (means "world" in Chinese)
	fmt.Println(" ")
	fmt.Printf("%v", []int{1, 2, 3, 4, 5, 6}) // [1 2 3 4 5 6]
	fmt.Println(" ")
	fmt.Printf("%v", [3]int{9, 8, 7}) // [9 8 7]
	fmt.Println(" ")
	fmt.Printf("%v", 1+5i) // 1+5i (this is a complex number)
	fmt.Println(" ")
	first := fmt.Sprintf("%d", 500) // 'first' variable now has the value of 500
	fmt.Printf("'first' variable now has the value of 500: %v \n", first)
	fmt.Println(" ")
	binaryVariable := fmt.Sprintf("%b", 500) // 'binaryVariable' variable now has the value of 111110100
	fmt.Printf("'binaryVariable' variable now has the value of 111110100 =  %v \n", binaryVariable)
	fmt.Println(" ")

	s := fmt.Sprintln("if you try to print s,",
		"it will automatically print a new line at the end of the string")

	fmt.Print(s, "Newline")

	fmt.Println(" ")

	fmt.Printf(`%s 
and
the
brave
new\n
world\n`, "Go")

}
