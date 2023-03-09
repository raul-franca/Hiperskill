package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n, s int
	fmt.Scan(&n, &s)
	rand.Seed(int64(s)) // set the seed for random number generation

	// generate the initial state of the universe using randomness
	universe := make([][]bool, n)
	for i := range universe {
		universe[i] = make([]bool, n)
		for j := range universe[i] {
			if rand.Intn(2) == 1 {
				universe[i][j] = true // cell is alive
			} else {
				universe[i][j] = false // cell is dead
			}
		}
	}

	// print the initial state of the universe
	for i := range universe {
		for j := range universe[i] {
			if universe[i][j] {
				fmt.Print("O")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
