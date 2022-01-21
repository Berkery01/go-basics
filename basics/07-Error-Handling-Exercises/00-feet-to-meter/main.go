package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: [number]")
		return
	}

	feet, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	meter := feet * 0.3048
	fmt.Printf("%g feet is %g meters.\n", feet, meter)

}
