package day4

import (
	"AdventOfCode2023/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type ScratchCard struct {
	id             int
	winningNumbers []int
	numbers        []int
}

type ScratchCardCopies struct {
	scratchCard ScratchCard
	copies      int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func DeserializeToList(stringList string) (list []int) {
	for _, part := range strings.Split(stringList, " ") {
		trimmedPart := strings.TrimSpace(part)
		if trimmedPart == "" {
			continue
		}
		number, err := strconv.Atoi(trimmedPart)
		utils.CheckError(err)
		list = append(list, number)
	}
	return
}

func Solve() {
	fmt.Println("Day 4")
	fileLines := utils.ReadFile("day4/data.txt")
	scratchCards := []ScratchCard{}
	for _, line := range fileLines {
		parts := strings.Split(line, ":")
		gameInfoPart := parts[0]
		idParts := strings.Split(gameInfoPart, " ")
		gameId, err := strconv.Atoi(strings.TrimSpace(idParts[len(idParts)-1]))
		utils.CheckError(err)
		scratchCardPart := parts[1]
		scratchCardSections := strings.Split(scratchCardPart, "|")
		scratchCards = append(
			scratchCards,
			ScratchCard{
				gameId,
				DeserializeToList(scratchCardSections[0]),
				DeserializeToList(scratchCardSections[1]),
			},
		)
	}
	fmt.Println(part1(scratchCards))
	fmt.Println(part2(scratchCards))
}

func calculatePoint(numbers []int, winningNumbers []int) int {
	numberOfMatches := getNumberOfMatches(numbers, winningNumbers)

	if numberOfMatches == 0 {
		return 0
	}

	return int(math.Pow(2, float64(numberOfMatches-1)))
}

func getNumberOfMatches(numbers []int, winningNumbers []int) (numberOfMatches int) {
	for _, winningNumber := range winningNumbers {
		if slices.Contains(numbers, winningNumber) {
			numberOfMatches += 1
		}
	}
	return
}

func part1(scratchCards []ScratchCard) (sum int) {
	fmt.Println("Part 1")
	for _, scratchCard := range scratchCards {
		sum += calculatePoint(scratchCard.numbers, scratchCard.winningNumbers)
	}
	return
}

func part2(scratchCards []ScratchCard) (sum int) {
	fmt.Println("Part 2")
	copies := map[int]int{}
	for _, card := range scratchCards {
		copies[card.id] = 1
	}
	for i, scratchCard := range scratchCards {
		numberOfMatches := getNumberOfMatches(scratchCard.numbers, scratchCard.winningNumbers)
		copiesOfThis := copies[scratchCard.id]

		for k := 0; k < numberOfMatches; k++ {
			next := i + k + 1
			if next == len(scratchCards) {
				break
			}
			nextCard := scratchCards[next]
			copies[nextCard.id] += copiesOfThis
		}
	}
	for _, copyCount := range copies {
		sum += copyCount
	}
	return
}
