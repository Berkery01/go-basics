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
// EXERCISE: Declare nil slices
//
//  1. Declare the following slices as nil slices:
//
//    1. The names of your friends (names slice)
//
//    2. The distances to locations (distances slice)
//
//    3. A data buffer (data slice)
//
//    4. Currency exchange ratios (ratios slice)
//
//    5. Up/Down status of web servers (alives slice)
//
//
//  2. Print their type, length and whether they're equal to nil value or not.
//
//
// EXPECTED OUTPUT
//  names    : []string 0 true
//  distances: []int 0 true
//  data     : []uint8 0 true
//  ratios   : []float64 0 true
//  alives   : []bool 0 true
// ---------------------------------------------------------

func main() {
	var (
		names     []string
		distances []int
		datas     []uint8
		ratios    []float64
		alives    []bool
	)
	fmt.Printf("Name       :%-10s %-10s %-10s\n", "Type", "Length", "Nil")
	fmt.Printf("names      :%-10T %-10d %-10t\n", names, len(names), names == nil)
	fmt.Printf("distances  :%-10T %-10d %-10t\n", distances, len(distances), distances == nil)
	fmt.Printf("datas      :%-10T %-10d %-10t\n", datas, len(datas), datas == nil)
	fmt.Printf("ratios     :%-10T %-10d %-10t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives     :%-10T %-10d %-10t\n", alives, len(alives), alives == nil)
}
