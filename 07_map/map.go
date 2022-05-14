package main

import "fmt"

func main() {
	bar := map[string]float64{
		"Amazone":   100.25,
		"Google":    11.23,
		"Microsoft": 28.23,
		"Apple":     456.20,
	}
	fmt.Println(bar["Microsoft"])

	//null
	fmt.Println(bar["Apple"])

	//condition
	Value, ok := bar["Apple"]

	if !ok {
		fmt.Println("No Mojud")
	} else {
		fmt.Println(Value)
	}
}
