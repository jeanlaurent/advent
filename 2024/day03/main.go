package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func step1(filename string) int {
	// Read the file content
	content, err := os.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}

	// Create regex to match mul(X,Y) where X and Y are 1-3 digits
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// Find all matches in the content
	matches := re.FindAllStringSubmatch(string(content), -1)

	sum := 0
	for _, match := range matches {
		// Convert string numbers to integers
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])

		// Multiply and add to sum
		sum += x * y
	}

	return sum
}

func step2(filename string) int {
	// Read the file content
	content, err := os.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}

	// Create regex patterns for all instructions
	mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	// doRe := regexp.MustCompile(`do\(\)`)
	// dontRe := regexp.MustCompile(`don't\(\)`)

	// Find all instructions in order
	allInstructions := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\))`)

	enabled := true // Instructions are enabled by default
	sum := 0

	matches := allInstructions.FindAllString(string(content), -1)
	for _, instruction := range matches {
		// Check for do() or don't() instructions
		if instruction == "do()" {
			enabled = true
			continue
		}
		if instruction == "don't()" {
			enabled = false
			continue
		}

		// Process mul instructions only if enabled
		if enabled {
			if mulMatch := mulRe.FindStringSubmatch(instruction); mulMatch != nil {
				x, _ := strconv.Atoi(mulMatch[1])
				y, _ := strconv.Atoi(mulMatch[2])
				sum += x * y
			}
		}
	}

	return sum
}

func main() {
	fmt.Printf("Sum of all multiplications: %d\n", step1("example1.txt"))
	fmt.Printf("Sum of all multiplications: %d\n", step1("input.txt"))
	fmt.Printf("Sum of all multiplications (step 2): %d\n", step2("example1.txt"))
	fmt.Printf("Sum of all multiplications (step 2): %d\n", step2("input.txt"))
}
