package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type node struct {
	number int
	name   string
}

func main() {
	dataAsString := load("puzzle_input.txt")
	//dataAsString := load("sample_input.txt")
	//dataAsString := load("sample_input2.txt")

	fmt.Println("step1 -->", step1(dataAsString))
	fmt.Println("step2 -->", step2(dataAsString))
}

func step2(dataAsString string) int {
	return findBags(node{name: "shiny gold", number: 1}, loadInputAsMap2(dataAsString)) - 1 // let's remove the shiny gold root node
}

func findBags(parentNode node, nodeMap map[string][]node) int {
	children := nodeMap[parentNode.name]
	sumOfChildrenBags := 0
	for _, child := range children {
		sumOfChildrenBags += findBags(child, nodeMap)
	}
	numberOfBags := parentNode.number + parentNode.number*sumOfChildrenBags
	return numberOfBags
}

func loadInputAsMap2(dataAsString string) map[string][]node {
	dic := map[string][]node{}
	for _, line := range strings.Split(dataAsString, "\n") {
		parts := strings.Split(line, " contain ")
		parent := strings.TrimSpace(parts[0][0 : len(parts[0])-4])
		childArray := []node{}
		for _, childLine := range strings.Split(parts[1], ",") {
			if childLine != "no other bags." {
				childParts := strings.Split(strings.TrimSpace(childLine), " ")
				child := childParts[1] + " " + childParts[2]
				childArray = append(childArray, node{name: strings.TrimSpace(child), number: atoi(childParts[0])})
			}
		}
		dic[parent] = childArray
	}
	return dic
}

func step1(dataAsString string) int {
	dic := loadInputAsMap(dataAsString)
	sum := 0
	listToFind := []string{"shiny gold"}
	found := []string{}
	for len(listToFind) > 0 {
		toFind := listToFind[0]
		for key, values := range dic {
			for _, value := range values {
				if toFind == value && !exist(found, key) {
					listToFind = append(listToFind, key)
					found = append(found, key)
					sum++
					break
				}
			}
		}
		listToFind = listToFind[1:]
	}
	return sum
}

func loadInputAsMap(dataAsString string) map[string][]string {
	dic := map[string][]string{}
	for _, line := range strings.Split(dataAsString, "\n") {
		parts := strings.Split(line, " contain ")
		parent := strings.TrimSpace(parts[0][0 : len(parts[0])-4])
		childArray := []string{}
		for _, childLine := range strings.Split(parts[1], ",") {
			if childLine != "no other bags." {
				childParts := strings.Split(strings.TrimSpace(childLine), " ")
				child := childParts[1] + " " + childParts[2]
				childArray = append(childArray, strings.TrimSpace(child))
			}
		}
		dic[parent] = childArray
	}
	return dic
}

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
