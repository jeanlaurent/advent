package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const north = 0
const east = 90
const south = 180
const west = 270

type step struct {
	order string
	steps int
}

type ferry struct {
	x         int
	y         int
	direction int
}

type coord struct {
	x int
	y int
}

func (f *ferry) move(direction int, distance int) {
	switch direction {
	case north:
		f.y += distance
	case east:
		f.x += distance
	case west:
		f.x -= distance
	case south:
		f.y -= distance
	}
}

func (f *ferry) turn(angle int) {
	f.direction = f.direction + angle
	if f.direction < 0 {
		f.direction += 360
	}
	if f.direction >= 360 {
		f.direction -= 360
	}
}

func main() {
	dataAsString := load("sample_input.txt")
	fmt.Println("step1 (sample) -->", step1(dataAsString))
	fmt.Println("step2 (sample) -->", step2(dataAsString))

	dataAsString = load("puzzle_input.txt")
	fmt.Println("step1 -->", step1(dataAsString))
	fmt.Println("step2 -->", step2(dataAsString))
}

func stringToSteps(dataAsString string) []step {
	steps := []step{}
	for _, line := range strings.Split(dataAsString, "\n") {
		currentStep := step{order: string(line[0]), steps: atoi(line[1:])}
		steps = append(steps, currentStep)
	}
	return steps
}

func wayPointTurn(x, y, angle int) (int, int) {
	if angle < 0 {
		angle = 360 + angle
	}
	switch angle {
	case 90:
		return y, x * -1
	case 180:
		return x * -1, y * -1
	case 270:
		return y * -1, x
	}
	return x, y
}

func step2(dataAsString string) int {
	steps := stringToSteps(dataAsString)
	boat := coord{x: 0, y: 0}
	waypoint := coord{x: 10, y: 1}
	for _, step := range steps {
		switch step.order {
		case "F":
			boat.x = boat.x + waypoint.x*step.steps
			boat.y = boat.y + waypoint.y*step.steps
		case "L":
			waypoint.x, waypoint.y = wayPointTurn(waypoint.x, waypoint.y, step.steps*-1)
		case "R":
			waypoint.x, waypoint.y = wayPointTurn(waypoint.x, waypoint.y, step.steps)
		case "N":
			waypoint.y += step.steps
		case "E":
			waypoint.x += step.steps
		case "W":
			waypoint.x -= step.steps
		case "S":
			waypoint.y -= step.steps
		}
		// fmt.Println("--> ", step)
		// fmt.Println("boat", boat)
		// fmt.Println("waypoint", waypoint)
		// fmt.Println()
	}
	return abs(boat.x) + abs(boat.y)
}

func step1(dataAsString string) int {
	steps := stringToSteps(dataAsString)
	boat := ferry{x: 0, y: 0, direction: east}
	for _, step := range steps {
		switch step.order {
		case "F":
			boat.move(boat.direction, step.steps)
		case "L":
			boat.turn(step.steps * -1)
		case "R":
			boat.turn(step.steps)
		case "N":
			boat.move(north, step.steps)
		case "E":
			boat.move(east, step.steps)
		case "W":
			boat.move(west, step.steps)
		case "S":
			boat.move(south, step.steps)
		}
	}
	return abs(boat.x) + abs(boat.y)
}

// ----

func abs(i int) int {
	if i > 0 {
		return i
	}
	return i * -1
}

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
