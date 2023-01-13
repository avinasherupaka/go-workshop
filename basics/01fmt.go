package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello young gopher")

	a := "double"
	b := "rainbow"
	// TODO print `a` and `b` to the console as part of a sentence
	fmt.Print(a, b)

	const c = 42
	const d = 6.022e23
	// TODO use math.Min to find the minimum of the two constants `c` and `d` and print the result to the console
	fmt.Print(math.Min(c, d))
	// TODO add a variable `e` to make the following console message return true
	const powerLevel = 9000
	const e = 10000
	fmt.Printf("It's over 9000? %v\n", e > powerLevel)
}
