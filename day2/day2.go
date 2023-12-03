package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// The Elf would first like to know which games would have been
	// possible if the bag contained only
	// 12 red cubes, 13 green cubes, and 14 blue cubes?
	file, _ := os.Open("day2/input.txt")
	defer file.Close()

	re := regexp.MustCompile(`^Game (\d+):`)

	scanner := bufio.NewScanner(file)
	var sum, powerSum int
	for scanner.Scan() {
		game := scanner.Text()
		matches := re.FindStringSubmatch(game)

		gameNumber, _ := strconv.Atoi(matches[1])

		actionStr := re.ReplaceAllString(game, "")
		actionParts := strings.Split(actionStr, ";")

		power := parseGameMins(actionParts)
		powerSum += power

		valid := parseGameValid(actionParts)
		if valid {
			sum += gameNumber
		}
	}

	fmt.Printf("Power %d, Valid sum %d\n", powerSum, sum)
}

func parseGameMins(actionParts []string) int {
	var minBlue, minRed, minGreen int

	for _, p := range actionParts {
		parts := strings.Split(strings.TrimSpace(p), ",")

		for _, part := range parts {
			var number int
			var color string
			fmt.Sscanf(part, "%d %s", &number, &color)

			if color == "red" && number > minRed {
				minRed = number
			} else if color == "green" && number > minGreen {
				minGreen = number
			} else if color == "blue" && number > minBlue {
				minBlue = number
			}
		}
	}

	return minBlue * minRed * minGreen
}

func parseGameValid(actionParts []string) bool {
	for _, p := range actionParts {
		parts := strings.Split(strings.TrimSpace(p), ",")

		for _, part := range parts {
			var number int
			var color string
			fmt.Sscanf(part, "%d %s", &number, &color)

			if color == "red" && number > 12 {
				return false
			} else if color == "green" && number > 13 {
				return false
			} else if color == "blue" && number > 14 {
				return false
			}
		}
	}
	return true
}
