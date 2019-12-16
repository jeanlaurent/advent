package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	computer := newIntCodeComputer(1, loadAndParse())
	keycode := computer.boost()
	fmt.Println("step1 -->:", keycode)
	computer = newIntCodeComputer(2, loadAndParse())
	keycode = computer.boost()
	fmt.Println("step2 -->:", keycode)

}

type intCodeComputer struct {
	input        int
	output       int64
	intcodes     []int64
	memory       map[int64]int64
	relativeBase int64
}

func newIntCodeComputer(input int, intcodes []int64) intCodeComputer {
	return intCodeComputer{input, -1, intcodes, make(map[int64]int64), 0}
}

func (c *intCodeComputer) boost() int64 {
	index := int64(0)
	for {
		// readOpCode and Param
		opcode := c.intcodes[index] % 100
		param1 := 0
		if c.intcodes[index] > 99 {
			param1 = getDigit(c.intcodes[index], 3)
		}
		param2 := 0
		if c.intcodes[index] > 999 {
			param2 = getDigit(c.intcodes[index], 4)
		}
		param3 := 0
		if c.intcodes[index] > 9999 {
			param3 = getDigit(c.intcodes[index], 5)
		}
		switch opcode {
		case 1: // addition
			firstParameter := c.getValueOrPointer(param1, c.intcodes[index+1])
			secondParameter := c.getValueOrPointer(param2, c.intcodes[index+2])
			fmt.Println("Add", firstParameter, secondParameter, " to ", c.intcodes[index+3])
			c.writeToMemory(c.getPointer(param3, c.intcodes[index+3]), firstParameter+secondParameter)
			index += 4
		case 2: // multiplication
			firstParameter := c.getValueOrPointer(param1, c.intcodes[index+1])
			secondParameter := c.getValueOrPointer(param2, c.intcodes[index+2])
			fmt.Println("Mult", firstParameter, secondParameter, " to ", c.intcodes[index+3])
			c.writeToMemory(c.getPointer(param3, c.intcodes[index+3]), firstParameter*secondParameter)
			index += 4
		case 3: //
			firstParameter := c.getPointer(param1, c.intcodes[index+1])
			fmt.Println("Store input", int64(c.input), "to", firstParameter)
			c.writeToMemory(firstParameter, int64(c.input))
			index += 2
		case 4: // output
			fmt.Println("Output", c.intcodes[index], param1, c.intcodes[index+1])
			c.output = c.getValueOrPointer(param1, c.intcodes[index+1])
			fmt.Println(">>>", c.output)
			index += 2
		case 5: // jump if true
			fmt.Println("jump if true")
			condition := c.getValueOrPointer(param1, c.intcodes[index+1])
			if condition != 0 {
				fmt.Println("jump !")
				index = c.getValueOrPointer(param2, c.intcodes[index+2])
			} else {
				index += 3
			}
		case 6: // jump if false
			fmt.Println("jump if false")
			condition := c.getValueOrPointer(param1, c.intcodes[index+1])
			if condition == 0 {
				index = c.getValueOrPointer(param2, c.intcodes[index+2])
				fmt.Println("jump !")
			} else {
				index += 3
			}
		case 7: // less than
			fmt.Println("less than")
			firstParameter := c.getValueOrPointer(param1, c.intcodes[index+1])
			secondParameter := c.getValueOrPointer(param2, c.intcodes[index+2])
			if firstParameter < secondParameter {
				c.writeToMemory(c.getPointer(param3, c.intcodes[index+3]), 1)
			} else {
				c.writeToMemory(c.getPointer(param3, c.intcodes[index+3]), 0)
			}
			index += 4
		case 8: // equal than
			fmt.Println("equal than")
			firstParameter := c.getValueOrPointer(param1, c.intcodes[index+1])
			secondParameter := c.getValueOrPointer(param2, c.intcodes[index+2])
			if firstParameter == secondParameter {
				c.writeToMemory(c.getPointer(param3, c.intcodes[index+3]), 1)
			} else {
				c.writeToMemory(c.getPointer(param3, c.intcodes[index+3]), 0)
			}
			index += 4
		case 9: // relative base
			firstParameter := c.getValueOrPointer(param1, c.intcodes[index+1])
			fmt.Println("relative base is ", c.relativeBase, " adding ", firstParameter)
			c.relativeBase += firstParameter
			index += 2
		case 99:
			return c.output
		default:
			fmt.Println("invalid intcode", c.intcodes[index], "at", index)
			os.Exit(-1)
		}
	}
}

func (c *intCodeComputer) writeToMemory(address int64, value int64) {
	fmt.Println("Write", value, " to ", address)
	if address > int64(len(c.intcodes)) {
		c.memory[address] = value
	} else {
		c.intcodes[address] = value
	}
}

func (c *intCodeComputer) readFromMemory(address int64) int64 {
	value := int64(0)
	if address > int64(len(c.intcodes)) {
		value = c.memory[address]
	} else {
		value = c.intcodes[address]
	}
	fmt.Println("read", value, "from ", address)
	return value
}

func (c *intCodeComputer) getValueOrPointer(paramMode int, value int64) int64 {
	switch paramMode {
	case 0: // Position Mode
		return c.readFromMemory(value)
	case 1: // Immediate Mode
		return value
	case 2: // Relative Mode
		return c.readFromMemory(c.relativeBase + value)
	default:
		fmt.Println("invalid paramMode", paramMode)
		os.Exit(-1)
		return -1 // non reachable but make compiler happy.
	}
}

func (c *intCodeComputer) getPointer(paramMode int, value int64) int64 {
	switch paramMode {
	case 0: // Position Mode
		return value
	case 1: // Immediate Mode
		fmt.Println("invalid paramMode[1]", paramMode)
		os.Exit(-1)
		return -1 // non reachable but make compiler happy.
	case 2: // Relative Mode
		return c.relativeBase + value
	default:
		fmt.Println("invalid paramMode", paramMode)
		os.Exit(-1)
		return -1 // non reachable but make compiler happy.
	}
}

// utilities
func loadAndParse() []int64 {
	ops := []int64{}
	dataAsString := load("input.txt")
	//dataAsString := "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"
	//dataAsString := "1102,34915192,34915192,7,4,7,99,0"
	//dataAsString := "104,1125899906842624,99"
	for _, opsAsString := range strings.Split(dataAsString, ",") {
		ops = append(ops, int64(atoi(opsAsString)))
	}
	return ops
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

func getDigit(num int64, place int) int {
	r := int(num) % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}
