package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 50
	height = 20
)

func main() {
	// Cria uma grade aleatória inicial.
	rand.Seed(time.Now().UnixNano())
	var grid [width][height]bool
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			grid[x][y] = rand.Intn(2) == 0
		}
	}

	// Reproduzir a grade.
	for {
		// Limpa a tela antes de exibir a nova grade.
		fmt.Print("\033[H\033[2J")

		// Imprime a grade atual.
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if grid[x][y] {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}

		// Cria uma nova grade com base na grade atual.
		var newGrid [width][height]bool
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				neighbors := 0
				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						if dx == 0 && dy == 0 {
							continue
						}
						nx, ny := x+dx, y+dy
						if nx < 0 {
							nx += width
						} else if nx >= width {
							nx -= width
						}
						if ny < 0 {
							ny += height
						} else if ny >= height {
							ny -= height
						}
						if grid[nx][ny] {
							neighbors++
						}
					}
				}
				if grid[x][y] {
					// A célula atual está viva.
					if neighbors == 2 || neighbors == 3 {
						newGrid[x][y] = true
					}
				} else {
					// A célula atual está morta.
					if neighbors == 3 {
						newGrid[x][y] = true
					}
				}
			}
		}

		// Usa a nova grade como a grade atual para a próxima iteração.
		grid = newGrid

		// Espera um pouco antes de exibir a próxima iteração.
		time.Sleep(100 * time.Millisecond)
	}
}
