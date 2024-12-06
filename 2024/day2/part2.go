package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	safeRecordCount := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		arr := lineToArray(line)
		if isSafe(arr, 0) {
			safeRecordCount += 1
		} else {
			// x, arr := arr[0], arr[1:]
			if isSafe(arr[1:], 1) {
				safeRecordCount += 1
			}
		}

	}
	fmt.Println(safeRecordCount)
	readFile.Close()
}

func isSafe(arr []int, depth int) bool {
	result := AllCorrectly(arr, "inc", depth) || AllCorrectly(arr, "dec", depth)
	return result
}

func AllCorrectly(items []int, order string, depth int) bool {
	for i := 1; i < len(items); i++ {
		inOrder := true
		if order == "inc" {
			inOrder = (items[i-1] < items[i])
		} else if order == "dec" {
			inOrder = (items[i-1] > items[i])
		}

		if !inOrder || !isWithinDifference(items[i-1], items[i]) {
			newItems := remove(items, i)

			if depth == 0 && AllCorrectly(newItems, order, 1) {
				return true
			} else {
				return false
			}
		}

	}
	return true
}

func remove(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func isWithinDifference(a, b int) bool {
	return a != b && absDiffInt(a, b) < 4
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func lineToArray(A string) []int {
	strs := strings.Split(A, " ")
	arr := make([]int, len(strs))
	for i := range arr {
		arr[i], _ = strconv.Atoi(strs[i])
	}

	return arr
}
