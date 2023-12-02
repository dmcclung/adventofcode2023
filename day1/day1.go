package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func main() {
	file, _ := os.Open("day1/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	reBackwards := regexp.MustCompile(`\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin`)

	var sum int
	for scanner.Scan() {
		line := scanner.Text()		

		numbers := re.FindAllString(line, -1)

		firstNum := convert(numbers[0])

		numbers = reBackwards.FindAllString(reverseString(line), -1)

		lastNum := convert(reverseString(numbers[0]))

		sum += ((firstNum * 10) + lastNum)
	}

	fmt.Println(sum)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func convert(s string) int {
	if unicode.IsDigit(rune(s[0])) {
		digit, _ := strconv.Atoi(s)
		return digit
	}

	switch s {	
	case "one":
		return 1
	case "oneight":
		return 8
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	}

	return -1
}
