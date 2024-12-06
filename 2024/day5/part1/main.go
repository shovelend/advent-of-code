package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	pageOrderingRules := parsePageOrderingRules()
	pagesToProduce := parsePagesToProduce()
	fmt.Println(pageOrderingRules)
	fmt.Println(pagesToProduce)

	correctlyOrderedUpdateLines := make([][]int, len(pagesToProduce))

	for _, pages := range pagesToProduce {
		goodLine := true
		for _, rule := range pageOrderingRules {
			if slices.Contains(pages, rule[0]) && slices.Contains(pages, rule[1]) && slices.Index(pages, rule[0]) > slices.Index(pages, rule[1]) {
				goodLine = false
			}
		}
		if goodLine {
			correctlyOrderedUpdateLines = append(correctlyOrderedUpdateLines, pages)
		}
	}
	fmt.Println(correctlyOrderedUpdateLines)
	total := 0
	for _, lines := range correctlyOrderedUpdateLines {
		if len(lines) > 0 {
			total += lines[len(lines)/2]
		}
	}
	fmt.Println(total)
}

func parsePageOrderingRules() [][2]int {
	content, err := os.ReadFile("../input1.txt")
	if err != nil {
	}
	lines := strings.Split(string(content), "\n")

	pageOrderingRules := make([][2]int, len(lines))
	for i, text := range lines {
		line := strings.Split(text, "|")
		pageOrderingRules[i][0], _ = strconv.Atoi(line[0])
		pageOrderingRules[i][1], _ = strconv.Atoi(line[1])

	}
	return pageOrderingRules
}

func parsePagesToProduce() [][]int {
	content, err := os.ReadFile("../input2.txt")
	if err != nil {
	}
	lines := strings.Split(string(content), "\n")

	pagesToProduce := make([][]int, len(lines))
	for i, text := range lines {
		strings := strings.Split(text, ",")
		pagesToProduce[i] = make([]int, len(strings))
		for j, stringVal := range strings {
			pagesToProduce[i][j], _ = strconv.Atoi(stringVal)
		}

	}
	return pagesToProduce
}
