package main

import (
	"fmt"
	"os"
)

func Find_Position(str string)int {
	countfloor := 0
	position := 0
	// fmt.Println(puzzles)
	for i, puzzle := range str{
		if puzzle == '(' {
			countfloor++
			i++
		} else if puzzle == ')' {
			countfloor--
			i++
		}
		position += i
		fmt.Println(i)
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
	Destination := Find_Position(puzzles)
	fmt.Println("Final Destination:", Destination, "floor")
}
