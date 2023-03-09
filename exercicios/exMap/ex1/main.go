package main

import "fmt"

func main() {
	var ranks map[string]string     // This only declares the 'ranks' map variable
	ranks = make(map[string]string) // Option #1 to initialize the previously declared 'ranks' map

	// var ranks = make(map[string]string) // Option #2 to declare & initialize the 'ranks' map
	// ranks := make(map[string]string) // Option #3 to assign, declare & initialize the 'ranks' map
	ranks["gold"] = "🥇"
	ranks["silver"] = "🥈"
	ranks["bronze"] = "🥉"

	fmt.Println(ranks) // map[bronze:🥉 gold:🥇 silver:🥈]
	fmt.Printf("The value of key: 'gold' = %v  \n", ranks["gold"])

	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
	}

	fmt.Println(elements) // map[H:Hydrogen He:Helium Li:Lithium]

}
