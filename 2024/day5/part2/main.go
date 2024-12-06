package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	total := 0
	rules := parsePageOrderingRules()
	lists := parsePagesToProduce()

	for _, line := range lists {
		line_is_safe := false
		if checkLine(line, rules) {
			continue
		}

		for !line_is_safe {
			for _, rule := range rules {
				if slices.Contains(line, rule[0]) && slices.Contains(line, rule[1]) {
					firstPos := slices.Index(line, rule[0])
					secondPos := slices.Index(line, rule[1])
					if firstPos > secondPos {
						tmp := line[firstPos]
						line[firstPos] = line[secondPos]
						line[secondPos] = tmp
					}
				}
			}

			if checkLine(line, rules) {
				line_is_safe = true
				total += line[len(line)/2]
			}
		}
	}

	fmt.Println(total)
}

func checkLine(page []int, rules [][2]int) bool {
	line_is_safe := true
	for _, rule := range rules {
		if slices.Contains(page, rule[0]) && slices.Contains(page, rule[1]) {
			if slices.Index(page, rule[0]) > slices.Index(page, rule[1]) {
				line_is_safe = false
				break
			}
		}
	}
	return line_is_safe
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
