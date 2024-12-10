package main

import (
	"fmt"
	"os"
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

	total := 0
	origArray = getPart(origArray)
	for i, num := range origArray {
		if num != -1 {
			total += (i * num)
		}
	}
	fmt.Println(total)
}

func getPart(origArray []int) []int {
	for j := len(origArray) - 1; j > 0; {
		idCounter := origArray[j]
		numOfThingsToMove := 0
		if origArray[j] == -1 {
			j -= 1
			continue
		}
		currj := j

		for currj != -1 && origArray[currj] == idCounter {
			numOfThingsToMove += 1
			currj -= 1
		}
		origArray = findFirstFreeBlockSpaceId(origArray, numOfThingsToMove, currj+1)
		j = currj
	}

	return origArray
}

func findFirstFreeBlockSpaceId(origArray []int, numOfThingsToMove int, idx int) []int {
	for i := 0; i < idx; i++ {
		if origArray[i] == -1 {
			fits := true
			for j := 0; j < numOfThingsToMove; j++ {
				if origArray[i+j] != -1 {
					fits = false
				}
			}
			if fits {
				elToInsert := origArray[idx]
				for j := 0; j < numOfThingsToMove; j++ {
					origArray[i+j] = elToInsert
					origArray[idx+j] = -1
				}
			}
		}
	}
	return origArray
}

func parseInput() string {
	byteContent, err := os.ReadFile("../input.txt")
	if err != nil {
	}
	content := string(byteContent)
	return content
}
