package day9

import (
	"AdventOfCode2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	fmt.Println("Day 9")
	lines := utils.ReadFile("day9/example.txt")
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

func part1(list [][]int) (total int) {
	fmt.Println("Part 1")
	return
}
