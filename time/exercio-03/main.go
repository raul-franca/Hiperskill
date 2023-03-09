package main

import (
	"fmt"
	"time"
)

// The first travel check conditions have been written for you:
func checkFirstTravel(date1 time.Time, date2 time.Time, date3 time.Time) {
	switch {
	case date1.After(date2) && date1.After(date3):
		fmt.Println(date1.Date())
	case date2.After(date3):
		fmt.Println(date2.Date())
	default:
		fmt.Println(date3.Date())
	}
}

// Write the correct check conditions for the second travel below:
func checkSecondTravel(date1 time.Time, date2 time.Time, date3 time.Time) {
	switch {
	case date1.Before(date2) && date1.Before(date3):
		fmt.Println(date1.Date())
	case date2.Before(date3):
		fmt.Println(date2.Date())
	default:
		fmt.Println(date3.Date())
	}
}

// Write the correct check conditions for the third travel below:
func checkThirdTravel(date1 time.Time, date2 time.Time, date3 time.Time) {
	switch {
	case date1.After(date2) && date1.Before(date3) || date1.Before(date2) && date1.After(date3):
		fmt.Println(date1.Date())
	case date2.After(date1) && date2.Before(date3) || date2.Before(date1) && date2.After(date3):
		fmt.Println(date2.Date())
	default:
		fmt.Println(date3.Date())
	}
}

// DO NOT delete or modify the contents of the main() function!
func main() {
	date1, date2, date3 := readDate(), readDate(), readDate()

	// The first travel — checks conditions for the latest date:
	checkFirstTravel(date1, date2, date3)

	// The second travel — checks conditions for the earliest date:
	checkSecondTravel(date1, date2, date3)

	// The third travel — checks conditions for the middle date:
	checkThirdTravel(date1, date2, date3)
}

// DO NOT modify the readDate function!
func readDate() time.Time {
	var year, month, day int
	fmt.Scan(&year, &month, &day)
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
