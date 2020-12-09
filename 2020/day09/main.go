package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	step1Result := step1(load("sample_input.txt"), 5)
	fmt.Println("step1-sample -->", step1Result)
	fmt.Println("step2-sample -->", step2(load("sample_input.txt"), 5, step1Result))

	step1Result = step1(load("puzzle_input.txt"), 25)
	fmt.Println("step1 -->", step1Result)
	fmt.Println("step2 -->", step2(load("puzzle_input.txt"), 25, step1Result))
}

func step2(dataAsString string, preambleLength int, target int) int {
	list := readInput(dataAsString)
	for rangeSize := 2; rangeSize < len(list)/2; rangeSize++ {
		for i := 0; i < len(list)-rangeSize; i++ {
			smallest := math.MaxInt64
			biggest := -1
			sum := 0
			for j := i; j < i+rangeSize; j++ {
				if list[j] < smallest {
					smallest = list[j]
				}
				if list[j] > biggest {
					biggest = list[j]
				}
				sum += list[j]
			}
			if sum == target {
				result := smallest + biggest
				return result
			}
		}
	}
	return -1
}

func step1(dataAsString string, preambleLength int) int {
	list := readInput(dataAsString)
	for position := preambleLength + 1; position < len(list); position++ {
		sum := []int{}
		for j := position - preambleLength - 1; j < position-1; j++ {
			for i := j + 1; i < position; i++ {
				// fmt.Println(list[j], "+", list[i])
				currentSum := list[j] + list[i]
				if !exist(sum, currentSum) {
					sum = append(sum, currentSum)
				}
			}
		}
		if !exist(sum, list[position]) {
			return list[position]
		}
	}
	return -1
}

func readInput(dataAsString string) []int {
	list := []int{}
	for _, numberAsString := range strings.Split(dataAsString, "\n") {
		list = append(list, atoi(numberAsString))
	}
	return list
}

// ----

func exist(array []int, item int) bool {
	for _, element := range array {
		if element == item {
			return true
		}
	}
	return false
}

func load(filename string) string {
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	return string(text)
}

func atoi(line string) int {
	number, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return number
}
