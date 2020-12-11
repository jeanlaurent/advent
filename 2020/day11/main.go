package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	// testWorld := createWorld(load("test_input.txt"))
	// testWorld.display()
	// fmt.Println(testWorld.adjSeatDirection(3, 3))
	dataAsString := load("sample_input.txt")
	fmt.Println("step1 (sample) -->", step1(dataAsString))
	fmt.Println("step2 (sample) -->", step2(dataAsString))

	dataAsString = load("puzzle_input.txt")
	fmt.Println("step1 -->", step1(dataAsString))
	fmt.Println("step2 -->", step2(dataAsString))
}

const floor = 0
const free = 1
const occupied = 2

type seatMap struct {
	seats             [][]int
	occupiedSeatCount int
}

func (s *seatMap) display() {
	for y := 0; y < len(s.seats); y++ {
		for x := 0; x < len(s.seats[y]); x++ {
			fmt.Print(s.displayChar(x, y))
		}
		fmt.Println()
	}
}

func (s *seatMap) displayChar(x, y int) string {
	if s.seats[y][x] == floor {
		return "."
	}
	if s.seats[y][x] == free {
		return "L"
	}
	if s.seats[y][x] == occupied {
		return "#"
	}
	return "@"
}

func (s *seatMap) isValid(x, y int) bool {
	return x >= 0 && x < len(s.seats[0]) && y >= 0 && y < len(s.seats)
}

func (s *seatMap) adjSeatDirection(x, y int) int {
	occupiedCount := 0
	for ym := -1; ym <= 1; ym++ {
		for xm := -1; xm <= 1; xm++ {
			if xm == 0 && ym == 0 {
				continue
			}
			xx := x + xm
			yy := y + ym
			for s.isValid(xx, yy) && s.seats[yy][xx] == floor {
				xx += xm
				yy += ym
			}
			if s.isValid(xx, yy) && s.seats[yy][xx] == occupied {
				occupiedCount++
			}
		}
	}
	return occupiedCount
}

func (s *seatMap) ajdSeat(x, y int) int {
	occupiedCount := 0
	for ym := -1; ym <= 1; ym++ {
		for xm := -1; xm <= 1; xm++ {
			if xm == 0 && ym == 0 {
				continue
			}
			xx := x + xm
			yy := y + ym
			if !s.isValid(xx, yy) {
				continue
			}
			if s.seats[yy][xx] == occupied {
				occupiedCount++
			}
		}
	}
	return occupiedCount
}

func (s *seatMap) run() bool {
	occupiedSeatCount := 0
	hasChanged := false
	newSeats := [][]int{}
	for y := 0; y < len(s.seats); y++ {
		newArray := []int{}
		for x := 0; x < len(s.seats[y]); x++ {
			if s.seats[y][x] == floor {
				newArray = append(newArray, floor)
			} else if s.seats[y][x] == free {
				if s.ajdSeat(x, y) == 0 {
					occupiedSeatCount++
					newArray = append(newArray, occupied)
					hasChanged = true
				} else {
					newArray = append(newArray, free)
				}
			} else if s.seats[y][x] == occupied {
				if s.ajdSeat(x, y) >= 4 {
					newArray = append(newArray, free)
					hasChanged = true
				} else {
					occupiedSeatCount++
					newArray = append(newArray, occupied)
				}
			}
		}
		newSeats = append(newSeats, newArray)
	}
	s.seats = newSeats
	s.occupiedSeatCount = occupiedSeatCount
	return hasChanged
}

func (s *seatMap) run2() bool {
	occupiedSeatCount := 0
	hasChanged := false
	newSeats := [][]int{}
	for y := 0; y < len(s.seats); y++ {
		newArray := []int{}
		for x := 0; x < len(s.seats[y]); x++ {
			if s.seats[y][x] == floor {
				newArray = append(newArray, floor)
			} else if s.seats[y][x] == free {
				if s.adjSeatDirection(x, y) == 0 {
					occupiedSeatCount++
					newArray = append(newArray, occupied)
					hasChanged = true
				} else {
					newArray = append(newArray, free)
				}
			} else if s.seats[y][x] == occupied {
				if s.adjSeatDirection(x, y) >= 5 {
					newArray = append(newArray, free)
					hasChanged = true
				} else {
					occupiedSeatCount++
					newArray = append(newArray, occupied)
				}
			}
		}
		newSeats = append(newSeats, newArray)
	}
	s.seats = newSeats
	s.occupiedSeatCount = occupiedSeatCount
	return hasChanged
}

func createWorld(dataAsString string) seatMap {
	world := seatMap{seats: [][]int{}}

	for _, line := range strings.Split(dataAsString, "\n") {
		newArray := []int{}
		for i := 0; i < len(line); i++ {
			char := string(line[i])
			if char == "." {
				newArray = append(newArray, floor)
			} else if char == "L" {
				newArray = append(newArray, free)
			} else if char == "#" {
				newArray = append(newArray, occupied)
			} else {
				fmt.Println("wow ! found char ", line[i])
			}
		}
		world.seats = append(world.seats, newArray)
	}
	return world
}

func step1(dataAsString string) int {
	world := createWorld(dataAsString)
	iteration := 0
	for world.run() {
		iteration++
	}
	return world.occupiedSeatCount
}

func step2(dataAsString string) int {
	world := createWorld(dataAsString)
	iteration := 0
	for world.run2() {
		iteration++
	}
	return world.occupiedSeatCount
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
