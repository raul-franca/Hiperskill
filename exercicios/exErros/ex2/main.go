package main

import (
	"errors"
	"fmt"
	"os"
)

//Go's Standard library offers two functions to create errors: errors.New
//from the errors package and fmt.Errorf from the fmt package.
//The difference between them is that fmt.Errorf provides the ability
//to add formatting to our error message, which means we can pass a parameter
//to the fmt.Errorf function and include it in the error message.
//
//Suppose we have the divide function in our Go program,
//and we want to show an error message when the user tries to divide a number by zero.
//Let's take a look at the implementation of the error message when using theerrors.
//New and the fmt.Errorf functions:

// Implementation with errors.New
func divide(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return num1 / num2, nil
}

// Output: cannot divide by zero

// Implementation with fmt.Errorf
func divide2(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("cannot divide %.2f by %.2f", num1, num2)
	}
	return num1 / num2, nil
}

// Output: cannot divide 10.00 by 0.00
// The function openFile returns a custom error message if opening the file fails
func openFile(filename string) error {
	if _, err := os.ReadFile(filename); err != nil {
		return fmt.Errorf("error opening %s: %w", filename, err)
	}

	return nil
}
func main() {

}
