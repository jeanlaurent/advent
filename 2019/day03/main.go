package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

func newCoord(x, y int) coord {
	return coord{x: x, y: y}
}

func main() {
	input := load("./input.txt")
	//input := "R8,U5,L5,D3\nU7,R6,D4,L4" //30 steps
	//input := "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83" //610 steps
	//input := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7" //410 steps

	wires := [][]coord{}
	for _, line := range strings.Split(input, "\n") {
		current := newCoord(1, 1)
		wire := []coord{}
		for _, word := range strings.Split(line, ",") {
			direction := string(word[0])
			length := atoi(word[1:])
			fmt.Println(direction, length)
			switch direction {
			case "R":
				for index := 1; index <= length; index++ {
					next := newCoord(current.x+1, current.y)
					wire = append(wire, next)
					current = next
				}
			case "L":
				for index := 1; index <= length; index++ {
					next := newCoord(current.x-1, current.y)
					wire = append(wire, next)
					current = next
				}
			case "U":
				for index := 1; index <= length; index++ {
					next := newCoord(current.x, current.y+1)
					wire = append(wire, next)
					current = next
				}
			case "D":
				for index := 1; index <= length; index++ {
					next := newCoord(current.x, current.y-1)
					wire = append(wire, next)
					current = next
				}
			}
		}
		wires = append(wires, wire)
	}
	fmt.Println(wires)
	matches := []coord{}
	minDistance := 99999
	for index := 0; index < len(wires[0]); index++ {
		for index2 := 0; index2 < len(wires[1]); index2++ {
			if wires[0][index].x == wires[1][index2].x && wires[0][index].y == wires[1][index2].y {
				matches = append(matches, wires[0][index])
				distance := int(math.Abs(float64(wires[0][index].x-1)) + math.Abs(float64(wires[0][index].y-1)))
				if distance < minDistance {
					minDistance = distance
				}
			}
		}
	}

	fmt.Println("matches")
	fmt.Println(matches)
	fmt.Println("step 1 -->", minDistance)

	minSteps := 999999
	for index := 0; index < len(matches); index++ {
		steps0 := findStepFor(wires[0], matches[index])
		steps1 := findStepFor(wires[1], matches[index])
		steps := steps0 + steps1
		fmt.Println(matches[index], steps0, steps1, steps)
		if steps < minSteps {
			minSteps = steps
		}
	}
	fmt.Println("step 2 -->", minSteps)
}

func findStepFor(wire []coord, match coord) int {
	index := 0
	for {
		if wire[index].x == match.x && wire[index].y == match.y {
			return index + 1 // +1 because we don't have origin (x:1,y:1) in the wire.
		}
		index++
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
