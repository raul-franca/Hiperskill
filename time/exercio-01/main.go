package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	nextYear := time.Now().Year() + 1

	firstJanuary := time.Date(nextYear, 1, 1, 0, 0, 0, 0, time.Local)

	fmt.Println(firstJanuary)

	// The hidden 'check' function checks your answer:
	check(firstJanuary.Sub(now))

}

func check(sub time.Duration) {

}
