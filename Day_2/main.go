package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <something.txt>")
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
	var totalRibbon int

	for scanner.Scan() {
		line := scanner.Text()
		dimensions := strings.Split(line, "x")

		if len(dimensions) != 3 {
			fmt.Printf("Invalid dimensions format: %s\n", line)
			continue
		}

		l, w, h, err := parseDimensions(dimensions)
		if err != nil {
			fmt.Printf("Error parsing dimensions: %v\n", err)
			continue
		}

		// Calculate the smallest perimeter directly in the assignment
		smallestPerimeter := 2 * min(l+w, min(w+h, h+l))

		// Calculate volume
		volume := l * w * h

		// Add the total ribbon needed for this present
		totalRibbon += smallestPerimeter + volume
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("Total feet of ribbon needed: %d\n", totalRibbon)
}

// Helper function to parse dimensions and handle errors
func parseDimensions(dimensions []string) (int, int, int, error) {
	l, err := strconv.Atoi(dimensions[0])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error converting length: %w", err)
	}
	w, err := strconv.Atoi(dimensions[1])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error converting width: %w", err)
	}
	h, err := strconv.Atoi(dimensions[2])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error converting height: %w", err)
	}
	return l, w, h, nil
}

// Helper function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
