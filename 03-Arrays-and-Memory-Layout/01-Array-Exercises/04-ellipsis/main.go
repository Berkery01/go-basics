// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Refactor to Ellipsis
//
//  1. Use the 03-array-literal exercise
//
//  2. Refactor the length of the array literals to ellipsis
//
//    This means: Use the ellipsis instead of defining the array's length
//                manually.
//
// EXPECTED OUTPUT
//   The output should be the same as the 03-array-literal exercise.
// ---------------------------------------------------------

func main() {

	var (
		names     = [...]string{"A", "B", "C"}
		distances = [...]int{100, 200, 300, 400, 500}
		data      = [...]uint8{12, 24, 36, 48, 60}
		ratios    = [...]float64{1.42}
		alives    = [...]bool{true, false, false, true}
		zero      = [...]uint8{}
	)
	fmt.Printf("names     : %q\n", names)
	fmt.Printf("distances : %d\n", distances)
	fmt.Printf("data      : %d\n", data)
	fmt.Printf("ratios    : %.2f\n", ratios)
	fmt.Printf("alives    : %t\n", alives)
	fmt.Printf("zero      : %d\n", zero)
}
