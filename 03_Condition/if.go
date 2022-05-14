package main

import "fmt"

func main() {
	a := 11.0
	b := 20.0
	if frac := a / b; frac > 0.5 {
		fmt.Printf("a is Big than b")
	}
}
