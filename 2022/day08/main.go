package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	step1(load("small.txt"))
	step1(load("input.txt"))
	step2(load("small.txt"))
	step2(load("input.txt"))
}

func step1(input string) {
	numberOfVisibleTree := 0
	world := buildForest(input)
	for y := 0; y < len(world); y++ {
		for x := 0; x < len(world[y]); x++ {
			if checkNorth(x, y, world) {
				numberOfVisibleTree++
				continue
			}
			if checkEast(x, y, world) {
				numberOfVisibleTree++
				continue
			}
			if checkSouth(x, y, world) {
				numberOfVisibleTree++
				continue
			}
			if checkWest(x, y, world) {
				numberOfVisibleTree++
				continue
			}
		}
	}
	fmt.Println("step1 -->", numberOfVisibleTree)
}

func step2(input string) {
	world := buildForest(input)
	highestScenicScore := 0
	for y := 0; y < len(world); y++ {
		for x := 0; x < len(world[y]); x++ {
			scenicScore := 1
			scenicScore *= distanceNorth(x, y, world)
			scenicScore *= distanceEast(x, y, world)
			scenicScore *= distanceSouth(x, y, world)
			scenicScore *= distanceWest(x, y, world)
			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
	}
	fmt.Println("step1 -->", highestScenicScore)
}

func checkNorth(x, y int, world [][]int) bool {
	//y negative
	height := world[y][x]
	xx := x
	for yy := y - 1; yy >= 0; yy-- {
		if world[yy][xx] >= height {
			return false
		}
	}
	return true
}

func checkEast(x, y int, world [][]int) bool {
	//x negative
	height := world[y][x]
	yy := y
	for xx := x - 1; xx >= 0; xx-- {
		if world[yy][xx] >= height {
			return false
		}
	}
	return true
}

func checkSouth(x, y int, world [][]int) bool {
	//y positive
	height := world[y][x]
	xx := x
	for yy := y + 1; yy < len(world); yy++ {
		if world[yy][xx] >= height {
			return false
		}
	}
	return true
}

func checkWest(x, y int, world [][]int) bool {
	//x positive
	height := world[y][x]
	yy := y
	for xx := x + 1; xx < len(world[0]); xx++ {
		if world[yy][xx] >= height {
			return false
		}
	}
	return true
}

func distanceNorth(x, y int, world [][]int) int {
	//y negative
	height := world[y][x]
	xx := x
	count := 0
	for yy := y - 1; yy >= 0; yy-- {
		count++
		if world[yy][xx] >= height {
			return count
		}
	}
	return count
}
func distanceEast(x, y int, world [][]int) int {
	//x negative
	height := world[y][x]
	yy := y
	count := 0
	for xx := x - 1; xx >= 0; xx-- {
		count++
		if world[yy][xx] >= height {
			return count
		}
	}
	return count
}

func distanceSouth(x, y int, world [][]int) int {
	//y positive
	height := world[y][x]
	xx := x
	count := 0
	for yy := y + 1; yy < len(world); yy++ {
		count++
		if world[yy][xx] >= height {
			return count
		}
	}
	return count
}

func distanceWest(x, y int, world [][]int) int {
	//x positive
	height := world[y][x]
	yy := y
	count := 0
	for xx := x + 1; xx < len(world[0]); xx++ {
		count++
		if world[yy][xx] >= height {
			return count
		}
	}
	return count
}

func buildForest(input string) [][]int {
	world := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		newLine := []int{}
		for _, char := range line {
			newLine = append(newLine, atoi(string(char)))
		}
		world = append(world, newLine)
	}

	return world
}

func printWorld(world [][]int) {
	for y := 0; y < len(world); y++ {
		for x := 0; x < len(world[y]); x++ {
			fmt.Print(world[y][x])
		}
		fmt.Println()
	}
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
