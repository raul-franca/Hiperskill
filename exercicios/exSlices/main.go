package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

func solve(s []string, x string, i int) []string {
	slCopy := make([]string, 1)
	copy(slCopy, s)

	// Assign x to the element at the 'i' index of the 'slCopy' slice:
	slCopy[i] = x
	fmt.Println(slCopy)
	return slCopy // Return the copy of the slice here!
}

// DO NOT delete or modify the contents of the main function!
func main() {
	var i int
	var x string

	mySlice := make([]string, 1)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	mySlice[i], x = scanner.Text(), scanner.Text()

	newSlice := solve(mySlice, x, i)
	if !reflect.DeepEqual(mySlice, newSlice) {
		log.Fatal("You didn't assign 'x' to the element at the 'i' index of the slice!")
	}
	fmt.Println(newSlice)
}
