package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2
`

func TestExample1(t *testing.T) {
	assert.Equal(t, 4, advent18(input))
}

func TestPart1(t *testing.T) {
	bytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Print(err)
	}
	assert.Equal(t, 8600, advent18(string(bytes)))

}

func advent18(program string) int {
	lastFrequency := -1
	registers := map[string]int{}
	lines := strings.Split(program, "\n")
	for index := 0; index < len(lines); index++ {
		instructions := strings.Split(lines[index], " ")
		fmt.Println("index", index, " /", len(lines))
		fmt.Println(instructions)
		fmt.Println(lines)
		switch instructions[0] {
		case "set":
			registers[instructions[1]] = valueOf(instructions[2], registers)
		case "mul":
			registers[instructions[1]] *= valueOf(instructions[2], registers)
		case "add":
			registers[instructions[1]] += valueOf(instructions[2], registers)
		case "mod":
			registers[instructions[1]] %= valueOf(instructions[2], registers)
		case "jgz":
			if registers[instructions[1]] != 0 {
				offset := toInt(instructions[2])
				if offset <= 0 {
					offset -= 1
				}
				index += offset
			}
		case "rcv":
			if registers[instructions[1]] != 0 {
				return lastFrequency
			}
		case "snd":
			lastFrequency = registers[instructions[1]]
		}
		fmt.Println(registers)
		fmt.Println(lastFrequency)
	}
	return lastFrequency
}

func toInt(someString string) int {
	value, err := strconv.Atoi(someString)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return value
}

func valueOf(someString string, registers map[string]int) int {
	value, err := strconv.Atoi(someString)
	if err != nil {
		fmt.Println(err)
		return registers[someString]
	}
	return value
}
