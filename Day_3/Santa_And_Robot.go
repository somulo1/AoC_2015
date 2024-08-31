package main

import (
	"bufio"
	"fmt"
	"os"
)

// Helper function to process directions for both Santa and Robo-Santa
func countUniqueHousesWithRobot(directions string) int {
	visited := make(map[[2]int]bool)
	santaX, santaY := 0, 0
	robotX, robotY := 0, 0

	// Mark the starting position as visited for both Santa and Robo-Santa
	visited[[2]int{santaX, santaY}] = true

	// Process each direction
	for i, dir := range directions {
		if i%2 == 0 {
			// Santa's move
			switch dir {
			case '^':
				santaY++
			case 'v':
				santaY--
			case '>':
				santaX++
			case '<':
				santaX--
			}
			visited[[2]int{santaX, santaY}] = true
		} else {
			// Robo-Santa's move
			switch dir {
			case '^':
				robotY++
			case 'v':
				robotY--
			case '>':
				robotX++
			case '<':
				robotX--
			}
			visited[[2]int{robotX, robotY}] = true
		}
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
		uniqueHouses := countUniqueHousesWithRobot(directions)
		fmt.Printf("Number of unique houses: %d\n", uniqueHouses)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
}
