package main

import (
	"fmt"
	"time"
)

func main() {
	//                    YEAR   MONTH    D  H  M  S  nS Location
	starWars := time.Date(1977, time.May, 4, 0, 0, 0, 0, time.UTC)

	fmt.Println(starWars)
	fmt.Println(starWars.Date())
	fmt.Println(time.Now())
	fmt.Println(time.Now().Date())
}

// Output:
// 1977-05-04 00:00:00 +0000 UTC
// 1977 May 4
// 2022-04-05 15:51:03-Operations.410972 +0500 +05 m=+0.000249376
