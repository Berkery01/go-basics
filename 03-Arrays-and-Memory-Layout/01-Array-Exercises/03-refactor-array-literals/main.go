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
// EXERCISE: Refactor to Array Literals
//
//  1. Use the 02-get-set-arrays exercise
//
//  2. Refactor the array assignments to array literals
//
//    1. You would need to change the array declarations to array literals
//
//    2. Then, you would need to move the right-hand side of the assignments,
//       into the array literals.
//
// EXPECTED OUTPUT
//   The output should be the same as the 02-get-set-arrays exercise.
// ---------------------------------------------------------

func main() {

	var (
		names     = [3]string{"A", "B", "C"}
		distances = [5]int{100, 200, 300, 400, 500}
		data      = [5]uint8{12, 24, 36, 48, 60}
		ratios    = [1]float64{1.42}
		alives    = [4]bool{true, false, false, true}
		zero      = [0]uint8{}
	)
	fmt.Printf("names     : %q\n", names)
	fmt.Printf("distances : %d\n", distances)
	fmt.Printf("data      : %d\n", data)
	fmt.Printf("ratios    : %.2f\n", ratios)
	fmt.Printf("alives    : %t\n", alives)
	fmt.Printf("zero      : %d\n", zero)
}
