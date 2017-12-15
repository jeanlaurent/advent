package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, 3, advent11("ne,ne,ne"))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 2, min(3, 2))
	assert.Equal(t, 2, min(2, 3))
	assert.Equal(t, 3, min(3, 3))
}

func TestExample2(t *testing.T) {
	assert.Equal(t, 0, advent11("ne,ne,sw,sw"))
}

func TestExample3(t *testing.T) {
	assert.Equal(t, 2, advent11("ne,ne,s,s"))
}

func TestExample4(t *testing.T) {
	assert.Equal(t, 3, advent11("se,sw,se,sw,sw"))
}

func TestStep1(t *testing.T) {
	text, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Print(err)
	}
	assert.Equal(t, 681, advent11(string(text)))
}

func cancelEachOther(dc map[string]int, el1 string, el2 string) {
	count := min(dc[el1], dc[el2])
	dc[el1] -= count
	dc[el2] -= count
}

func replace(dc map[string]int, el1 string, el2 string, el3 string) {
	count := min(dc[el1], dc[el2])
	dc[el1] -= count
	dc[el2] -= count
	dc[el3] += count
}

func advent11(input string) int {
	return reduce(strings.Split(input, ","))
}

func reduce(directions []string) int {
	directionCount := map[string]int{}
	directionCount["ne"] = count(directions, "ne")
	directionCount["sw"] = count(directions, "sw")
	directionCount["nw"] = count(directions, "nw")
	directionCount["se"] = count(directions, "se")
	directionCount["n"] = count(directions, "n")
	directionCount["s"] = count(directions, "s")

	cancelEachOther(directionCount, "n", "s")
	cancelEachOther(directionCount, "ne", "sw")
	cancelEachOther(directionCount, "nw", "se")

	replace(directionCount, "ne", "nw", "n")
	replace(directionCount, "se", "sw", "s")
	replace(directionCount, "se", "n", "ne")
	replace(directionCount, "sw", "n", "nw")
	replace(directionCount, "s", "ne", "se")
	replace(directionCount, "s", "nw", "sw")

	count := 0
	for _, v := range directionCount {
		count += v
	}
	return count

}

func count(directions []string, targetDirection string) int {
	count := 0
	for _, direction := range directions {
		if direction == targetDirection {
			count++
		}
	}
	return count
}

func min(el1 int, el2 int) int {
	if el1 < el2 {
		return el1
	}
	return el2
}
