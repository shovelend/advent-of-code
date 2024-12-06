package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	table := parseInput()
	guardOnMap := true
	posVisited := 0
	guardChars := []rune{'^', '>', 'v', '<'}

	for guardOnMap {
		for i, row := range table {
			for j := range row {
				guardFace := slices.Index(guardChars, table[i][j])
				if guardFace != -1 {
					fmt.Println(i, j)
					switch guardFace {
					case 0:
						if i == 0 {
							table[i][j] = 'X'
							guardOnMap = false
							continue
						}
						if table[i-1][j] != '#' {
							table[i-1][j] = guardChars[0]
							table[i][j] = 'X'
						} else {
							table[i][j] = guardChars[1]
						}
						continue
					case 1:
						if j == len(table[i])-1 {
							table[i][j] = 'X'
							guardOnMap = false
							continue
						}
						if table[i][j+1] != '#' {
							table[i][j+1] = guardChars[1]
							table[i][j] = 'X'
						} else {
							table[i][j] = guardChars[2]
						}
						continue
					case 2:
						if i == len(table)-1 {
							table[i][j] = 'X'
							guardOnMap = false
							continue
						}
						if table[i+1][j] != '#' {
							table[i+1][j] = guardChars[2]
							table[i][j] = 'X'
						} else {
							table[i][j] = guardChars[3]
						}
						continue
					case 3:
						if j == 0 {
							table[i][j] = 'X'
							guardOnMap = false
							continue
						}
						if table[i][j-1] != '#' {
							table[i][j-1] = guardChars[3]
							table[i][j] = 'X'
						} else {
							table[i][j] = guardChars[0]
						}
						continue
					}
				}
			}
		}
	}

	for i, row := range table {
		for j := range row {

			if table[i][j] == 'X' {
				posVisited += 1
			}
		}
	}
	fmt.Println(posVisited)
}

func parseInput() [][]rune {
	content, err := os.ReadFile("../input.txt")
	if err != nil {
	}
	lines := strings.Split(string(content), "\n")

	table := make([][]rune, len(lines))
	for i, text := range lines {
		table[i] = make([]rune, len(text))
		for j, char := range text {
			table[i][j] = char
		}
	}
	return table
}
