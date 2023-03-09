package main

import "fmt"

func main() {

	elements := make(map[string]string, 3) // the initial capacity of 'elements' is 3
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"

	fmt.Println("element length is:", len(elements)) // element length is: 3
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"

	fmt.Println("element length is:", len(elements))
	delete(elements, "H")

	fmt.Println(elements)                            // map[B:Boron Be:Beryllium He:Helium Li:Lithium]
	fmt.Println("element length is:", len(elements)) // element length is: 4

	for key := range elements {
		fmt.Println(key)
	}

	movieRatings := map[string]int{
		"The Matrix":          88,
		"Speed":               94,
		"The Matrix Reloaded": 73,
		"John Wick":           86,
	}

	fmt.Println("List Option #1")
	// Option #1 - create the 'val' variable to print the values of the map
	for key, val := range movieRatings {
		fmt.Println(key, val)
	}
	fmt.Println("List Option #2")
	// Option #2 - pass the 'key' variable within [] square brackets after the map's name
	for key := range movieRatings {
		fmt.Println(key, movieRatings[key])
	}

	vegetables := map[string]bool{
		"🥕": true,
		"🧅": true,
		"🥦": true,
	}

	if _, ok := vegetables["🥕"]; ok {
		fmt.Println("🥕 is in the set.")
	}

	if _, ok := vegetables["🍇"]; ok {
		fmt.Println("🍇 is in the set.")
	}

}
