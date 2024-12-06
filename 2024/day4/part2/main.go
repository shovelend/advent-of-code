package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	matrix := parseFile("input.txt")
	fmt.Println(totalCount(matrix))
}

func parseFile(fileName string) [][]byte {
	content, err := os.ReadFile(fileName)
	if err != nil {
	}
	lines := strings.Split(string(content), "\n")

	matrix := make([][]byte, len(lines))
	for i := range lines {
		matrix[i] = make([]byte, len(lines[i]))
		for j, char := range lines[i] {
			matrix[i][j] = byte(char)
		}
	}
	return matrix
}

func totalCount(board [][]byte) int {

	rowLength := len(board)
	colLength := len(board[0])

	total := 0

	for row := 1; row < rowLength-1; row++ {
		for col := 1; col < colLength-1; col++ {
			if board[row][col] == 'A' {
				leftdiagmatch := false
				rightdiagmatch := false
				if board[row-1][col-1] == 'S' && board[row+1][col+1] == 'M' {
					leftdiagmatch = true
				}
				if board[row-1][col-1] == 'M' && board[row+1][col+1] == 'S' {
					leftdiagmatch = true
				}

				if board[row-1][col+1] == 'S' && board[row+1][col-1] == 'M' {
					rightdiagmatch = true
				}
				if board[row-1][col+1] == 'M' && board[row+1][col-1] == 'S' {
					rightdiagmatch = true
				}

				if rightdiagmatch && leftdiagmatch {
					total += 1
				}
			}
		}
	}

	return total
}

// 1433 too low
// 1933 not right
// 2869 too high
// 3622 too high
// 3841 not right
