package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	diskMap := parseInput()
	origArray := []int{}
	for i, elem := range diskMap {
		element, _ := strconv.Atoi(string(elem))
		switch i % 2 {
		case 0:
			for j := 0; j < element; j++ {
				origArray = append(origArray, i/2)
			}
		case 1:
			for j := 0; j < element; j++ {
				origArray = append(origArray, -1)
			}
		}
	}
	for j := len(origArray) - 1; j >= 0; j-- {
		earliestIndex := slices.Index(origArray, -1)
		if origArray[j] != -1 {
			origArray[earliestIndex] = origArray[j]
			origArray[j] = -1
		}
	}
	origArray = append(origArray[:0], origArray[1:]...)
	origArray = append(origArray, -1)
	total := 0
	for i, num := range origArray {
		if num != -1 {
			total += (i * num)
		}
	}

	fmt.Println(total)
}

func parseInput() string {
	byteContent, err := os.ReadFile("../input.txt")
	if err != nil {
	}
	content := string(byteContent)
	return content
}
