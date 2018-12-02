package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := load()
	frequencies := []int{}
	frequency := 0
	frequencies = append(frequencies, 0)
	for {
		for _, line := range strings.Split(input, "\n") {
			change, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
				return
			}
			frequency += change
			fmt.Println("new Frequency --->", frequency, change)
			for _, oldFrequency := range frequencies {
				// fmt.Println(">  checking", frequency, oldFrequency)
				if oldFrequency == frequency {
					fmt.Println(">  first double frequency is ", frequency)
					os.Exit(1)
				}
			}
			fmt.Println(">  adding", frequency)
			frequencies = append(frequencies, frequency)
		}
	}
	fmt.Println("Final Frequency", frequency)
}

func load() string {
	text, err := ioutil.ReadFile("./input1.txt")
	if err != nil {
		fmt.Print(err)
	}
	return string(text)
}
