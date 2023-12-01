package main

import (
	"strconv"
	"strings"
	"unicode"
)

var testInpPartOne = `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`
var testInpPartTwo = `
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

func partOne() {
	numbers := make([]int, 0)

	for _, line := range strings.Split(testInpPartOne, "\n") {
		var digitOne, digitTwo string
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				digitOne = string(line[i])
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				digitTwo = string(line[i])
				break
			}
		}
		twoDigits := string(digitOne) + string(digitTwo)
		number, _ := strconv.Atoi(twoDigits)

		numbers = append(numbers, number)

	}
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	print(sum, "\n")
}

func startingDigit(line string) int {
	digitStrings := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	digitStringsToDigit := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	for i := 0; i < len(digitStrings); i++ {
		if strings.HasPrefix(line, digitStrings[i]) {
			return digitStringsToDigit[digitStrings[i]]
		}
	}
	return -1

}

func partTwo() {

	numbers := make([]int, 0)

	for _, line := range strings.Split(testInpPartTwo, "\n") {

		var digitOne, digitTwo string
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				digitOne = string(line[i])
				break
			} else {
				potentialDigit := startingDigit(line[i:])
				if potentialDigit == -1 {
					continue
				} else {
					digitOne = strconv.Itoa(potentialDigit)
					break
				}
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				digitTwo = string(line[i])
				break
			} else {
				potentialDigit := startingDigit(line[i:])
				if potentialDigit == -1 {
					continue
				} else {
					digitTwo = strconv.Itoa(potentialDigit)
					break
				}
			}
		}
		twoDigits := string(digitOne) + string(digitTwo)
		number, _ := strconv.Atoi(twoDigits)
		numbers = append(numbers, number)
	}

	sum := 0
	for _, number := range numbers {
		sum += number
	}

	print(sum, "\n")

}

func main() {
	partOne()
	partTwo()
}
