package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `    |
    |  +--+
    A  |  C
F---|----E|--+
    |  |  |  D
    +B-+  +--+
`
const input2 = `    |
    |  +--+
    A  |  C
F---+----E|--+
    |  |  |  D
    +B-+  +--+
`

func TestExample1(t *testing.T) {
	assert.Equal(t, "ABCDEF", advent19(input))
}

func TestExample2(t *testing.T) {
	assert.Equal(t, "ABCDEF", advent19(input2))
}

func TestPart1(t *testing.T) {
	bytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Print(err)
	}
	assert.Equal(t, "abc", advent19(string(bytes)))
}

const empty = 0
const horizontal = 1
const vertical = 2
const cross = 3

type coord struct {
	x int
	y int
}

func (c *coord) isValid(width int, height int) bool {
	fmt.Println("isvalid ", c, c.x >= 0 && c.x < width && c.y >= 0 && c.y < height)
	return c.x >= 0 && c.x < width && c.y >= 0 && c.y < height
}

func (c *coord) move(direction int) coord {
	switch direction {
	case up:
		return c.up()
	case down:
		return c.down()
	case right:
		return c.right()
	case left:
		return c.left()
	}
	return *c
}

func (c *coord) left() coord {
	return coord{x: c.x - 1, y: c.y}
}

func (c *coord) right() coord {
	return coord{x: c.x + 1, y: c.y}
}

func (c *coord) up() coord {
	return coord{x: c.x, y: c.y - 1}
}

func (c *coord) down() coord {
	return coord{x: c.x, y: c.y + 1}
}

const down = 0
const up = 1
const left = 2
const right = 3
const blocked = -1

func padMap(rawNetworkMap string) [][]byte {
	rawLines := strings.Split(rawNetworkMap, "\n")
	width := 0
	for _, line := range rawLines {
		if len(line) > width {
			width = len(line)
		}
	}
	height := len(rawLines)
	width += 2
	height += 2
	lines := make([][]byte, height)
	for y := 0; y < height; y++ {
		lines[y] = make([]byte, width)
		for x := 0; x < width; x++ {
			if x == 0 || y == 0 || x == width-1 || y == height-1 {
				lines[y][x] = '.'
			} else if len(rawLines[y-1]) <= x-1 {
				lines[y][x] = '.'
			} else {
				lines[y][x] = rawLines[y-1][x-1]
				if lines[y][x] == ' ' {
					lines[y][x] = '.'
				}
			}
		}
	}

	drawMap(lines, coord{x: -1, y: -1}, width, height)
	return lines
}

func advent19(rawNetworkMap string) string {
	chars := []byte{}
	IsLetter := regexp.MustCompile(`[A-Z]`).MatchString
	lines := padMap(rawNetworkMap)
	direction := down
	width := len(lines[0])
	height := len(lines)
	fmt.Println(width, height)
	pos := coord{x: findX(1, '|', lines), y: 1}
	next := pos.down()
	fmt.Println(pos)
	for {
		if lines[pos.y][pos.x] == '+' {
			// drawMap(lines, pos, width, height)
			next = pos.move(direction)
			if (direction == left || direction == right) && (lines[next.y][next.x] == '-' || IsLetter(string(lines[next.y][next.x]))) {
				// fmt.Println("skip + since we horizontal and next is ", string(lines[next.y][next.x]))
			} else if (direction == up || direction == down) && (lines[next.y][next.x] == '|' || IsLetter(string(lines[next.y][next.x]))) {
				// fmt.Println("skip + since we vertical and next is ", string(lines[next.y][next.x]))
			} else {
				fmt.Println("Turning on  +")
				fmt.Println("Next is (", next.x, "/", next.y, ") -> ", string(lines[next.y][next.x]))
				fmt.Println("Direction is ", direction)
				pos, direction = chooseNextDirection(lines, direction, pos)
			}
		}
		if IsLetter(string(lines[pos.y][pos.x])) {
			fmt.Println("FOUND ", string(lines[pos.y][pos.x]))
			chars = append(chars, lines[pos.y][pos.x])
		}

		pos = pos.move(direction)

		if lines[pos.y][pos.x] == '.' {
			return string(chars)
		}
	}
	return string(chars)
}

