package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	nSizeOfUniverser := 0
	seed := 0
	fmt.Scan(&nSizeOfUniverser, &seed)
	if nSizeOfUniverser < 0 {
		log.Fatalf("the size of universe most be great then 0")
	}
	for y := 0; y < nSizeOfUniverser; y++ {
		for x := 0; x < nSizeOfUniverser; x++ {
			if rand.Intn(2) == 1 {
				fmt.Print("O")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
