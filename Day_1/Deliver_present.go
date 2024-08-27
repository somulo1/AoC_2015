package main

import (
	"fmt"
	"os"
)

// count the number of "(" and the number of ")"
func Locate_Floor(puzzles string) int {
	countfloor := 0

	// fmt.Println(puzzles)
	for _, puzzle := range puzzles {
		if puzzle == '(' {
			countfloor++
		} else if puzzle == ')' {
			countfloor--
		}
	}
	return countfloor
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("USAGE: go run . something.txt")
		os.Exit(1)
	}
	str := os.Args[1]
	puzzlebit, err := os.ReadFile(str)
	if err != nil {
		fmt.Printf("error reading the file")
		os.Exit(1)
	}
	puzzles := string(puzzlebit)
	Destination := Locate_Floor(puzzles)
	fmt.Println("Final Destination:", Destination,"floor")
}
