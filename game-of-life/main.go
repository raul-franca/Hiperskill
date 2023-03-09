package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	fmt.Scan(&n)

	currentGeneration := make([][]bool, n)
	nextGeneration := make([][]bool, n)

	for i := range currentGeneration {
		currentGeneration[i] = make([]bool, n)
		nextGeneration[i] = make([]bool, n)
		for j := range currentGeneration[i] {
			currentGeneration[i][j] = rand.Intn(2) == 1
		}
	}

	for gen := 1; gen <= 11; gen++ {
		fmt.Printf("Generation #%d\n", gen)
		printUniverse(currentGeneration)

		time.Sleep(500 * time.Millisecond)
		//fmt.Print("\033[H\033[2J") // Clear console

		for i := range currentGeneration {
			for j := range currentGeneration[i] {
				neighbors := countLiveNeighbors(currentGeneration, i, j)

				if currentGeneration[i][j] && (neighbors == 2 || neighbors == 3) {
					nextGeneration[i][j] = true
				} else if !currentGeneration[i][j] && neighbors == 3 {
					nextGeneration[i][j] = true
				} else {
					nextGeneration[i][j] = false
				}
			}
		}

		currentGeneration, nextGeneration = nextGeneration, currentGeneration
	}
}

func countLiveNeighbors(currentGeneration [][]bool, i, j int) int {
	count := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			r := i + x
			c := j + y
			if r < 0 || r >= len(currentGeneration) || c < 0 || c >= len(currentGeneration[i]) {
				continue
			}
			if currentGeneration[r][c] {
				count++
			}
		}
	}
	return count
}

func printUniverse(currentGeneration [][]bool) {
	aliveCount := 0
	for i := range currentGeneration {
		for j := range currentGeneration[i] {
			if currentGeneration[i][j] {
				fmt.Print("O")
				aliveCount++
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Printf("Alive: %d\n", aliveCount)
}
