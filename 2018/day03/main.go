package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type rectangle struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

func newRect(id int, x int, y int, width int, height int) rectangle {
	return rectangle{id, x, y, width, height}
}

func main() {
	input := load()
	step1(input)
}

func createRectFromLine(line string) rectangle {
	rect := rectangle{}
	for i, chunk := range strings.Split(line, " ") {
		switch i {
		case 0:
			rect.id = toNumber(chunk[1:])
		case 1: // ignore
		case 2:
			cleanChunk := chunk[:len(chunk)-1]
			coords := strings.Split(cleanChunk, ",")
			rect.x = toNumber(coords[0])
			rect.y = toNumber(coords[1])
		case 3:
			size := strings.Split(chunk, "x")
			rect.width = toNumber(size[0])
			rect.height = toNumber(size[1])
		}
	}
	return rect
}

func step1(input string) {
	rectangles := []rectangle{}
	for _, line := range strings.Split(input, "\n") {
		rectangles = append(rectangles, createRectFromLine(line))
	}
	world := make([][]int, 1000)
	for i := range world {
		world[i] = make([]int, 1000)
	}
	for _, rect := range rectangles {
		for x := rect.x; x < (rect.x + rect.width); x++ {
			for y := rect.y; y < (rect.y + rect.height); y++ {
				world[x][y]++
			}
		}
	}
	count := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if world[x][y] > 1 {
				count++
			}
		}
	}
	fmt.Println("Step1", count)
	for _, rect := range rectangles {
		if isIntact(world, rect) {
			fmt.Println("Step2", rect)
			break
		}
	}
}

func isIntact(world [][]int, rect rectangle) bool {
	for x := rect.x; x < (rect.x + rect.width); x++ {
		for y := rect.y; y < (rect.y + rect.height); y++ {
			if world[x][y] != 1 {
				return false
			}
		}
	}
	return true
}

func toNumber(input string) int {
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return number
}

func load() string {
	//text, err := ioutil.ReadFile("./inputtest.txt")
	text, err := ioutil.ReadFile("./input03.txt")
	if err != nil {
		fmt.Print(err)
	}
	return string(text)
}
