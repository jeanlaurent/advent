package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	dataAsString := load("puzzle_input.txt")
	maxSeatID := -1
	planePlan := [128][8]int{}
	for y := 0; y < 128; y++ {
		for x := 0; x < 7; x++ {
			planePlan[y][x] = 0
		}
	}

	for _, line := range strings.Split(dataAsString, "\n") {
		seatID, y, x := findSeatId(line)
		planePlan[y][x] = 1
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	fmt.Println("step1 ->", maxSeatID)

	// printPlanePlan(planePlan)
	previousRowFull := false
	index := -1
	for y := 0; y < 128; y++ {
		rowFull := true
		for x := 0; x < 8; x++ {
			if planePlan[y][x] == 0 {
				rowFull = false
				index = x
				break
			}
		}
		if !rowFull && previousRowFull {
			fmt.Println("step2-->", computeSeatId(y, index))
			break
		}
		previousRowFull = rowFull
	}

}

func printPlanePlan(plan [128][8]int) {
	for y := 0; y < 128; y++ {
		fmt.Print(y, "\t")
		for x := 0; x < 8; x++ {
			if plan[y][x] == 1 {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func findSeatId(ticket string) (int, int, int) {
	row := findRow(ticket)
	col := findColumn(ticket)
	return computeSeatId(row, col), row, col
}

func computeSeatId(row, col int) int {
	return row*8 + col
}

func findColumn(ticket string) int {
	min := 0
	max := 7
	for i := 7; i < 10; i++ {
		char := ticket[i]
		if string(char) == "L" {
			max = (max-min)/2 + min
		} else {
			min = (max-min)/2 + min + 1
		}
	}
	return max
}

func findRow(ticket string) int {
	min := 0
	max := 127
	for i := 0; i < 7; i++ {
		char := ticket[i]
		if string(char) == "F" {
			max = (max-min)/2 + min
		} else {
			min = (max-min)/2 + min + 1
		}
	}
	return max
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
