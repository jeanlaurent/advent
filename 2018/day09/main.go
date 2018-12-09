package main

import "fmt"

type game struct {
	turn            int
	currentPlayer   int
	nextMarbleValue int
	start           *marble
	circle          *marble
	currentMarble   *marble
	playerLength    int
	maxMarble       int
	scores          map[int]int
}

func (g *game) isFinished() bool {
	return g.turn == g.maxMarble
}

func (g *game) winnerScore() int {
	max := 0
	for _, score := range g.scores {
		if max < score {
			max = score
		}
	}
	return max
}

func (g *game) nextTurn() {
	g.currentPlayer = g.currentPlayer%g.playerLength + 1

	if g.nextMarbleValue%23 == 0 {
		g.scores[g.currentPlayer] += g.nextMarbleValue // not sure of next or previous player here
		toBeRemoved := g.currentMarble.previous.previous.previous.previous.previous.previous.previous
		g.scores[g.currentPlayer] += toBeRemoved.value
		source := toBeRemoved.previous
		destination := toBeRemoved.next
		source.next = destination
		destination.previous = source
		g.currentMarble = destination
	} else {
		destination := g.currentMarble.next.next
		source := g.currentMarble.next
		newMarble := marble{g.nextMarbleValue, source, destination}
		source.next = &newMarble
		destination.previous = &newMarble
		g.currentMarble = &newMarble
	}
	g.nextMarbleValue++
	g.turn++
}

func (g *game) print() {
	fmt.Print("[", g.currentPlayer, "] ")
	if g.start == g.currentMarble {
		fmt.Print("(", g.start.value, ")")
	} else {
		fmt.Print(" ", g.start.value, "  ")
	}
	next := g.start.next
	for next != g.start {
		if next == g.currentMarble {
			fmt.Print("(", next.value, ") ")
		} else {
			fmt.Print(" ", next.value, "  ")
		}
		next = next.next
	}
	fmt.Println()
}

func newGame(players, maxMarble int) game {
	zero := marble{0, nil, nil}
	zero.previous = &zero
	zero.next = &zero
	scores := map[int]int{}
	for index := 0; index < players; index++ {
		scores[index] = 0
	}
	return game{0, 0, 1, &zero, &zero, &zero, players, maxMarble, scores}
}

type marble struct {
	value    int
	previous *marble
	next     *marble
}

func runGame(players, marble int) int {
	game := newGame(players, marble)
	// game.print()
	for !game.isFinished() {
		game.nextTurn()
		//game.print()
	}
	return game.winnerScore()
}

func main() {
	fmt.Println(9, 25, "-->", runGame(9, 25))
	fmt.Println(10, 1618, "-->", runGame(10, 1618))
	fmt.Println(13, 7999, "-->", runGame(13, 7999))
	fmt.Println(17, 1104, "-->", runGame(17, 1104))
	fmt.Println(21, 6111, "-->", runGame(21, 6111))
	fmt.Println(30, 5807, "-->", runGame(30, 5807))
	fmt.Println("Step1", 430, 71588, "-->", runGame(430, 71588))
	fmt.Println("Step1", 430, 7158800, "-->", runGame(430, 7158800))
}
