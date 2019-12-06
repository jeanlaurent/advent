package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type planet struct {
	name     string
	orbiters []planet
	depth    int
}

func main() {
	orbits := loadOrbits()

	orbitMap := make(map[string][]string)
	for i := 0; i < len(orbits); i++ {
		if orbitMap[orbits[i][0]] == nil {
			orbitMap[orbits[i][0]] = []string{orbits[i][1]}
		} else {
			orbitMap[orbits[i][0]] = append(orbitMap[orbits[i][0]], orbits[i][1])
		}
	}

	sum := 0
	root := findOrbiters("COM", orbitMap, 1, &sum)
	fmt.Println("step1 --->", sum)

	// find orbiter of YOU, keep the full chain
	pathToYou := make(map[string]int)
	findPathTo(root, "YOU", pathToYou)
	youKeys := mapKeys(pathToYou)

	// find orbiter of SAN, keep the full chain
	pathToSan := make(map[string]int)
	findPathTo(root, "SAN", pathToSan)
	sanKeys := mapKeys(pathToSan)

	// Find a common node between the two
	commonMax := 0
	for _, key := range youKeys {
		if pathToSan[key] != 0 {
			if pathToSan[key] > commonMax {
				commonMax = pathToSan[key]
			}
		}
	}

	// find max value of YOU
	youMax := 0
	for _, key := range youKeys {
		if pathToYou[key] > youMax {
			youMax = pathToYou[key]
		}
	}
	// find max value of SAN
	sanMax := 0
	for _, key := range sanKeys {
		if pathToSan[key] > sanMax {
			sanMax = pathToSan[key]
		}
	}

	// add distance from SAN minus distance of common node + YOU and distance of common node
	fmt.Println("step2 --->", sanMax-commonMax+youMax-commonMax)
}

func mapKeys(aMap map[string]int) []string {
	sanKeys := make([]string, 0, len(aMap))
	for key := range aMap {
		sanKeys = append(sanKeys, key)
	}
	return sanKeys
}

func findPathTo(root planet, target string, path map[string]int) bool {
	if root.name == target {
		return true
	}
	for _, orbiter := range root.orbiters {
		if findPathTo(orbiter, target, path) {
			if orbiter.name != target {
				path[orbiter.name] = orbiter.depth
			}
			return true
		}
	}
	return false
}

func findOrbiters(rootName string, orbitMap map[string][]string, depth int, sum *int) planet {
	rootPlanet := planet{rootName, []planet{}, depth}
	for _, orbiter := range orbitMap[rootName] {
		*sum += depth
		rootPlanet.orbiters = append(rootPlanet.orbiters, findOrbiters(orbiter, orbitMap, depth+1, sum))
	}
	return rootPlanet
}

func loadOrbits() [][]string {
	//rawListOfOrbits := load("orbit-test.txt")
	rawListOfOrbits := load("orbit.txt")
	orbits := [][]string{}
	for _, orbitAsString := range strings.Split(rawListOfOrbits, "\n") {
		orbit := strings.Split(string(orbitAsString), ")")
		orbits = append(orbits, orbit)

	}
	return orbits
}

func load(filename string) string {
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	return string(text)
}
