package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Initialize the grid with all lights turned off (0).
func initializeGrid(size int) [][]int {
	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
	}
	return grid
}

// Apply a command to the grid based on the instruction.
func applyInstruction(grid [][]int, command string, x1, y1, x2, y2 int) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			switch command {
			case "turn on":
				grid[i][j] = 1
			case "turn off":
				grid[i][j] = 0
			case "toggle":
				if grid[i][j] == 1 {
					grid[i][j] = 0
				} else {
					grid[i][j] = 1
				}
			}
		}
	}
}

// Parse an instruction line to extract the command and coordinates.
func parseInstruction(instruction string) (string, int, int, int, int, error) {
	parts := strings.Fields(instruction)
	if len(parts) < 5 {
		return "", 0, 0, 0, 0, fmt.Errorf("invalid instruction format: %s", instruction)
	}

	command := parts[0] + " " + parts[1] // 'turn on', 'turn off', or 'toggle'

	// Extract starting coordinates.
	x1y1 := strings.Split(parts[2], ",")
	if len(x1y1) != 2 {
		return "", 0, 0, 0, 0, fmt.Errorf("invalid starting coordinates format: %s", parts[2])
	}
	x1, err := strconv.Atoi(x1y1[0])
	if err != nil {
		return "", 0, 0, 0, 0, fmt.Errorf("invalid x1 coordinate: %s", x1y1[0])
	}
	y1, err := strconv.Atoi(x1y1[1])
	if err != nil {
		return "", 0, 0, 0, 0, fmt.Errorf("invalid y1 coordinate: %s", x1y1[1])
	}

	// Extract ending coordinates.
	x2y2 := strings.Split(parts[4], ",")
	if len(x2y2) != 2 {
		return "", 0, 0, 0, 0, fmt.Errorf("invalid ending coordinates format: %s", parts[4])
	}
	x2, err := strconv.Atoi(x2y2[0])
	if err != nil {
		return "", 0, 0, 0, 0, fmt.Errorf("invalid x2 coordinate: %s", x2y2[0])
	}
	y2, err := strconv.Atoi(x2y2[1])
	if err != nil {
		return "", 0, 0, 0, 0, fmt.Errorf("invalid y2 coordinate: %s", x2y2[1])
	}

	return command, x1, y1, x2, y2, nil
}

// Count the number of lights that are turned on (value 1) in the grid.
func countLitLights(grid [][]int) int {
	count := 0
	for _, row := range grid {
		for _, light := range row {
			if light > 0 {
				count++
			}
		}
	}
	return count
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <data_file>")
		return
	}

	dataFile := os.Args[1]

	// Initialize the grid
	grid := initializeGrid(1000)

	// Open the file
	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := scanner.Text()
		command, x1, y1, x2, y2, parseErr := parseInstruction(instruction)
		if parseErr != nil {
			fmt.Println("Error parsing instruction:", parseErr)
			continue
		}
		applyInstruction(grid, command, x1, y1, x2, y2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Count the lit lights
	litLights := countLitLights(grid)

	// Print the result
	fmt.Println("Number of lights lit:", litLights)
}
