package main

import (
	"bufio"
	"fmt"
	"os"
)

// Helper function to process directions and return the number of unique houses
func countUniqueHouses(directions string) int {
	visited := make(map[[2]int]bool)
	x, y := 0, 0

	// Mark the starting position as visited
	visited[[2]int{x, y}] = true

	// Process each direction
	for _, dir := range directions {
		switch dir {
		case '^':
			y++
		case 'v':
			y--
		case '>':
			x++
		case '<':
			x--
		}
		visited[[2]int{x, y}] = true
	}

	return len(visited)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <Data.txt>")
		return
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		directions := scanner.Text()
		uniqueHouses := countUniqueHouses(directions)
		fmt.Printf("Number of unique houses: %d\n", uniqueHouses)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
}
