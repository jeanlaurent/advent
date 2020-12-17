package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("step1 (sample) -->", step1(load("sample_input.txt")))
	fmt.Println("step2 (sample2) -->", step2(load("sample2_input.txt")))

	fmt.Println("step1 -->", step1(load("puzzle_input.txt")))
	fmt.Println("step2 -->", step2(load("puzzle_input.txt")))

}

func step2(dataAsString string) int {
	registers := map[int]int{}
	originalMask := ""
	for _, line := range strings.Split(dataAsString, "\n") {
		if line[0:4] == "mask" {
			originalMask = line[7:]
			fmt.Println("Found new mask")
			fmt.Println(originalMask)
			fmt.Println("====")
		} else {
			parts := strings.Split(line, "=")

			address := atoi(parts[0][4 : len(parts[0])-2])
			value := atoi(strings.TrimSpace(parts[1]))

			newMask := ""
			fmt.Printf("Address: %036b (%d)\n", address, address)
			for bitPos := 0; bitPos < len(originalMask); bitPos++ {
				bitMask := string(originalMask[bitPos])
				if bitMask == "0" {
					if hasBit(address, uint(len(originalMask)-bitPos-1)) {
						newMask += "1"
					} else {
						newMask += "0"
					}
				} else {
					newMask += bitMask
				}
			}
			fmt.Printf("Mask   : %s\n", newMask)

			adresses := []string{} // lol generate int rather than string here, divide code by 2.
			for i := 0; i < len(newMask); i++ {
				if string(newMask[i]) == "X" {
					if len(adresses) == 0 {
						adresses = []string{"0", "1"}
					} else {
						newAddress := []string{}
						for _, add := range adresses {
							newAddress = append(newAddress, add+"0")
							newAddress = append(newAddress, add+"1")
						}
						adresses = newAddress
					}
				} else {
					if len(adresses) == 0 {
						adresses = []string{string(newMask[i])}
					} else {
						newAddress := []string{}
						for _, add := range adresses {
							newAddress = append(newAddress, add+string(newMask[i]))
						}
						adresses = newAddress
					}

				}
			}
			for _, add := range adresses {
				fmt.Print("       : ", add)
				addressAsInt := 0
				for bitPos := 0; bitPos < len(add); bitPos++ {
					val := string(add[len(add)-bitPos-1])
					if val == "1" {
						addressAsInt = setBit(addressAsInt, uint(bitPos))
					} else if val == "0" {
						addressAsInt = clearBit(addressAsInt, uint(bitPos))
					}
				}
				fmt.Println(" (", addressAsInt, ")")
				registers[addressAsInt] = value
			}

		}
	}

	sum := 0
	for _, v := range registers {
		sum += v
	}
	return int(sum)
}

func step1(dataAsString string) int {
	registers := map[int]int{}
	mask := ""
	for _, line := range strings.Split(dataAsString, "\n") {
		if line[0:4] == "mask" {
			mask = line[7:]
			// fmt.Println(mask)
		} else {
			// fmt.Println("====")
			parts := strings.Split(line, "=")
			offset := atoi(strings.TrimSpace(parts[1]))
			address := atoi(parts[0][4 : len(parts[0])-2])
			// fmt.Println("new offset", address, offset)
			value := offset
			// fmt.Printf("%036b (%d)\n", value, value)
			// fmt.Println(mask)
			for bitPos := 0; bitPos < len(mask); bitPos++ {
				val := string(mask[len(mask)-bitPos-1])
				if val == "1" {
					value = setBit(value, uint(bitPos))
				} else if val == "0" {
					value = clearBit(value, uint(bitPos))
				}
			}
			// fmt.Printf("%036b (%d)\n", value, value)
			registers[address] = value
			// fmt.Println(registers)
		}
	}
	sum := 0
	for _, v := range registers {
		sum += v
	}
	return int(sum)
}

func setBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}

func clearBit(n int, pos uint) int {
	mask := ^(1 << pos)
	n &= mask
	return n
}

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
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
