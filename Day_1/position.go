package main

import (
	"fmt"
	"os"
)

// Find_Position calculates the position where Santa first reaches the basement (floor -1)
func Find_Position(str string) int {
	countup := 0
	countdown := 0
	position := 0
	// Iterate over each character in the puzzles
	for i, puzzle := range str {
		if puzzle == '(' {
			countup++
		} else if puzzle == ')' {
			countdown++
		}

		// Check if Santa has entered the basement
		if countdown > countup {
			position = i + 1 // Return position
			break
		}
	}

	//  if the basement is never reached
	return position
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

	fmt.Printf("The first time santa gets basement is at position %d\n", position)
}
