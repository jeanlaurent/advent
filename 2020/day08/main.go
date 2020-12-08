package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	command  string
	modifier int
	visited  bool
}

type program struct {
	acc     int
	lines   []instruction
	current int
}

func (p *program) step() {
	p.lines[p.current].visited = true
	switch p.lines[p.current].command {
	case "nop":
		p.current++
	case "jmp":
		p.current += p.lines[p.current].modifier
	case "acc":
		p.acc += p.lines[p.current].modifier
		p.current++
	}
}

func (p *program) reset() {
	p.acc = 0
	p.current = 0
	for i := 0; i < len(p.lines); i++ {
		p.lines[i].visited = false
	}
}

func (p *program) print() {
	fmt.Println("acc", p.acc)
	fmt.Println("cur", p.current)
	for i := 0; i < len(p.lines); i++ {
		fmt.Println(i, p.lines[i].command, p.lines[i].modifier)
	}
	fmt.Println("-----")
}

const notTerminated = 0
const infiniteLoop = 1
const normalTermination = 2

func (p *program) isTerminated() int {
	if p.current >= len(p.lines) {
		return normalTermination
	}
	if p.lines[p.current].visited == true {
		return infiniteLoop
	}
	return notTerminated
}

func main() {
	dataAsString := load("puzzle_input.txt")
	//dataAsString := load("sample_input.txt")

	fmt.Println("step1 -->", step1(dataAsString))
	fmt.Println("step2 -->", step2(dataAsString))
}

func step2(dataAsString string) int {
	prog := readInputIntoProgram(dataAsString)
	swapLine := -1
	for lineNumber, instruction := range prog.lines {
		prog = readInputIntoProgram(dataAsString)
		if lineNumber > swapLine && instruction.command == "nop" || instruction.command == "jmp" {
			if instruction.command == "jmp" {
				swapLine = lineNumber
				prog.lines[swapLine].command = "nop"
			} else if instruction.command == "nop" {
				swapLine = lineNumber
				prog.lines[swapLine].command = "nop"
			}
			for prog.isTerminated() == notTerminated {
				prog.step()
			}
			if prog.isTerminated() == normalTermination {
				return prog.acc
			}
		}
	}
	return -1
}

func step1(dataAsString string) int {
	prog := readInputIntoProgram(dataAsString)
	for prog.isTerminated() != infiniteLoop {
		prog.step()
	}
	return prog.acc
}

func readInputIntoProgram(dataAsString string) program {
	prog := program{acc: 0, lines: []instruction{}, current: 0}
	for _, line := range strings.Split(dataAsString, "\n") {
		parts := strings.Split(line, " ")
		inst := instruction{command: parts[0], modifier: atoi(parts[1]), visited: false}
		prog.lines = append(prog.lines, inst)
	}
	return prog
}

// ----

func exist(array []string, item string) bool {
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
