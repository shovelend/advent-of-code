package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Combine struct {
	target  int64
	numbers []int64
}

func main() {
	table := parseInput()
	total := int64(0)
	for _, combine := range table {
		if equals(combine, 0) {
			total += (combine.target)
		}
	}
	fmt.Println(total)
}

func equals(c Combine, total int64) bool {
	if total > c.target {
		return false
	}
	if len(c.numbers) == 0 {
		return total == c.target
	}
	currHead := c.numbers[0]
	c = Combine{target: c.target, numbers: c.numbers[1:len(c.numbers)]}
	appended := strings.Join([]string{strconv.FormatInt(total, 10), strconv.FormatInt(currHead, 10)}, "")
	intAppended, _ := strconv.Atoi(appended)
	return equals(c, total+currHead) || equals(c, total*currHead) || equals(c, int64(intAppended))
}

func parseInput() []Combine {
	content, err := os.ReadFile("../input.txt")
	if err != nil {
	}
	lines := strings.Split(string(content), "\n")

	table := make([]Combine, len(lines))
	for i, line := range lines {
		splintered := strings.Split(line, ":")
		nums := parseNums(splintered[1])
		res, _ := strconv.Atoi(splintered[0])
		table[i] = Combine{target: int64(res), numbers: nums}
	}
	return table
}

func parseNums(nums string) []int64 {

	strs := strings.Split(nums, " ")
	ary := make([]int64, len(strs))
	for i := range ary {
		res, _ := strconv.Atoi(strs[i])
		ary[i] = int64(res)
	}

	return ary
}
