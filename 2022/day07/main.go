package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	step1(load("small.txt"))
	step1(load("input.txt"))
	step2(load("small.txt"))
	step2(load("input.txt"))
}

type fs struct {
	name     string
	size     int
	children []*fs
	parent   *fs
}

func step1(input string) {
	root := buildTree(input)
	// printTree(root, 0)
	fmt.Println("Step1 --> ", sumBigDir(root))
}

func sumBigDir(root *fs) int {
	bigDirSum := 0
	if root.children != nil {
		if root.size < 100000 && root.name != "" {
			bigDirSum += root.size
		} else {
		}
		for _, child := range root.children {
			bigDirSum += sumBigDir(child)
		}
	}
	return bigDirSum
}

func printTree(root *fs, level int) {
	for i := 0; i < level; i++ {
		fmt.Print(" ")
	}

	if root.children != nil {
		fmt.Println("-", root.name+"/ (dir, size=", root.size, ")")
		for _, child := range root.children {
			printTree(child, level+1)
		}
	} else {
		fmt.Println(">", root.name, "(file, size=", root.size, ")")
	}
}

func findDirToDelete(root *fs, freeSpaceToReclaim int) []int {
	dirSize := []int{}
	if root.children != nil {
		if root.size > freeSpaceToReclaim {
			dirSize = append(dirSize, root.size)
		}
		for _, child := range root.children {
			childSizes := findDirToDelete(child, freeSpaceToReclaim)
			for _, childSize := range childSizes {
				dirSize = append(dirSize, childSize)
			}
		}
	}
	return dirSize
}

func step2(input string) {
	root := buildTree(input)
	unusedSpace := 70000000 - root.size
	freeSpaceToReclaim := 30000000 - unusedSpace

	dirSizeToDelete := findDirToDelete(root, freeSpaceToReclaim)[:]
	sort.Ints(dirSizeToDelete)

	fmt.Println("Step2 --> ", dirSizeToDelete[0])
}

// We compute dir size as we go, add to hack the inputs and add a "$ cd .." for the
// root dir to be accurate. it's probably not useful, but for the sake of correctness..
func buildTree(input string) *fs {
	root := &fs{"", 0, []*fs{}, nil}
	currentDir := root
	for _, line := range strings.Split(input, "\n") {
		tokens := strings.Split(line, " ")
		if tokens[0] == "$" {
			if tokens[1] == "cd" {
				if tokens[2] == ".." {
					currentDir.parent.size += currentDir.size
					currentDir = currentDir.parent
				} else {
					for _, node := range currentDir.children {
						if node.name == tokens[2] {
							currentDir = node
							break
						}
					}
				}
			} else {
				// ls -- ignore
			}
		} else {
			var file *fs
			if tokens[0] == "dir" {
				file = &fs{tokens[1], 0, []*fs{}, currentDir}
			} else {
				file = &fs{tokens[1], atoi(tokens[0]), nil, nil}
				currentDir.size += file.size
			}
			currentDir.children = append(currentDir.children, file)
		}
	}
	return root
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
