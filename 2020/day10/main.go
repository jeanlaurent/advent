package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dataAsString := load("simple_input.txt")
	fmt.Println("step1 (simple) -->", step1(dataAsString))
	fmt.Println("step2 (simple) -->", step2(dataAsString))

	dataAsString = load("sample_input.txt")
	fmt.Println("step1 (sample) -->", step1(dataAsString))
	fmt.Println("step2 (sample) -->", step2(dataAsString))

	dataAsString = load("puzzle_input.txt")
	fmt.Println("step1 -->", step1(dataAsString))
	fmt.Println("step2 -->", step2(dataAsString))
}

func step2(dataAsString string) int {
	list := makeList(dataAsString)

	pathCount := map[int]int{0: 1}

	for origin := 0; origin < len(list); origin++ {
		nextStep := origin + 1

		for nextStep < len(list) && list[nextStep] <= list[origin]+3 {

			_, exists := pathCount[nextStep]
			if !exists {
				pathCount[nextStep] = 0
			}
			pathCount[nextStep] = pathCount[nextStep] + pathCount[origin]
			nextStep++
		}
	}
	return pathCount[len(list)-1]
}

func step1(dataAsString string) int {
	list := makeList(dataAsString)

	diff := map[int]int{0: 0, 1: 0, 2: 0, 3: 0}

	for i := 0; i < len(list)-1; i++ {
		diff[list[i+1]-list[i]]++
	}

	return diff[1] * diff[3]
}

func makeList(dataAsString string) []int {
	list := readInput(dataAsString)
	list = append(list, 0)
	sort.Ints(list)
	return append(list, list[len(list)-1]+3)
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
