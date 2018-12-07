package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type node struct {
	origin      string
	destination string
}

func main() {
	input := load()
	step1 := ""
	origins := []string{}
	destinations := []string{}
	nodes := []node{}
	for _, line := range strings.Split(input, "\n") {
		origin := string(line[5])
		origins = append(origins, origin)
		destination := string(line[36])
		destinations = append(destinations, destination)
		fmt.Println(string(origin), "->", string(destination), ";")
		nodes = append(nodes, node{origin, destination})
	}
	roots := findRoots(origins, destinations)
	root := ""
	step1 = ""
	paths := roots

	loop := 0
	for len(nodes) > 0 {
		loop++
		fmt.Println(step1, paths, nodes)
		fmt.Println("<<<<<<< round", loop, ">>>>>>>>")
		fmt.Println("root", root, "paths", paths, "step1", step1)
		newNodes := []node{}
		fmt.Println("looking at node starting with", root)
		for _, aNode := range nodes {
			if aNode.origin == root {
				destinationAlreadyExist := false
				for _, path := range paths {
					if path == aNode.destination {
						destinationAlreadyExist = true
						break
					}
				}
				if !destinationAlreadyExist {
					fmt.Println(aNode)
					fmt.Println(" adding", aNode.destination, "to path")
					fmt.Println(" removing", aNode, "from available nodes")
					paths = append(paths, aNode.destination)
				} else {
					fmt.Println("Not adding", aNode.destination, "since it already exist")
				}
			} else {
				newNodes = append(newNodes, aNode)
			}
		}
		nodes = newNodes
		sort.Strings(paths)
		index := 0
		fmt.Println("Finding next root from", paths)
		foundMatchingRoot := false
		for i := 0; i < len(paths); i++ {
			root = paths[i]
			index = i
			fmt.Println("Checking if", step1, "contains", root)
			if strings.Contains(step1, root) {
				fmt.Println("root", root, "already added to", step1)
			} else {
				blocked := false
				for _, aNode := range nodes {
					if root == aNode.destination {
						fmt.Println(root, "blocked by", aNode)
						blocked = true
						break
					}
				}
				if !blocked {
					foundMatchingRoot = true
					break
				}
			}
		}
		if !foundMatchingRoot {
			fmt.Println("BLOCKED, no matching root")
			os.Exit(1)
		}
		fmt.Println("new root is", root)
		step1 = step1 + root
		fmt.Println("solution so far is", step1)
		paths = append(paths[:index], paths[index+1:]...)
		fmt.Println("Next steps are ", paths)
	}

	fmt.Println(step1)
}

func findRoots(origins, destinations []string) []string {
	roots := []string{}
	for _, origin := range origins {
		found := false
		for _, destination := range destinations {
			if origin == destination {
				found = true
				break
			}
		}
		if !found {
			rootAlreadyExist := false
			for _, root := range roots {
				if root == origin {
					rootAlreadyExist = true
				}
			}
			if !rootAlreadyExist {
				roots = append(roots, origin)
			}
		}
	}
	return roots
}

func load() string {
	text, err := ioutil.ReadFile("./input07.txt")
	if err != nil {
		fmt.Print(err)
	}
	return string(text)
}
