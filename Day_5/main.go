package main

import (
	"bufio"
	"fmt"
	"os"
)

// hasNonOverlappingPair checks if the string has a pair of any two letters that appears at least twice without overlapping.
func hasNonOverlappingPair(s string) bool {
	seen := make(map[string]int)
	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		if index, exists := seen[pair]; exists && index < i-1 {
			return true
		}
		seen[pair] = i
	}
	return false
}

// hasRepeatingLetterWithOneBetween checks if the string contains at least one letter which repeats with exactly one letter between them.
func hasRepeatingLetterWithOneBetween(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

// isNice checks if a string meets the new criteria to be considered nice.
func isNice(s string) bool {
	return hasNonOverlappingPair(s) && hasRepeatingLetterWithOneBetween(s)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <inputfile>")
		return
	}

	inputFile := os.Args[1]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	niceCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if isNice(line) {
			niceCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Number of nice strings: %d\n", niceCount)
}
