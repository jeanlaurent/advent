package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, "eabcd", advent16("abcde", "s1"))
	assert.Equal(t, "deabc", advent16("abcde", "s2"))
}

func TestExample2(t *testing.T) {
	assert.Equal(t, "eabdc", advent16("eabcd", "x3/4"))
}

func TestExample3(t *testing.T) {
	assert.Equal(t, "baedc", advent16("eabdc", "pe/b"))
}

func TestExampleAll(t *testing.T) {
	assert.Equal(t, "baedc", advent16("abcde", "s1,x3/4,pe/b"))
}

func TestExamplePart2(t *testing.T) {
	assert.Equal(t, "ceadb", advent16Part2("abcde", "s1,x3/4,pe/b", 2))
}

func TestPart1(t *testing.T) {
	bytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Print(err)
	}
	assert.Equal(t, "doeaimlbnpjchfkg", advent16("abcdefghijklmnop", string(bytes[0:len(bytes)-1])))
}

// func TestPart2(t *testing.T) {
// 	bytes, err := ioutil.ReadFile("./input.txt")
// 	if err != nil {
// 		fmt.Print(err)
// 	}
// 	assert.Equal(t, "doeaimlbnpjchfkg", advent16Part2("abcdefghijklmnop", string(bytes[0:len(bytes)-1]), 1000000000))
// }

func advent16(dancers string, danceMovesAsString string) string {
	return advent16Part2(dancers, danceMovesAsString, 1)
}

const spin = 1
const exchange = 2
const partner = 3

type Move struct {
	orderType int
	left      int
	right     int
	leftChar  string
	rightChar string
}

func parseInput(danceMovesAsString string) []Move {
	stringMoves := strings.Split(danceMovesAsString, ",")
	moves := make([]Move, len(stringMoves))
	for index, stringMove := range stringMoves {
		if string(stringMove[0]) == "s" {
			span, err := strconv.Atoi(string(stringMove[1:]))
			if err != nil {
				fmt.Println(err)
			}
			moves[index] = Move{orderType: spin, left: span}
		}
		if string(stringMove[0]) == "x" {
			values := strings.Split(string(stringMove[1:]), "/")
			left, err := strconv.Atoi(values[0])
			if err != nil {
				fmt.Println(err)
			}
			right, err := strconv.Atoi(values[1])
			if err != nil {
				fmt.Println(err)
			}
			moves[index] = Move{orderType: exchange, left: left, right: right}
		}
		if string(stringMove[0]) == "p" {
			values := strings.Split(string(stringMove[1:]), "/")
			moves[index] = Move{orderType: partner, leftChar: values[0], rightChar: values[1]}
		}
	}
	return moves
}

func advent16Part2(dancers string, danceMovesAsString string, iteration int) string {
	moves := parseInput(danceMovesAsString)
	dancersArray := make([]string, len(dancers))
	for index := 0; index < len(dancers); index++ {
		dancersArray[index] = string(dancers[index])
	}
	start := time.Now()
	spot := start
	for loopCount := 0; loopCount < iteration; loopCount++ {
		if loopCount%1000 == 0 && loopCount != 0 {
			diff := time.Since(spot)
			diffStart := time.Since(start)
			fmt.Println(loopCount, diff, diffStart)
			spot = time.Now()
		}

		for _, move := range moves {
			switch move.orderType {
			case spin:
				newDancersArray := make([]string, len(dancers))
				for index := 0; index < len(dancers); index++ {
					newDancersArray[(index+move.left)%len(dancers)] = dancersArray[index]
				}
				dancersArray = newDancersArray
			case exchange:
				swap(dancersArray, move.left, move.right)
			case partner:
				left := find(dancersArray, move.leftChar)
				right := find(dancersArray, move.rightChar)
				swap(dancersArray, left, right)
			}
		}
	}
	newDancers := ""
	for index := range dancersArray {
		newDancers += dancersArray[index]
	}
	return newDancers
}

func swap(slice []string, left int, right int) {
	swap := slice[right]
	slice[right] = slice[left]
	slice[left] = swap
}

func find(slice []string, target string) int {
	for index, element := range slice {
		if target == element {
			return index
		}
	}
	return -1
}
