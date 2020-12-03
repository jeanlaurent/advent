package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const tree = 1
const space = 0

func main() {
	dataAsString := load("puzzle_input.txt")
	//dataAsString := load("sample_input.txt")
	world := [][]int{}
	for _, line := range strings.Split(dataAsString, "\n") {
		worldLine := []int{}
		for _, char := range line {
			if string(char) == "#" {
				worldLine = append(worldLine, tree)
			} else {
				worldLine = append(worldLine, space)
			}
		}
		world = append(world, worldLine)
	}

	fmt.Println("step1 -->", foundTreeForSlope(1, 3, world))

	tree1 := foundTreeForSlope(1, 1, world)
	tree2 := foundTreeForSlope(1, 3, world)
	tree3 := foundTreeForSlope(1, 5, world)
	tree4 := foundTreeForSlope(1, 7, world)
	tree5 := foundTreeForSlope(2, 1, world)
	fmt.Println("step2 -->", tree1*tree2*tree3*tree4*tree5)
}

func foundTreeForSlope(xIncrement, yIncrement int, world [][]int) int {
	maxX := len(world)
	maxY := len(world[0])

	treeEncountered := 0
	x := 0
	y := 0

	for x < maxX-1 {
		y = (y + yIncrement) % maxY
		x = x + xIncrement
		if world[x][y] == tree {
			treeEncountered++
		}
	}

	return treeEncountered
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
