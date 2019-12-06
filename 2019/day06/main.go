package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type planet struct {
	name     string
	orbiters []planet
}

func main() {
	//rawListOfOrbits := load("orbit-test.txt")
	rawListOfOrbits := load("orbit.txt")
	orbits := [][]string{}
	for _, orbitAsString := range strings.Split(rawListOfOrbits, "\n") {
		orbit := strings.Split(string(orbitAsString), ")")
		orbits = append(orbits, orbit)

	}
	fmt.Println(orbits)
	orbitMap := make(map[string][]string)
	for i := 0; i < len(orbits); i++ {
		if orbitMap[orbits[i][0]] == nil {
			orbitMap[orbits[i][0]] = []string{orbits[i][1]}
		} else {
			orbitMap[orbits[i][0]] = append(orbitMap[orbits[i][0]], orbits[i][1])
		}
	}
	fmt.Println(orbitMap)
	sum := 0
	root := findOrbiters("COM", orbitMap, 1, &sum)
	fmt.Println(root, sum)
}

func findOrbiters(rootName string, orbitMap map[string][]string, depth int, sum *int) planet {
	rootPlanet := planet{rootName, []planet{}}
	for _, orbiter := range orbitMap[rootName] {
		*sum += depth
		rootPlanet.orbiters = append(rootPlanet.orbiters, findOrbiters(orbiter, orbitMap, depth+1, sum))
	}
	return rootPlanet
}

func load(filename string) string {
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	return string(text)
}
