package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    // Open the file
    file, err := os.Open("data.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    totalWrappingPaper := 0

    // Process each line in the file
    for scanner.Scan() {
        line := scanner.Text()
        dimensions := strings.Split(line, "x")

        if len(dimensions) != 3 {
            fmt.Println("Invalid dimensions format:", line)
            continue
        }

        l, _ := strconv.Atoi(dimensions[0])
        w, _ := strconv.Atoi(dimensions[1])
        h, _ := strconv.Atoi(dimensions[2])

        // Calculate surface area
        surfaceArea := 2*l*w + 2*w*h + 2*h*l

        // Find the smallest side area for slack
        minSide := min(l*w, min(w*h, h*l))

        // Total wrapping paper needed for this present
        totalPaperForPresent := surfaceArea + minSide
        totalWrappingPaper += totalPaperForPresent
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    fmt.Printf("Total square feet of wrapping paper needed: %d\n", totalWrappingPaper)
}

// Helper function to find the minimum of three integers
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
