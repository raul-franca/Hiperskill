package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n, s, m int
	fmt.Scan(&n, &s, &m)

	// Initialize current and next generations
	currentGen := make([][]bool, n)
	nextGen := make([][]bool, n)
	for i := range currentGen {
		currentGen[i] = make([]bool, n)
		nextGen[i] = make([]bool, n)
	}

	// Seed the random number generator
	rand.Seed(int64(s))

	// Fill the initial generation randomly
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			currentGen[i][j] = rand.Intn(2) == 1
		}
	}

	// Apply the rules for M generations
	for gen := 1; gen <= m; gen++ {
		// Compute the next generation
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// Count the live neighbors
				liveNeighbors := 0
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						if !(x == 0 && y == 0) { // exclude self
							// Get the coordinates of the neighbor
							neighborX := (i + x + n) % n // handle border cells
							neighborY := (j + y + n) % n

							if currentGen[neighborX][neighborY] {
								liveNeighbors++
							}
						}
					}
				}

				// Apply the rules
				if currentGen[i][j] {
					if liveNeighbors == 2 || liveNeighbors == 3 {
						nextGen[i][j] = true
					} else {
						nextGen[i][j] = false
					}
				} else {
					if liveNeighbors == 3 {
						nextGen[i][j] = true
					} else {
						nextGen[i][j] = false
					}
				}
			}
		}

		// Copy the next generation to the current generation
		for i := range currentGen {
			copy(currentGen[i], nextGen[i])
		}
	}

	// Print the final generation
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if currentGen[i][j] {
				fmt.Print("O")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
