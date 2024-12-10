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
	fmt.Println(origArray)
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
		for j != -1 && origArray[j] == idCounter {
			numOfThingsToMove += 1
			j -= 1
		}
		fmt.Println(origArray[j+1])
		origArray = findFirstFreeBlockSpaceId(origArray, numOfThingsToMove, j+1)
	}

	return origArray
}

func findFirstFreeBlockSpaceId(origArray []int, numOfThingsToMove int, idx int) []int {
	for i := 0; i < len(origArray)-numOfThingsToMove; i++ {
		if origArray[i] == -1 {
			fits := true
			for j := 0; j < numOfThingsToMove; j++ {
				// fmt.Println(numOfThingsToMove, i, j, idx)
				if origArray[i+j] != -1 {
					fits = false
				}
			}
			if fits {
				for j := 0; j < numOfThingsToMove; j++ {
					// fmt.Println(origArray)
					origArray[i+j] = origArray[idx]
					origArray[idx+j] = -1
				}
			}
		}
	}
	return origArray
}

func parseInput() string {
	byteContent, err := os.ReadFile("test.txt")
	if err != nil {
	}
	content := string(byteContent)
	return content
}
