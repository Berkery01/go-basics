// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: List
//
//  Now, it's time to add an interface to your program using
//  the bufio.Scanner. So the users can list the games, or
//  search for the games by id.
//
//  1. Scan for the input in a loop (use bufio.Scanner)
//
//  2. Print the available commands.
//
//  3. Implement the quit command: Quits from the loop.
//
//  4. Implement the list command: Lists all the games.
//
//
// EXPECTED OUTPUT
//  Please run the solution and try the program with list and
//  quit commands.
// ---------------------------------------------------------

func main() {
	// use your solution from the previous exercise
	type item struct {
		id    int
		name  string
		price int
	}
	type game struct {
		item
		genre string
	}

	games := []game{
		{item: item{id: 1, name: "God of War", price: 50}, genre: "action adventure"},
		{item: item{id: 2, name: "x-com 2", price: 30}, genre: "strategy"},
		{item: item{id: 3, name: "Minecraft", price: 20}, genre: "sandbox"},
	}

	//fmt.Printf("%-5s %-15s %-10s %-20s\n", "ID", "Name", "Price", "Genre")
	//for _, v := range games {
	//	fmt.Printf("%-5d %-15s %-10d %-20s\n", v.id, v.name, v.price, v.genre)
	//}
	gamesbyID := map[int]game{}
	for _, v := range games {
		gamesbyID[v.id] = v
	}

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(`
  > list   : list all items
  > id [ID]   : list item with [ID]
  > quit   : quits
`)

		if !in.Scan() {
			break
		}

		fmt.Println()

		cmd := strings.Fields(in.Text())
		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n", g.id, g.name, "("+g.genre+")", g.price)
			}

		case "id":
			if len(cmd) != 2 {
				fmt.Println("Usage: id [ID]")
				continue
			}
			queryID, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Printf("Error: %q\n", err)
				return
			}
			query, ok := gamesbyID[queryID]
			if !ok {
				fmt.Printf("%q is not a valid ID.\n", queryID)
				continue
			}
			fmt.Printf("#%d: %-15q %-20s $%d\n", query.id, query.name, "("+query.genre+")", query.price)
		}
	}
}
