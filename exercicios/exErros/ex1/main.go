package main

import (
	"log"
	"os"
)

func main() {
	// Try to open a non-existing file: "new_file.txt"
	file, err := os.Open("new_file.txt") // os.Open returns two values: the file and an error
	if err != nil {
		log.Fatal(err) // Log the error & exit the program with the exit code 1 - meaning: general error
	}
	defer file.Close() // This line closes the file before exiting the program

	// Output:
	// 2022/10/04 06:09:55 open new_file.txt: The system cannot find the file specified.

}
