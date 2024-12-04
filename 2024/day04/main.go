package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	Row, Col             int
	RowOffset, ColOffset int
}

func findLetters(grid [][]rune, letter rune, start Coord) []Coord {
	var coords []Coord
	for rowOffset := -1; rowOffset <= 1; rowOffset++ {
		for colOffset := -1; colOffset <= 1; colOffset++ {
			if rowOffset == 0 && colOffset == 0 {
				continue
			}
			row, col := start.Row+rowOffset, start.Col+colOffset
			if row >= 0 && row < len(grid) && col >= 0 && col < len(grid[row]) && grid[row][col] == letter {
				coords = append(coords, Coord{row, col, rowOffset, colOffset})
			}
		}
	}
	return coords
}

func loadGrid(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var grid [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		grid = append(grid, row)
	}
	return grid
}

func step1(filename string) int {
	grid := loadGrid(filename)
	foundCount := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 'X' {
				xCoord := Coord{row, col, 0, 0}
				mCoords := findLetters(grid, 'M', xCoord)
				for _, mCoord := range mCoords {
					aRow, aCol := mCoord.Row+mCoord.RowOffset, mCoord.Col+mCoord.ColOffset
					if aRow < 0 || aRow >= len(grid) || aCol < 0 || aCol >= len(grid[aRow]) {
						continue
					}
					aCoord := Coord{aRow, aCol, mCoord.RowOffset, mCoord.ColOffset}
					if grid[aCoord.Row][aCoord.Col] == 'A' {
						sRow, sCol := aCoord.Row+aCoord.RowOffset, aCoord.Col+aCoord.ColOffset
						if sRow < 0 || sRow >= len(grid) || sCol < 0 || sCol >= len(grid[sRow]) {
							continue
						}
						if grid[sRow][sCol] == 'S' {
							foundCount++
						}
					}
				}
			}
		}
	}
	return foundCount
}

func isInBounds(grid [][]rune, coord Coord) bool {
	row, col := coord.Row, coord.Col
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[row])
}

func step2(filename string) int {
	grid := loadGrid(filename)
	foundCount := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 'M' {
				firstMCoord := Coord{row, col, 0, 0}
				for rowOffset := -2; rowOffset <= 2; rowOffset += 2 {
					for colOffset := -2; colOffset <= 2; colOffset += 2 {
						if rowOffset == 0 && colOffset == 0 {
							continue
						}
						row2, col2 := row+rowOffset, col+colOffset
						if row2 >= 0 && row2 < len(grid) && col2 >= 0 && col2 < len(grid[row2]) {
							if grid[row2][col2] == 'M' {
								// found two matching M
								secondMCoord := Coord{row2, col2, 0, 0}
								if !isInBounds(grid, secondMCoord) {
									continue
								}
								// if the M are aligned horizontally or vertically, then we can predict where to look for an A
								if firstMCoord.Row == secondMCoord.Row {
									for rowOffset := -1; rowOffset <= 1; rowOffset += 2 {
										aCoord := Coord{firstMCoord.Row + rowOffset, (firstMCoord.Col + secondMCoord.Col) / 2, rowOffset, 0}
										if isInBounds(grid, aCoord) && grid[aCoord.Row][aCoord.Col] == 'A' {
											firstSCoord := Coord{aCoord.Row + rowOffset, firstMCoord.Col, rowOffset, 0}
											secondSCoord := Coord{aCoord.Row + rowOffset, secondMCoord.Col, rowOffset, 0}
											if isInBounds(grid, firstSCoord) &&
												isInBounds(grid, secondSCoord) &&
												grid[firstSCoord.Row][firstSCoord.Col] == 'S' &&
												grid[secondSCoord.Row][secondSCoord.Col] == 'S' {
												foundCount++
											}
										}
									}
								} else if firstMCoord.Col == secondMCoord.Col {
									for colOffset := -1; colOffset <= 1; colOffset += 2 {
										aCoord := Coord{(firstMCoord.Row + secondMCoord.Row) / 2, firstMCoord.Col + colOffset, 0, colOffset}
										if isInBounds(grid, aCoord) && grid[aCoord.Row][aCoord.Col] == 'A' {
											firstSCoord := Coord{firstMCoord.Row, aCoord.Col + aCoord.ColOffset, rowOffset, 0}
											secondSCoord := Coord{secondMCoord.Row, aCoord.Col + aCoord.ColOffset, rowOffset, 0}
											if isInBounds(grid, firstSCoord) &&
												isInBounds(grid, secondSCoord) &&
												grid[firstSCoord.Row][firstSCoord.Col] == 'S' &&
												grid[secondSCoord.Row][secondSCoord.Col] == 'S' {
												foundCount++
											}
										}
									}
								}
							}
						}

					}
				}
			}
		}
	}
	return foundCount / 2
}

func main() {
	fmt.Printf("step1, example1.txt: Found %d S's\n", step1("example1.txt"))
	fmt.Printf("step1, step1.txt: Found %d S's\n", step1("step1.txt"))
	fmt.Printf("step2, example2.txt: Found %d X-MAS's\n", step2("example2.txt"))
	fmt.Printf("step2, example3.txt: Found %d X-MAS's\n", step2("example3.txt"))
	fmt.Printf("step2, example1.txt: Found %d X-MAS's\n", step2("example1.txt"))
	fmt.Printf("step2, step1.txt: Found %d X-MAS's\n", step2("step1.txt"))
}
