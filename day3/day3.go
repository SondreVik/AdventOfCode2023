package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Solve() {
	fmt.Println("Day 3")
	readFile, err := os.Open("day3/data.txt")
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	fileLines := []string{}

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	matrix := [][]rune{}
	for _, line := range fileLines {
		array := []rune{}
		for _, char := range line {
			array = append(array, char)
		}
		matrix = append(matrix, array)
	}
	fmt.Println(Part1(matrix))
	fmt.Println(part2(matrix))
}

type CharCoord struct {
	index int
	char  rune
}

func Part1(matrix [][]rune) (sum int) {
	fmt.Println("Part 1")
	for i, chars := range matrix {
		charCoords := []rune{}
		charText := ""
		minX := 0
		maxX := 0
		for k, char := range chars {
			if unicode.IsNumber(char) {
				if len(charCoords) == 0 {
					minX = k - 1
				}
				maxX = k + 1
				charCoords = append(charCoords, char)
				charText += string(char)
			}
			if len(charCoords) > 0 && (!unicode.IsNumber(char) || k+1 == len(chars)) {
				isValid := false
				maxX = min(maxX, len(chars)-1)
				minX = max(minX, 0)
				minY := max(i-1, 0)
				maxY := min(i+1, len(matrix)-1)
				for y := minY; y <= maxY; y++ {

					for x := minX; x <= maxX; x++ {
						if matrix[y][x] != '.' && !unicode.IsNumber(matrix[y][x]) {
							isValid = true
							x = maxX + 1
						}
					}
					if isValid {
						y = maxY + 1
					}
				}
				if isValid {
					number, err := strconv.Atoi(charText)
					check(err)
					sum += number
					isValid = false
				}
				charCoords = []rune{}
				charText = ""
			}
		}
	}
	return
}

type Coord struct {
	x int
	y int
}

func getNumber(chars []rune, index int) (number int) {
	word := ""
	start := index
	for i := index; i >= 0; i-- {
		if !unicode.IsNumber(chars[i]) {
			break
		}
		start = i
	}
	for i := start; i < len(chars); i++ {
		if !unicode.IsNumber(chars[i]) {
			break
		}
		word += string(chars[i])
	}

	result, err := strconv.Atoi(word)
	check(err)
	number = result
	return
}
func part2(matrix [][]rune) (sum int) {
	fmt.Println("Part 2")
	for i, chars := range matrix {
		for k, char := range chars {
			if char == '*' {
				numbers := []int{}
				for y := max(i-1, 0); y <= min(i+1, len(matrix)-1); y++ {
					lastWasNumber := false
					for x := max(k-1, 0); x <= min(k+1, len(chars)-1); x++ {
						if unicode.IsNumber(matrix[y][x]) {
							if !lastWasNumber {
								numbers = append(numbers, getNumber(matrix[y], x))
							}
							lastWasNumber = true
						} else {
							lastWasNumber = false
						}
					}
				}
				if len(numbers) == 2 {
					product := 1
					for _, number := range numbers {
						product *= number
					}
					sum += product
				}
			}
		}
	}
	return
}
