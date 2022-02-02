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
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Decode
//
//  At the beginning of the file:
//
//  1. Load the initial data to the game store from json.
//     (see the data constant below)
//
//  2. Load the decoded values into the usual `game` values (to the games slice as well).
//
//     So the rest of the program can work intact.
//
//
// HINT
//
//  Move the jsonGame type to the top and reuse it both when
//  loading the initial data, and in the "save" command.
//
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

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

	type jsonGame struct {
		ID    int    `json:"ID"`
		Name  string `json:"name"`
		Genre string `json:"genre"`
		Price int    `json:"price"`
	}

	var decoded []jsonGame
	if err := json.Unmarshal([]byte(data), &decoded); err != nil {
		fmt.Printf("Error: %q", err)
		return
	}

	var games []game
	for _, v := range decoded {
		games = append(games, game{item{v.ID, v.Name, v.Price}, v.Genre})
	}

	gamesbyID := map[int]game{}
	for _, v := range games {
		gamesbyID[v.id] = v
	}

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(`
  > list     : list all items
  > id [ID]  : list item with [ID]
  > save     : save as json and quit
  > quit     : quits
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

		case "save":
			var encodable []jsonGame
			for _, g := range games {
				encodable = append(encodable, jsonGame{g.id, g.name, g.genre, g.price})
			}
			output, err := json.MarshalIndent(encodable, "", "\t")
			if err != nil {
				fmt.Printf("Error: %q\n", err)
				continue
			}
			fmt.Println(string(output))
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
