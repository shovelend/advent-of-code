package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	r, _ := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)")
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	total := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		total += calculateSubResult(r.FindAllString(line, -1))
	}

	fmt.Println(total)
	readFile.Close()
}

func calculateSubResult(s []string) int {
	re := regexp.MustCompile("[0-9]+")
	subTotal := 0
	for _, el := range s {
		numArray := re.FindAllString(el, -1)
		firstEl, _ := strconv.Atoi(numArray[0])
		secondEl, _ := strconv.Atoi(numArray[1])
		fmt.Println(firstEl, secondEl)
		subTotal += firstEl * secondEl
	}
	return subTotal
}

// 81458033 TOO HIGH!!!
// 74838033
// 70812129 TOO LOW
