// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices
//
//  We have received housing prices. Your task is to load the data into
//  appropriately typed slices then print them.
//
//  1. Check out the expected output
//
//
//  2. Check out the code below
//
//     1. header   : stores the column headers
//     2. data     : stores the real data related to each column
//     3. separator: you will use it to separate the data
//
//
//  3. Parse the data
//
//     1. Separate it into rows by using the newline character ("\n")
//
//     2. For each row, separate it by using the separator (",")
//
//
//  4. Load the data into distinct slices
//
//     1. Load the locations into a []string
//     2. Load the sizes into []int
//     3. Load the beds into []int
//     4. Load the baths into []int
//     5. Load the prices into []int
//
//
//  5. Print the header
//
//     1. Separate it by using the separator
//
//     2. Print each column 15 character wide ("%-15s")
//
//
//  6. Print the rows from the slices that you've created, line by line
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//
// HINTS
//
//  + strings.Split function can separate a string into a []string
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	h := strings.Split(header, separator)
	d := strings.Split(data, "\n")

	var (
		loc                      []string
		size, beds, baths, price []int
		temp                     int
	)
	for _, v := range d {
		cols := strings.Split(v, separator)
		loc = append(loc, cols[0])

		temp, _ = strconv.Atoi(cols[1])
		size = append(size, temp)

		temp, _ = strconv.Atoi(cols[2])
		beds = append(beds, temp)

		temp, _ = strconv.Atoi(cols[3])
		baths = append(baths, temp)

		temp, _ = strconv.Atoi(cols[4])
		price = append(price, temp)
	}

	for i := range h {
		fmt.Printf("%-15s", h[i])
	}

	fmt.Println()

	for i := range d {
		fmt.Printf("%-15s", loc[i])
		fmt.Printf("%-15d", size[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d", price[i])
		fmt.Println()
	}

	var (
		totalSize, totalBeds, totalBaths, totalPrice float64
		avgSize, avgBeds, avgBaths, avgPrice         float64
	)

	for i := range d {
		totalSize += float64(size[i])
		totalBeds += float64(beds[i])
		totalBaths += float64(baths[i])
		totalPrice += float64(price[i])
	}
	n := float64(len(d))
	avgSize = totalSize / n
	avgBeds = totalBeds / n
	avgBaths = totalBaths / n
	avgPrice = totalPrice / n
	fmt.Printf("%-15.2s", "")
	fmt.Printf("%-15.2f", avgSize)
	fmt.Printf("%-15.2f", avgBeds)
	fmt.Printf("%-15.2f", avgBaths)
	fmt.Printf("%-15.2f", avgPrice)
	fmt.Println()
}
