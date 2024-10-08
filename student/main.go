package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"student/helperfunctions"
)

func main() {
	// Define the window size
	
	var data []float64

	// Create a scanner to read from standard input
	input := os.Stdin
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		lines := strings.ReplaceAll(line, " ", "")
		num, err := strconv.ParseFloat(lines, 64)
		if err != nil {
			fmt.Printf("Invalid input: %v\n", err)
		}
		if num < math.MinInt64 || num > math.MaxInt64 {
			fmt.Printf("cannot use an overflow value %.f\n", num)
			return
		}

		// Add the new number to the data slice
		data = append(data, num)

		// Maintain a fixed window size
	

		// Calculate mean and standard deviation for the current window
		stdDev := helperfunctions.StandardDeviation(data)
		average := helperfunctions.CalculateMean(data)

		// Calculate and print the lower and upper bounds
		if len(data) > 1 {
			lower := math.Round(average - 1.96*stdDev)
			upper := math.Round(average + 1.96*stdDev)

			if lower < math.MinInt64 || upper > math.MaxInt64 {
				fmt.Printf("cannot use an overflow value %.f\n", num)
				return
			}
			fmt.Println(lower, upper)
		}
	}

	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading standard input: %v", err)
	}
}
