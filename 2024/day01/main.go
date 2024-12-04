package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// step1("example.txt")
	// step1("input.txt")
	step2("example.txt")
	step2("input.txt")
}

func step1(filename string) {
	// Create slices to store numbers from each column
	var column1 []int
	var column2 []int

	// Read the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Process each line
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line by whitespace to get the two numbers
		fields := strings.Fields(line)

		if len(fields) == 2 {
			// Convert strings to integers
			num1, err1 := strconv.Atoi(fields[0])
			num2, err2 := strconv.Atoi(fields[1])

			if err1 != nil || err2 != nil {
				fmt.Printf("Error converting numbers on line: %s\n", line)
				continue
			}

			// Add numbers to their respective columns
			column1 = append(column1, num1)
			column2 = append(column2, num2)
		}
	}

	// Sort both columns
	sort.Ints(column1)
	sort.Ints(column2)

	// Print sorted columns
	fmt.Println("Sorted Column 1:", column1)
	fmt.Println("Sorted Column 2:", column2)

	// Calculate differences and sum
	totalDiff := 0
	for i := 0; i < len(column1); i++ {
		diff := int(math.Abs(float64(column2[i] - column1[i])))
		fmt.Printf("Difference at position %d: %d - %d = %d\n", i, column2[i], column1[i], diff)
		totalDiff += diff
	}
	fmt.Printf("Total sum of differences: %d\n", totalDiff)

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}

func step2(filename string) {
	// Create slices to store numbers from each column
	var column1 []int
	var column2 []int

	// Read the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Process each line
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)

		if len(fields) == 2 {
			num1, err1 := strconv.Atoi(fields[0])
			num2, err2 := strconv.Atoi(fields[1])

			if err1 != nil || err2 != nil {
				fmt.Printf("Error converting numbers on line: %s\n", line)
				continue
			}

			column1 = append(column1, num1)
			column2 = append(column2, num2)
		}
	}

	// Create a frequency map for the right list
	rightFreq := make(map[int]int)
	for _, num := range column2 {
		rightFreq[num]++
	}

	// Calculate similarity score
	totalScore := 0
	for _, leftNum := range column1 {
		frequency := rightFreq[leftNum]
		score := leftNum * frequency
		fmt.Printf("Step 2 - Number %d appears %d times in right list, score: %d\n",
			leftNum, frequency, score)
		totalScore += score
	}

	fmt.Printf("Step 2 - Total similarity score: %d\n", totalScore)

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
