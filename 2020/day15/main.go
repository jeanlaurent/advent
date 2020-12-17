package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(step2("0,3,6", 2020))
	fmt.Println(step2("1,0,16,5,17,4", 2020))
	fmt.Println(step2("0,3,6", 30000000))
	fmt.Println(step2("1,0,16,5,17,4", 30000000))
}

func step2(dataAsString string, end int) int {
	fmt.Println("Running", dataAsString, "over", end, "times")
	turn := 1
	memory := map[int][]int{}
	lastNumberSpoken := -1

	for _, intAsString := range strings.Split(dataAsString, ",") {
		// fmt.Println("turn:", turn)
		number := atoi(intAsString)
		memory[number] = []int{turn, -1}
		lastNumberSpoken = number
		// fmt.Println(">", lastNumberSpoken)
		// fmt.Println(memory)
		turn++
	}
	for turn <= end {
		// fmt.Println("turn:", turn)
		first, second := getFirstAndSecond(memory, lastNumberSpoken)
		// fmt.Println("lastNumberSpoken:", lastNumberSpoken, "have been seen", first, second)
		if first == -1 {
			newNumber := 0
			// fmt.Println("New Number is ----------->", newNumber)
			// fmt.Println("first time seeing")
			// fmt.Println("adding", []int{turn, -1}, "to", newNumber)
			memory[newNumber] = []int{turn, -1}
			lastNumberSpoken = newNumber
		} else if second == -1 {
			newNumber := turn - (first + 1)
			// fmt.Println("New Number is ----------->", newNumber)
			firstNew, _ := getFirstAndSecond(memory, newNumber)
			// fmt.Println("adding", []int{turn, -1}, "to", newNumber)
			memory[newNumber] = []int{turn, firstNew}
			lastNumberSpoken = newNumber
		} else {
			newNumber := first - second
			// fmt.Println("New Number is ----------->", newNumber)
			firstNew, _ := getFirstAndSecond(memory, newNumber)
			// fmt.Println("adding", []int{turn, -1}, "to", newNumber)
			memory[newNumber] = []int{turn, firstNew}
			lastNumberSpoken = newNumber
		}

		// fmt.Println(memory)
		// fmt.Println()
		turn++
	}
	return lastNumberSpoken
}

func getFirstAndSecond(memory map[int][]int, lastNumberSpoken int) (int, int) {
	firstAndSecond, exists := memory[lastNumberSpoken]
	if !exists {
		return -1, -1
	}
	return firstAndSecond[0], firstAndSecond[1]
}

func step1(dataAsString string, end int) int {
	turn := 1
	lane := []int{}
	for _, intAsString := range strings.Split(dataAsString, ",") {
		number := atoi(intAsString)
		lane = append(lane, number)
		turn++
	}

	// fmt.Println(lane)
	for turn <= end {
		// fmt.Println("turn:", turn)
		lastNumberSpoken := lane[turn-2]
		// fmt.Println("lastNumberSpoken:", lastNumberSpoken)
		first, second := last2IndexOf(lane, lastNumberSpoken)
		// fmt.Prsintln(first, second)
		if first == -1 {
			// fmt.Println("lastNumberSpoken:", lastNumberSpoken, " is brand new")
			lane = append(lane, 0)
		} else if second == -1 {
			if first+1 == turn-1 {
				// fmt.Println("lastNumberSpoken:", lastNumberSpoken, " is new from last turn")
				lane = append(lane, 0)
			} else {
				// fmt.Println("lastNumberSpoken:", lastNumberSpoken, " has been seen once")
				lane = append(lane, turn-(first+1))
			}
		} else {
			// fmt.Println("lastNumberSpoken:", lastNumberSpoken, " has been seen twice")
			lane = append(lane, first-second)
		}
		// fmt.Println(lane)
		turn++
	}
	return lane[len(lane)-1]
}

func lastIndexOf(array []int, target int) int {
	for index := len(array) - 1; index >= 0; index-- {
		if array[index] == target {
			return index
		}
	}
	return -1
}

func last2IndexOf(array []int, target int) (int, int) {
	firstIndex := -1
	secondIndex := -1
	for index := len(array) - 1; index >= 0; index-- {
		if array[index] == target {
			if firstIndex == -1 {
				firstIndex = index
			} else {
				return firstIndex, index
			}
		}
	}
	return firstIndex, secondIndex
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
