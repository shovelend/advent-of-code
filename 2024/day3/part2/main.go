package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	r, _ := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)|do\\(\\)|don't\\(\\)")
	b, err := os.ReadFile("input.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(b) // print the content as 'bytes'

	matches := r.FindAllString(string(b), -1)

	res := 0
	flag := true
	for _, match := range matches {
		if match == "do()" {
			flag = true
		} else if match == "don't()" {
			flag = false
		} else {
			if flag {
				res += calculateSubResult(match)
			}
		}
	}
	fmt.Println(res)
}

func calculateSubResult(s string) int {
	re := regexp.MustCompile("[0-9]+")
	subTotal := 0

	numArray := re.FindAllString(s, -1)
	firstEl, _ := strconv.Atoi(numArray[0])
	secondEl, _ := strconv.Atoi(numArray[1])
	fmt.Println(firstEl, secondEl)
	subTotal += firstEl * secondEl

	return subTotal
}

// 81458033 TOO HIGH!!!
// 70812129 TOO LOW
