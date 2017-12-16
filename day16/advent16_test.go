package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

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

func TestPart2(t *testing.T) {
	bytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Print(err)
	}
	assert.Equal(t, "doeaimlbnpjchfkg", advent16Part2("abcdefghijklmnop", string(bytes[0:len(bytes)-1]), 1000000000))
}

func advent16(dancers string, danceMovesAsString string) string {
	return advent16Part2(dancers, danceMovesAsString, 1)
}

func advent16Part2(dancers string, danceMovesAsString string, iteration int) string {
	moves := strings.Split(danceMovesAsString, ",")
	dancersArray := make([]string, len(dancers))
	for loopCount := 0; loopCount < iteration; loopCount++ {
		if (loopCount%1000 == 0) && loopCount > 0 {
			fmt.Println(loopCount)
		}
		for _, move := range moves {
			if string(move[0]) == "s" {
				span, err := strconv.Atoi(string(move[1:]))
				if err != nil {
					fmt.Println(err)
					return "error"
				}
				for index := 0; index < len(dancers); index++ {
					dancersArray[(index+span)%len(dancers)] = string(dancers[index])
				}
			}
			if string(move[0]) == "x" {
				values := strings.Split(string(move[1:]), "/")
				left, err := strconv.Atoi(values[0])
				if err != nil {
					fmt.Println(err)
					return "error"
				}
				right, err := strconv.Atoi(values[1])
				if err != nil {
					fmt.Println(err)
					return "error"
				}
				for index := 0; index < len(dancers); index++ {
					if index == left {
						dancersArray[index] = string(dancers[right])
					} else if index == right {
						dancersArray[index] = string(dancers[left])
					} else {
						dancersArray[index] = string(dancers[index])
					}
				}
			}
			if string(move[0]) == "p" {
				values := strings.Split(string(move[1:]), "/")
				left := strings.Index(dancers, string(values[0]))
				right := strings.Index(dancers, string(values[1]))
				for index := 0; index < len(dancers); index++ {
					if index == left {
						dancersArray[index] = string(dancers[right])
					} else if index == right {
						dancersArray[index] = string(dancers[left])
					} else {
						dancersArray[index] = string(dancers[index])
					}
				}
			}
			newDancers := ""
			for index := range dancersArray {
				newDancers += dancersArray[index]
			}
			dancers = newDancers
		}
	}
	return dancers
}
