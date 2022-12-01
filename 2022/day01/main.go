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
	step1("input.txt")
	step2("input.txt")
}

func step1(filename string) {
	dataAsString := load(filename)

	maxCalories := -1
	calories := 0

	for _, line := range strings.Split(dataAsString, "\n") {
		if line == "" {
			if calories > maxCalories {
				maxCalories = calories
			}
			calories = 0
			continue
		}
		calories += atoi(line)
	}

	fmt.Println("step1 -->", maxCalories)
}

func step2(filename string) {
	dataAsString := load(filename)

	elve := []int{}
	calories := 0

	for _, line := range strings.Split(dataAsString, "\n") {
		if line == "" {
			elve = append(elve, calories)
			calories = 0
			continue
		}
		calories += atoi(line)
	}

	elveSlice := elve[:]
	sort.Sort(sort.Reverse(sort.IntSlice(elveSlice)))

	sumCalories := 0
	for i := 0; i < 3; i++ {
		sumCalories += elveSlice[i]
	}

	fmt.Println("step2 -->", sumCalories)
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
