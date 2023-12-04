
package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// function, which takes a string as
// argument and return the reverse of string.
func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func firstNum(text string) string {
	for _, s := range text {
		_, err := strconv.Atoi(string(s))
		if err == nil {
			return string(s)
		}
	}
	return ""
}

func translateNumerics(text string) string {
	digitMap := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}
	numberedText := text
	for key, digit := range digitMap {
		numberedText = strings.ReplaceAll(numberedText, key, digit)
	}
	return numberedText
}

func lastNum(text string) string {
	return firstNum(reverse(text))
}

func Solve() {

	dat, err := os.ReadFile("day1/data.txt")
	check(err)
	text := string(dat)
	part1(text)
	part2(text)
}

func part1(text string) {
	fmt.Println("Part 1")
	fmt.Println(sumUp(text))
}

func part2(text string) {
	fmt.Println("Part 2")
	fmt.Println(sumUp(translateNumerics(text)))
}

func sumUp(text string) int {

	res1 := strings.Split(text, "\n")
	sum := 0
	for _, s := range res1 {
		translated := s
		digit1 := firstNum(translated)
		digit2 := lastNum(translated)
		digitSum, err := strconv.Atoi(digit1 + digit2)
		check(err)
		if err == nil {
			sum += digitSum
		}
	}
	return sum
}
