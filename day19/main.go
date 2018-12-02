package main

import (
    "strings"
    "io/ioutil"
    "fmt"
)

type coord struct {
    Y int
    X int
}

type direction struct {
    Y int
    X int
}

func walk(in string) (string, int) {
    maze := strings.Split(strings.Trim(in, "\n"), "\n")

    pos := coord{Y: 0, X: strings.IndexRune(maze[0], '|')}
    dir := direction{Y: 1, X: 0}
    out := make([]byte, 0)
    count := 0
    for {
        var prev coord
        prev, pos = pos, coord{pos.Y + dir.Y, pos.X + dir.X}
        count++
        if pos.Y < 0 || pos.Y >= len(maze) || pos.X < 0 || pos.X >= len(maze[0]) {
            break
        }

        char := maze[pos.Y][pos.X]
        if char >= 'A' && char <= 'Z' {
            out = append(out, char)
        }

        if char == ' ' {
            break
        }

        if char == '+' {
            // find another way out
            for _, nextDir := range []direction{direction{-1, 0}, direction{0, -1}, direction{0, 1}, direction{1, 0}} {
                next := coord{pos.Y + nextDir.Y, pos.X + nextDir.X}
                if next == prev {
                    continue
                }

                if next.Y < 0 || next.Y >= len(maze) || next.X < 0 || next.X >= len(maze[0]) {
                    continue
                }

                nextChar := maze[next.Y][next.X]
                if nextChar == ' ' {
                    continue
                }
                dir = nextDir
            }
        }
    }
    return string(out), count
}

func puzzle1(in string) string {
    letters, _ := walk(in)
    return letters
}

func puzzle2(in string) int {
    _, count := walk(in)
    return count
}

func main() {
  bytes, err := ioutil.ReadFile("./input.txt")
  if err != nil {
    fmt.Print(err)
  }
  fmt.Println(puzzle1((string(bytes))))
  fmt.Println(puzzle2((string(bytes))))
}
