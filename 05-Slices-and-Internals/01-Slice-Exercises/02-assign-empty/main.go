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
// EXERCISE: Assign empty slices
//
//   Assign empty slices to all the slices that you've declared in the previous
//   exercise, and print them here.
//
//
// EXPECTED OUTPUT
//  names    : []string 0 false
//  distances: []int 0 false
//  data     : []uint8 0 false
//  ratios   : []float64 0 false
//  alives   : []bool 0 false
// ---------------------------------------------------------

func main() {
	var (
		names     []string
		distances []int
		datas     []uint8
		ratios    []float64
		alives    []bool
	)

	names = []string{}
	distances = []int{}
	datas = []uint8{}
	ratios = []float64{}
	alives = []bool{}

	fmt.Printf("Name       :%-10s %-10s %-10s\n", "Type", "Length", "Nil")
	fmt.Printf("names      :%-10T %-10d %-10t\n", names, len(names), names == nil)
	fmt.Printf("distances  :%-10T %-10d %-10t\n", distances, len(distances), distances == nil)
	fmt.Printf("datas      :%-10T %-10d %-10t\n", datas, len(datas), datas == nil)
	fmt.Printf("ratios     :%-10T %-10d %-10t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives     :%-10T %-10d %-10t\n", alives, len(alives), alives == nil)
}
