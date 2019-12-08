package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const width = 25
const height = 6
const layerMax = 100
const layerlength = width * height

type layerstat struct {
	stat []int
}

func main() {
	step1(load("input.txt"))
	step2(load("input.txt"))
}

func step2(dataAsString string) {
	layers := [layerMax][height][width]int{}
	for pos, char := range dataAsString {
		layerNum := pos / layerlength
		x := pos % width
		y := (pos - layerNum*layerlength) / width
		layers[layerNum][y][x] = atoi(string(char))
	}
	fmt.Println("step2 -->")
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			for layer := 0; layer < layerMax; layer++ {
				if layers[layer][y][x] == 1 {
					fmt.Print("#")
					break
				}
				if layers[layer][y][x] == 0 {
					fmt.Print(" ")
					break
				}
			}
		}
		fmt.Println()
	}
}

func step1(dataAsString string) {
	statistics := make(map[int]layerstat)
	maxLayer := -1
	for pos, char := range dataAsString {
		layerNum := pos / layerlength
		num := atoi(string(char))
		if len(statistics[layerNum].stat) == 0 {
			statistics[layerNum] = layerstat{[]int{0, 0, 0}}
		}
		statistics[layerNum].stat[num]++
		maxLayer = layerNum
	}
	minZero := 9999
	minLayer := -1
	for index := 0; index < maxLayer; index++ {
		if statistics[index].stat[0] < minZero {
			minZero = statistics[index].stat[0]
			minLayer = index
		}
	}
	fmt.Println("step1 --> ", statistics[minLayer].stat[1]*statistics[minLayer].stat[2])

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
