package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	table := parseInput()
	origTable := parseInput()
	infinite := 0
	guardChars := []rune{'^', '>', 'v', '<'}

	for io, rowo := range origTable {
		for jo := range rowo {
			guardOnMap := true
			posVisited := 0

			tmp := origTable[io][jo]
			if !slices.Contains(guardChars, origTable[io][jo]) {
				table[io][jo] = '#'
			}
			// TODO fix infinite loop
			for guardOnMap {
				fmt.Println(io, jo)
				fmt.Println(posVisited)
				if posVisited > 10000 {
					infinite += 1
					break
				}
				for i, row := range table {
					for j := range row {
						guardFace := slices.Index(guardChars, table[i][j])
						if guardFace != -1 {
							switch guardFace {
							case 0:
								if i == 0 {
									setVisitedInc(table, i, j, &posVisited)
									guardOnMap = false
									continue
								}
								if table[i-1][j] != '#' {
									table[i-1][j] = guardChars[0]
									setVisitedInc(table, i, j, &posVisited)
								} else {
									table[i][j] = guardChars[1]
								}
								continue
							case 1:
								if j == len(table[i])-1 {
									setVisitedInc(table, i, j, &posVisited)
									guardOnMap = false
									continue
								}
								if table[i][j+1] != '#' {
									table[i][j+1] = guardChars[1]
									setVisitedInc(table, i, j, &posVisited)
								} else {
									table[i][j] = guardChars[2]
								}
								continue
							case 2:
								if i == len(table)-1 {
									setVisitedInc(table, i, j, &posVisited)
									guardOnMap = false
									continue
								}
								if table[i+1][j] != '#' {
									table[i+1][j] = guardChars[2]
									setVisitedInc(table, i, j, &posVisited)
								} else {
									table[i][j] = guardChars[3]
								}
								continue
							case 3:
								if j == 0 {
									setVisitedInc(table, i, j, &posVisited)
									guardOnMap = false
									continue
								}
								if table[i][j-1] != '#' {
									table[i][j-1] = guardChars[3]
									setVisitedInc(table, i, j, &posVisited)
								} else {
									table[i][j] = guardChars[0]
								}
								continue
							}
						}
					}
				}
			}
			origTable[io][jo] = tmp
		}
	}

	fmt.Println(infinite)
}

func setVisitedInc(table [][]rune, i, j int, posVisited *int) {
	if table[i][j] != '#' {
		if table[i][j] != 'X' {
			*posVisited += 1
		}
		table[i][j] = 'X'
	}
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
