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
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//
// ---------------------------------------------------------

func main() {
	// Hint: Store phone numbers as text

	// #1
	// Key        : Last name
	// Element    : Phone number

	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable

	// #3
	// Key        : Last name
	// Element    : Phone numbers

	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity

	phone := map[string]string{
		"lastname1": "5002001010",
		"lastname2": "5003002020",
		"lastname3": "5004004040",
	}

	products := map[int]bool{
		223322: true,
		112233: false,
		001122: true,
	}

	numbers := map[string][]string{
		"lastname1": {"5002001010", "5003001010"},
		"lastname2": {"5003002020", "5003331010"},
		"lastname3": {"5004004040", "5223445555"},
	}

	basket := map[int]map[int]int{
		1: {123123: 10, 1333: 4},
	}

	user, p := "lastname1", "N/A"
	if v, ok := phone[user]; ok {
		p = v
	}
	fmt.Printf("%s 's phone number is %q\n", user, p)

	id, status := 223322, "available"
	if !products[id] {
		status = "not" + status
	}
	fmt.Printf("Product: %d is %s\n", id, status)

	user2, p := "lastname2", "N/A"
	if v, ok := numbers[user2]; ok {
		p = v[1]
	}
	fmt.Printf("%s 's second phone number is %q\n", user2, p)

	cid, pid := 1, 123123
	count := basket[cid][pid]
	fmt.Printf("User %d has product %d in the basket with amount of %d\n", cid, pid, count)
}
