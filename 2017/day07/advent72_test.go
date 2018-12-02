package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestFirstMap(t *testing.T) {
	fmt.Println(advent72("tknk", sampleInput))
}

func TestSecondExercice(t *testing.T) {
	text, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(advent72("fbgguv", string(text)))
	fmt.Println(advent72("sfruur", string(text)))
	fmt.Println(advent72("zsasjr", string(text)))
	fmt.Println(advent72("jdxfsa", string(text)))

}

type nodeCache struct {
	name          string
	value         int
	childrenNames []string
}

type node struct {
	name     string
	value    int
	sum      int
	parent   *node
	children []node
}

func advent72(rootName string, input string) string {
	if input == "" {
		return ""
	}
	nodesCache := buildNodeCache(input)
	rootNode := buildNode(rootName, nodesCache, nil)
	for _, node := range rootNode.children {
		fmt.Println(node.name, node.sum, node.value)
	}
	// fmt.Println(rootNode.children)
	return ""
}

func buildNode(nodeName string, nodesCache map[string]nodeCache, parentNode *node) node {
	//fmt.Println("creating new node ", nodeName)
	nodeCache := nodesCache[nodeName]
	newNode := node{name: nodeCache.name, value: nodeCache.value, parent: parentNode, sum: nodeCache.value}
	for _, name := range nodeCache.childrenNames {
		//fmt.Println("Attaching ", name, " to ", nodeName)
		childNode := buildNode(name, nodesCache, &newNode)
		newNode.sum += childNode.sum
		newNode.children = append(newNode.children, childNode)
	}
	return newNode
}

func buildNodeCache(input string) map[string]nodeCache {
	nodeMap := map[string]nodeCache{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		line = strings.Replace(line, ",", "", -1)
		words := strings.Split(line, " ")
		name := words[0]
		value, err := strconv.Atoi(words[1][1 : len(words[1])-1])
		if err != nil {
			break
		}
		childrenNames := []string{}
		if len(words) > 3 {
			for _, childName := range words[3:] {
				childrenNames = append(childrenNames, childName)
			}
		}
		nodeMap[name] = nodeCache{name: name, value: value, childrenNames: childrenNames}
		// fmt.Println(nodeMap[name])
	}
	return nodeMap
}
