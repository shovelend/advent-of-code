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
		if isSafe(line) {
			safeRecordCount += 1
		}

	}
	fmt.Println(safeRecordCount)
	readFile.Close()
}

func isSafe(text string) bool {
	arr := lineToArray(text)
	return allCorrectly(arr, "inc") || allCorrectly(arr, "dec")
}

func allCorrectly(items []int, order string) bool {
	for i := 1; i < len(items); i++ {
		isOrdered := true
		if order == "inc" {
			isOrdered = (items[i-1] < items[i])
		} else if order == "dec" {
			isOrdered = (items[i-1] > items[i])
		}
		if !isOrdered || !isWithinDifference(items[i-1], items[i]) {
			return false
		}

	}
	return true
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
