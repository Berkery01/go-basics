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
	"os"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Students
//
//  Create a program that returns the students by the given
//  Hogwarts house name (see the data below).
//
//  Print the students sorted by name.
//
//  "bobo" doesn't belong to Hogwarts, remove it by using
//  the delete function.
//
//
// RESTRICTIONS
//
//  + Add the following data to your map as is.
//    Do not sort it manually and do not modify it.
//
//  + Slices in the map shouldn't be sorted (changed).
//    HINT: Copy them.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//  Please type a Hogwarts house name.
//
//
//  go run main.go bobo
//
//  Sorry. I don't know anything about "bobo".
//
//
//  go run main.go hufflepuf
//
//  ~~~ hufflepuf students ~~~
//
//        + diggory
//        + helga
//        + scamander
//        + wenlock
//
// ---------------------------------------------------------

func main() {
	// House        Student Name
	// ---------------------------
	// gryffindor   weasley
	// gryffindor   hagrid
	// gryffindor   dumbledore
	// gryffindor   lupin
	// hufflepuf    wenlock
	// hufflepuf    scamander
	// hufflepuf    helga
	// hufflepuf    diggory
	// ravenclaw    flitwick
	// ravenclaw    bagnold
	// ravenclaw    wildsmith
	// ravenclaw    montmorency
	// slytherin    horace
	// slytherin    nigellus
	// slytherin    higgs
	// slytherin    scorpius
	// bobo         wizardry
	// bobo         unwanted

	houses := map[string][]string{
		"gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuf":  {"wenlock", "scamander", "helga", "diggory"},
		"ravenclaw":  {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin":  {"horace", "nigellus", "higgs", "scorpius"},
		"bobo":       {"wizardry", "unwanted"},
	}

	delete(houses, "bobo")

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Try writing a Hogwards house name.")
		return
	}

	if names, ok := houses[args[0]]; ok {
		fmt.Printf("~~~ %s students ~~~\n", args[0])

		clone := append([]string{}, names...)
		sort.Strings(clone)
		for _, v := range clone {
			fmt.Printf("\t+ %s\n", v)
		}
		return
	}
	fmt.Printf("There is no house named %s\n", args[0])

}
