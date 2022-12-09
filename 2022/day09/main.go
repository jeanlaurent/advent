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
	step2(load("small2.txt"))
	step2(load("input.txt"))
}

type coord struct {
	x int
	y int
}

func step1(input string) {
	visitedPlace := []coord{}
	head := coord{0, 0}
	tail := coord{0, 0}
	for _, line := range strings.Split(input, "\n") {
		moves := strings.Split(line, " ")
		direction := moves[0]
		distance := atoi(moves[1])
		for i := 0; i < distance; i++ {
			previousHead := head
			switch direction {
			case "U":
				head.y++
			case "D":
				head.y--
			case "L":
				head.x--
			case "R":
				head.x++
			}
			tail = moveTail(head, tail, previousHead)
			if !exist(tail, visitedPlace) {
				visitedPlace = append(visitedPlace, tail)
			}
		}
	}
	fmt.Println("step1 -->", len(visitedPlace))
}

func step2(input string) {
	visitedPlace := []coord{}
	rope := []coord{}
	for i := 0; i < 10; i++ {
		rope = append(rope, coord{0, 0})
	}
	for _, line := range strings.Split(input, "\n") {
		moves := strings.Split(line, " ")
		direction := moves[0]
		distance := atoi(moves[1])
		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				rope[0].y++
			case "D":
				rope[0].y--
			case "L":
				rope[0].x--
			case "R":
				rope[0].x++
			}

			for i := 1; i < 10; i++ {
				diff := coord{rope[i-1].x - rope[i].x, rope[i-1].y - rope[i].y}
				if abs(diff.x) > 1 || abs(diff.y) > 1 {
					rope[i].x += posneg(diff.x)
					rope[i].y += posneg(diff.y)
				}

			}
			if !exist(rope[9], visitedPlace) {
				visitedPlace = append(visitedPlace, rope[9])
			}
		}

	}
	fmt.Println("step2 -->", len(visitedPlace))
}

func printWorld(head, tail coord) {
	maxY := 5
	maxX := 6
	for yy := maxY - 1; yy >= 0; yy-- {
		for xx := 0; xx < maxX; xx++ {
			if xx == head.x && yy == head.y {
				fmt.Print("H")
			} else if xx == tail.x && yy == tail.y {
				fmt.Print("T")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func moveTail(newHead coord, tail coord, previousHead coord) coord {
	gap := max(abs(newHead.x-tail.x), abs(newHead.y-tail.y))
	if gap <= 1 {
		return tail
	}
	return previousHead
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func posneg(a int) int {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	} else {
		return 0
	}
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return a * -1
	}
}

func exist(target coord, list []coord) bool {
	for _, current := range list {
		if target.x == current.x && target.y == current.y {
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
