package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	matrix := parseFile("input.txt")
	fmt.Println(totalCount(matrix, "SAMX"))
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

func totalCount(board [][]byte, word string) int {

	rowLength := len(board)
	colLength := len(board[0])

	total := 0

	var dfs func(row, col, idx int, direction string) bool

	dfs = func(row, col, idx int, direction string) bool {
		if idx == len(word) { // match all prefix
			return true
		}
		if row < 0 || row >= rowLength || col < 0 || col >= colLength {
			return false
		}
		// not match
		if board[row][col] != word[idx] {
			return false
		}
		switch direction {
		case "above":
			return dfs(row-1, col, idx+1, direction)
		case "below":
			return dfs(row+1, col, idx+1, direction)
		case "left":
			return dfs(row, col-1, idx+1, direction)
		case "right":
			return dfs(row, col+1, idx+1, direction)
		case "diagupleft":
			return dfs(row-1, col-1, idx+1, direction)
		case "diagupright":
			return dfs(row-1, col+1, idx+1, direction)
		case "diagdownright":
			return dfs(row+1, col+1, idx+1, direction)
		case "diagdownleft":
			return dfs(row+1, col-1, idx+1, direction)
		}
		return false
	}
	for row := 0; row < rowLength; row++ {
		for col := 0; col < colLength; col++ {
			if dfs(row, col, 0, "above") {
				total += 1
			}
			if dfs(row, col, 0, "below") {
				total += 1
			}
			if dfs(row, col, 0, "left") {
				total += 1
			}
			if dfs(row, col, 0, "right") {
				total += 1
			}
			if dfs(row, col, 0, "diagupleft") {
				total += 1
			}
			if dfs(row, col, 0, "diagupright") {
				total += 1
			}
			if dfs(row, col, 0, "diagdownright") {
				total += 1
			}
			if dfs(row, col, 0, "diagdownleft") {
				total += 1
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
