package main

import (
	"strconv"
	"strings"
	"unicode"
)

var testInput = `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

func isPart(element string) bool {
	if element == "." {
		return false
	}
	if unicode.IsDigit([]rune(element)[0]) {
		return false
	}
	return true
}

func checkAdjacent(xStart, xEnd, y int, input [][]string) bool {
	if xStart > 0 {
		left := input[y][xStart-1]
		if isPart(left) {
			return true
		}
	}
	if xEnd < len(input[y])-1 {
		right := input[y][xEnd+1]
		if isPart(right) {
			return true
		}
	}
	// top row
	for i := xStart - 1; i <= xEnd+1; i++ {
		if i < 0 || i >= len(input[y]) {
			continue
		}
		if y > 0 {
			top := input[y-1][i]
			if isPart(top) {
				return true
			}
		}
		if y < len(input)-1 {
			bottom := input[y+1][i]
			if isPart(bottom) {
				return true
			}
		}
	}
	return false
}

func partOne() {
	var parsedInput = make([][]string, 0)
	sum := 0
	for _, line := range strings.Split(testInput, "\n") {
		if len(line) == 0 {
			continue
		}
		parsedInput = append(parsedInput, strings.Split(strings.TrimSpace(line), ""))
	}
	for lineI, line := range parsedInput {
		start := 0
		end := 0
		hasStart := false
		for i, char := range line {
			if unicode.IsDigit([]rune(char)[0]) {
				if !hasStart {
					start = i
					hasStart = true
				}
				end = i
			}
			// check if current element is last element of the line or if next element is not a digit
			if hasStart && (i == len(line)-1 || !unicode.IsDigit([]rune(line[i+1])[0])) {
				number, _ := strconv.Atoi(strings.Join(line[start:end+1], ""))
				if checkAdjacent(start, end, lineI, parsedInput) {
					sum += number
				}
				start = 0
				end = 0
				hasStart = false
			}
		}
	}
	println("sum", sum)
}

// ---------- PART TWO ------------

func findNumbersInRow(row []string, middlePoint int) []string {

	if !unicode.IsDigit([]rune(row[middlePoint])[0]) {
		// if middle point is not a number then check left and right, there might be 2 number
		leftCursor := middlePoint - 1
		rightCursor := middlePoint + 1
		leftNumber := ""
		rightNumber := ""
		for leftCursor >= 0 {

			if unicode.IsDigit([]rune(row[leftCursor])[0]) {
				leftNumber = row[leftCursor] + leftNumber
			} else {
				break
			}
			leftCursor--
		}
		for rightCursor < len(row) {

			if unicode.IsDigit([]rune(row[rightCursor])[0]) {
				rightNumber = rightNumber + row[rightCursor]
			} else {
				break
			}
			rightCursor++
		}

		result := make([]string, 0)
		if leftNumber != "" {
			result = append(result, leftNumber)
		}
		if rightNumber != "" {
			result = append(result, rightNumber)
		}
		return result

	} else {
		number := row[middlePoint]
		leftCursor := middlePoint - 1
		rightCursor := middlePoint + 1
		for leftCursor >= 0 {

			if unicode.IsDigit([]rune(row[leftCursor])[0]) {
				number = row[leftCursor] + number
			} else {
				break
			}
			leftCursor--
		}
		for rightCursor < len(row) {

			if unicode.IsDigit([]rune(row[rightCursor])[0]) {
				number = number + row[rightCursor]
			} else {
				break
			}
			rightCursor++
		}
		return []string{number}
		// if middle point is a number then check left and right, there might be 1 number
	}

}

func findAdjacentGearNumbers(input [][]string, x, y int) []string {
	combined := make([]string, 0)
	// check left
	if x > 0 && unicode.IsDigit([]rune(input[y][x-1])[0]) {
		left := findNumbersInRow(input[y], x-1)
		combined = append(combined, left...)
	}
	if x < len(input[y])-1 && unicode.IsDigit([]rune(input[y][x+1])[0]) {
		right := findNumbersInRow(input[y], x+1)
		combined = append(combined, right...)
	}
	// check top
	if y > 0 {
		top := findNumbersInRow(input[y-1], x)
		combined = append(combined, top...)
	}

	// check bottom
	if y < len(input)-1 {
		bottom := findNumbersInRow(input[y+1], x)
		combined = append(combined, bottom...)
	}
	if len(combined) == 2 {
		return combined
	}
	return []string{}
}

func partTwo() {
	var parsedInput = make([][]string, 0)
	sum := 0
	for _, line := range strings.Split(testInput, "\n") {
		if len(line) == 0 {
			continue
		}
		parsedInput = append(parsedInput, strings.Split(strings.TrimSpace(line), ""))
	}
	for lineI, line := range parsedInput {
		for i, char := range line {
			if char == "*" {
				ratio := 1
				numbers := findAdjacentGearNumbers(parsedInput, i, lineI)
				for _, number := range numbers {
					numberInt, _ := strconv.Atoi(number)
					ratio *= numberInt
				}
				if ratio != 1 {
					sum += ratio
				}
			}
		}
	}
	println("sum", sum)
}

func main() {
	partOne()
	partTwo()
}
