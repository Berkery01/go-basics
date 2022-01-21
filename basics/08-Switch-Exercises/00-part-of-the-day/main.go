package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Now().Hour()
	fmt.Println(time)

	switch {
	case time >= 18:
		fmt.Println("Good evening")
	case time >= 12:
		fmt.Println("Good afternoon")
	case time >= 6:
		fmt.Println("Good morning")
	default:
		fmt.Println("Good night")
	}
}
