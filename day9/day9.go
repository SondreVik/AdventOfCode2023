package day9

import (
	"AdventOfCode2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	fmt.Println("Day 9")
	lines := utils.ReadFile("day9/data.txt")
	sequences := [][]int{}
	for _, line := range lines {
		sequence := []int{}
		for _, textNumber := range strings.Split(line, " ") {
			number, err := strconv.Atoi(textNumber)
			utils.CheckError(err)
			sequence = append(sequence, number)
		}
		sequences = append(sequences, sequence)
	}
	fmt.Println(part1(sequences))

}

func part1(sequences [][]int) (total int) {
	fmt.Println("Part 1")
	for _, sequence := range sequences {
		total += calculateNextDigit(sequence, getChangeSequence(sequence))
	}
	return
}

// Helpers

func getChangeSequence(sequence []int) (changeSequence []int) {
	for i := 1; i < len(sequence); i++ {
		changeSequence = append(changeSequence, sequence[i]-sequence[i-1])
	}
	return
}

func calculateNextDigit(sequences []int, changeList []int) int {
	if allZeroes(changeList) {
		return sequences[len(sequences)-1]
	}
	return sequences[len(sequences)-1] + calculateNextDigit(changeList, getChangeSequence(changeList))
}

func allZeroes(sequence []int) bool {
	for _, number := range sequence {
		if number != 0 {
			return false
		}
	}
	return true
}
