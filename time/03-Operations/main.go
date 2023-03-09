package main

import (
	"fmt"
	"time"
)

func main() {
	twoSecond := time.Millisecond*1000 + time.Second
	sevenDays := time.Hour * 24 * 7
	oneYear := time.Hour * 24 * 365
	threeYearsAndTwoMonths := oneYear*3 + time.Hour*24*30

	fmt.Println(twoSecond)
	fmt.Println(sevenDays)
	fmt.Println(oneYear)
	fmt.Println(threeYearsAndTwoMonths)

	// Output:
	// 2s
	// 168h0m0s
	// 8760h0m0s
	// 27000h0m0s

	CopernicusBDay := time.Date(1473, time.February, 19, 0, 0, 0, 0, time.UTC)
	NewtonBDay := CopernicusBDay.Add(time.Hour * 1489080)
	EinsteinBDay := time.Date(1879, time.March, 14, 0, 0, 0, 0, time.UTC)

	fmt.Println("Timeline:")
	fmt.Printf("Nicolaus Copernicus: %v\n", CopernicusBDay)
	fmt.Printf(" | %v\n", NewtonBDay.Sub(CopernicusBDay))
	fmt.Printf("Isaac Newton: %v\n", NewtonBDay)
	fmt.Printf(" | %v\n", EinsteinBDay.Sub(NewtonBDay))
	fmt.Printf("Albert Einstein: %v\n", EinsteinBDay)
}

// Output:
// Timeline:
// Nicolaus Copernicus: 1473-02-19 00:00:00 +0000 UTC
//  | 1489080h0m0s
// Isaac Newton: 1643-01-04 00:00:00 +0000 UTC
//  | 2070384h0m0s
// Albert Einstein: 1879-03-14 00:00:00 +0000 UTC
