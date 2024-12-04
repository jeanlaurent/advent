package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isMonotonic(numbers []int) bool {
	increasing := true
	decreasing := true

	for i := 1; i < len(numbers); i++ {
		if numbers[i] <= numbers[i-1] {
			increasing = false
		}
		if numbers[i] >= numbers[i-1] {
			decreasing = false
		}
	}

	return increasing || decreasing
}

func hasValidDifferences(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		diff := abs(numbers[i] - numbers[i-1])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isSafeReport(numbers []int) bool {
	return isMonotonic(numbers) && hasValidDifferences(numbers)
}

func main() {
	step1("example1.txt")
	step1("input.txt")
	step2("example1.txt")
	step2("input.txt")
}

func step1(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		strNumbers := strings.Fields(line)
		numbers := make([]int, len(strNumbers))

		for i, str := range strNumbers {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("Error converting string to number: %v\n", err)
				continue
			}
			numbers[i] = num
		}

		if isSafeReport(numbers) {
			safeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("Number of safe reports: %d\n", safeCount)
}

func step2(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		strNumbers := strings.Fields(line)
		numbers := make([]int, len(strNumbers))

		for i, str := range strNumbers {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("Error converting string to number: %v\n", err)
				continue
			}
			numbers[i] = num
		}

		// Check if it's already safe
		if isSafeReport(numbers) {
			safeCount++
			continue
		}

		// Try removing each number to see if it becomes safe
		for i := 0; i < len(numbers); i++ {
			// Create a new slice without the i-th number
			tempNumbers := make([]int, 0, len(numbers)-1)
			tempNumbers = append(tempNumbers, numbers[:i]...)
			tempNumbers = append(tempNumbers, numbers[i+1:]...)

			if isSafeReport(tempNumbers) {
				safeCount++
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("Number of safe reports in step 2: %d\n", safeCount)
}
