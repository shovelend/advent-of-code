package main

//4655 TOO HIGH
import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Pos struct {
	x int
	y int
}

func main() {
	origTable := parseInput()
	traversedPath := getTraversedPath()
	guardChars := []rune{'^', '>', 'v', '<'}
	infinite := 0

	for _, pos := range traversedPath {
		guardOnMap := true

		beenThereDoneThat := make([][][4]int, len(origTable))
		for i := range origTable {
			beenThereDoneThat[i] = make([][4]int, len(origTable[i]))
			for j := range origTable[i] {
				beenThereDoneThat[i][j][0] = 0
				beenThereDoneThat[i][j][1] = 0
				beenThereDoneThat[i][j][2] = 0
				beenThereDoneThat[i][j][3] = 0
			}
		}
		table := parseInput()
		if !slices.Contains(guardChars, origTable[pos.x][pos.y]) && origTable[pos.x][pos.y] != '#' {
			table[pos.x][pos.y] = '#'
		} else {
			continue
		}
		for guardOnMap {
			for i, rows := range table {
				for j := range rows {
					guardFace := slices.Index(guardChars, table[i][j])
					if guardFace != -1 {
						if beenThereDoneThat[i][j][guardFace] > 0 {
							infinite += 1
							guardOnMap = false
							break
						}
						switch guardFace {
						case 0:
							if i == 0 {
								guardOnMap = false
								break
							}
							if table[i-1][j] != '#' {
								table[i-1][j] = guardChars[0]
								beenThereDoneThat[i][j][guardFace] = 1
							} else {
								table[i][j] = guardChars[1]
							}
						case 1:
							if j == len(table[i])-1 {
								guardOnMap = false
								break
							}
							if table[i][j+1] != '#' {
								table[i][j+1] = guardChars[1]
								beenThereDoneThat[i][j][guardFace] = 1
							} else {
								table[i][j] = guardChars[2]
							}
						case 2:
							if i == len(table)-1 {
								guardOnMap = false
								break
							}
							if table[i+1][j] != '#' {
								table[i+1][j] = guardChars[2]
								beenThereDoneThat[i][j][guardFace] = 1
							} else {
								table[i][j] = guardChars[3]
							}
						case 3:
							if j == 0 {
								guardOnMap = false
								break
							}
							if table[i][j-1] != '#' {
								table[i][j-1] = guardChars[3]
								beenThereDoneThat[i][j][guardFace] = 1
							} else {
								table[i][j] = guardChars[0]
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(infinite)
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

func getTraversedPath() []Pos {
	table := parseInput()
	guardOnMap := true
	guardChars := []rune{'^', '>', 'v', '<'}

	for guardOnMap {
		for i, row := range table {
			for j := range row {
				guardFace := slices.Index(guardChars, table[i][j])
				if guardFace != -1 {
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

	traversedPath := []Pos{}
	for i, row := range table {
		for j := range row {

			if table[i][j] == 'X' {
				traversedPath = append(traversedPath, Pos{x: i, y: j})
			}
		}
	}
	return traversedPath
}
