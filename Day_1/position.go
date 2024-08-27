package main

import (
	"fmt"
	"os"
)

// Find_Position calculates the position where Santa first reaches the basement (floor -1)
func Find_Position(str string) int {
	countfloor := 0
	// Iterate over each character in the puzzles
	for i, puzzle := range str {
		if puzzle == '(' {
			countfloor++
		} else if puzzle == ')' {
			countfloor--
		}
		
		// Check if Santa has entered the basement
		if countfloor == -1 {
			return i + 1 // Return position
		}
	}
	
	//  if the basement is never reached
	return -1
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("USAGE: go run . filename")
		os.Exit(1)
	}
	str := os.Args[1]
	puzzlebit, err := os.ReadFile(str)
	if err != nil {
		fmt.Printf("error reading the file: %v\n", err)
		os.Exit(1)
	}
	puzzles := string(puzzlebit)
	position := Find_Position(puzzles)
	if position != -1 {
		fmt.Printf("The first time santa gets basement is at position %d\n", position)
	} else {
		fmt.Println("Santa never enters the basement.")
	}
}
