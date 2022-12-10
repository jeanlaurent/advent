package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	step1(load("small.txt"))
	step1(load("input.txt"))
	step2(load("small.txt"))
	step2(load("input.txt"))
}

func step1(input string) {
	mark := 40
	x := 1
	cycle := 0
	signalStrength := 0

	for _, line := range strings.Split(input, "\n") {
		//finally trying scanf per dgageot/ulrich tweet
		var cmd string
		var val int
		fmt.Sscanf(line, "%s %d", &cmd, &val)

		//noop or in between addx
		cycle++
		if (cycle+mark/2)%mark == 0 {
			signalStrength += x * cycle
		}
		if cmd == "addx" {
			cycle++
			if (cycle+mark/2)%mark == 0 {
				signalStrength += x * cycle
			}
			x += val
		}
	}
	fmt.Println("step1 -->", signalStrength)
}

func step2(input string) {
	screen := ""
	mark := 40
	x := 1
	cycle := 0

	for _, line := range strings.Split(input, "\n") {
		//finally trying scanf per dgageot/ulrich tweet
		var cmd string
		var val int
		fmt.Sscanf(line, "%s %d", &cmd, &val)

		// before instruction
		screen += drawPixel(cycle, mark, x)
		//noop or in between addx
		cycle++
		if cmd == "addx" {
			screen += drawPixel(cycle, mark, x)
			cycle++
			x += val
		}
	}
	fmt.Println("step2 -->")
	fmt.Println(screen)
}

func drawPixel(cycle, mark, x int) string {
	pixel := ""
	if (cycle%mark-1) <= x && x <= cycle%mark+1 {
		pixel = "#"
	} else {
		pixel = " "
	}
	if cycle%mark == mark-1 {
		pixel += "\n"
	}
	return pixel
}

func load(filename string) string {
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	return string(text)
}