func chooseNextDirection(lines [][]byte, direction int, pos coord) (coord, int) {
	IsLetter := regexp.MustCompile(`[A-Z]`).MatchString
	if direction == up || direction == down {
		newpos := pos.right()
		if lines[newpos.y][newpos.x] == '-' || IsLetter(string(lines[newpos.y][newpos.x])) {
			return newpos, right
		}
		newpos = pos.left()
		if lines[newpos.y][newpos.x] == '-' || IsLetter(string(lines[newpos.y][newpos.x])) {
			return newpos, left
		}
	} else if direction == right || direction == left {
		newpos := pos.up()
		if lines[newpos.y][newpos.x] == '|' || IsLetter(string(lines[newpos.y][newpos.x])) {
			return newpos, up
		}
		newpos = pos.down()
		if lines[newpos.y][newpos.x] == '|' || IsLetter(string(lines[newpos.y][newpos.x])) {
			return newpos, down
		}
	}
	return pos, blocked
}

func drawMap(lines [][]byte, pos coord, width int, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if pos.x == x && pos.y == y {
				fmt.Print("@")
			} else {
				fmt.Print(string(lines[y][x]))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func findX(y int, search byte, lines [][]byte) int {
	for x, char := range lines[y] {
		if char == search {
			return x
		}
	}
	return -1
}

// for index, line := range rawLines {
// 	missing := len(line) - width
// 	if missing > 0 {
// 		for index := 0; index < missing; index++ {
//
// 		}
// 	}
// }
//
// fmt.Println("width ", width, " height ", height)
// pos := coord{x: findX(0, '|', lines), y: 0}
// direction := down
//
// for pos.isValid(width, height) {
// 	fmt.Println("start")
// 	for pos.isValid(width, height) && len(lines[pos.y]) >= pos.x && lines[pos.y][pos.x] != '+' {
// 		fmt.Println(string(lines[pos.y][pos.x]), pos)
// 		if IsLetter(string(lines[pos.y][pos.x])) {
// 			chars = append(chars, lines[pos.y][pos.x])
// 		}
// 		switch direction {
// 		case up:
// 			pos = pos.up()
// 		case down:
// 			pos = pos.down()
// 		case right:
// 			pos = pos.right()
// 		case left:
// 			pos = pos.left()
// 		}
// 	}
// 	if !pos.isValid(width, height) || len(lines[pos.y]) < pos.x {
// 		return string(chars)
// 	}
// 	fmt.Println("plus found at ", pos.x, pos.y)
// 	if direction == right || direction == left {
// 		upPos := pos.up()
// 		fmt.Println("consider up ?")
// 		if upPos.isValid(width, height) && len(lines[upPos.y]) > pos.x && (lines[upPos.y][upPos.x] == '|' || IsLetter(string(lines[upPos.y][upPos.x]))) {
// 			fmt.Println("follow up")
// 			pos = upPos
// 			direction = up
// 		} else {
// 			downPos := pos.down()
// 			fmt.Println("consider down ?")
// 			if downPos.isValid(width, height) && len(lines[downPos.y]) > pos.x && (lines[downPos.y][downPos.x] == '|' || IsLetter(string(lines[downPos.y][downPos.x]))) {
// 				fmt.Println("follow down")
// 				pos = downPos
// 				direction = down
// 			} else {
// 				fmt.Println("no proper direction")
// 			}
// 		}
// 	} else if direction == up || direction == down {
// 		rightPos := pos.right()
// 		if rightPos.isValid(width, height) && len(lines[rightPos.y]) > pos.x && (lines[rightPos.y][rightPos.x] == '-' || IsLetter(string(lines[rightPos.y][rightPos.x]))) {
// 			fmt.Println("follow right")
// 			pos = rightPos
// 			direction = right
// 		} else {
// 			leftPos := pos.left()
// 			if leftPos.isValid(width, height) && len(lines[leftPos.y]) > pos.x && (lines[leftPos.y][leftPos.x] == '-' || IsLetter(string(lines[leftPos.y][leftPos.x]))) {
// 				fmt.Println("follow left")
// 				pos = leftPos
// 				direction = left
// 			}
// 		}
// 	}
// 	fmt.Println("end loop with ", pos)
// }
