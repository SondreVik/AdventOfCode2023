package day7

import (
	"AdventOfCode2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bid   int
}

func Solve() {
	fmt.Println("Day 7")
	lines := utils.ReadFile("day7/data.txt")
	hands := deserialize(lines)
	fmt.Println(part1(hands))
	fmt.Println(part2(hands))
}

//Parts

func part1(hands []Hand) (result int) {
	fmt.Println("Part 1")
	sorted := quickSort(hands)
	for id, hand := range sorted {
		result += (id + 1) * hand.bid
	}
	return
}

func part2(hands []Hand) (result int) {
	fmt.Println("Part 2")
	sorted := quickSortJ(hands)
	for id, hand := range sorted {
		result += (id + 1) * hand.bid
	}
	return
}

// Deserialize
func deserialize(input []string) (hands []Hand) {
	for _, line := range input {
		parts := strings.Split(line, " ")
		bid, err := strconv.Atoi(parts[1])
		utils.CheckError(err)
		hands = append(hands, Hand{parts[0], bid})
	}
	return
}

// Sorting

func quickSort(hands []Hand) (sortedHands []Hand) {
	if len(hands) == 0 {
		return hands
	}
	less := []Hand{}
	greater := []Hand{}
	first := hands[0]
	rest := hands[1:]
	for _, hand := range rest {
		if greaterThan(first.cards, hand.cards) {
			less = append(less, hand)
		} else {
			greater = append(greater, hand)
		}
	}
	return append(append(quickSort(less), first), quickSort(greater)...)
}

func quickSortJ(hands []Hand) (sortedHands []Hand) {
	if len(hands) == 0 {
		return hands
	}
	less := []Hand{}
	greater := []Hand{}
	first := hands[0]
	rest := hands[1:]
	for _, hand := range rest {
		if greaterThanJ(first.cards, hand.cards) {
			less = append(less, hand)
		} else {
			greater = append(greater, hand)
		}
	}
	return append(append(quickSortJ(less), first), quickSortJ(greater)...)
}

func greaterThan(cards1, cards2 string) bool {
	typeValCards1 := getTypeValue(cards1)
	typeValCards2 := getTypeValue(cards2)
	if typeValCards1 != typeValCards2 {
		return typeValCards1 > typeValCards2
	}
	for cardId := range cards1 {
		cardVal1 := getCardValue(rune(cards1[cardId]))
		cardVal2 := getCardValue(rune(cards2[cardId]))
		if cardVal1 == cardVal2 {
			continue
		}
		return cardVal1 > cardVal2
	}
	return false
}
func greaterThanJ(cards1, cards2 string) bool {
	typeValCards1 := getTypeValue(translateHand(cards1))
	typeValCards2 := getTypeValue(translateHand(cards2))
	if typeValCards1 != typeValCards2 {
		return typeValCards1 > typeValCards2
	}
	for cardId := range cards1 {
		cardVal1 := getCardValueJ(rune(cards1[cardId]))
		cardVal2 := getCardValueJ(rune(cards2[cardId]))
		if cardVal1 == cardVal2 {
			continue
		}
		return cardVal1 > cardVal2
	}
	return false
}

// Helpers
func getTypeValue(hand string) int {
	types := map[string]int{
		"FiveOfAKind":  70,
		"FourOfAKind":  60,
		"FullHouse":    50,
		"ThreeOfAKind": 40,
		"TwoPair":      30,
		"OnePair":      20,
		"HighCard":     10,
	}
	return types[getType(hand)]
}

func getType(hand string) string {
	handMap := map[rune]int{}
	for _, char := range hand {
		handMap[char]++
	}

	switch len(handMap) {
	case len(hand):
		return "HighCard"
	case 1:
		return "FiveOfAKind"
	case 2:
		for _, amount := range handMap {
			switch amount {
			case 4, 1:
				return "FourOfAKind"
			case 3, 2:
				return "FullHouse"
			}
		}
	case 3:
		for _, amount := range handMap {
			if amount == 3 {
				return "ThreeOfAKind"
			}
		}
		return "TwoPair"
	}

	return "OnePair"
}

func translateHand(hand string) string {
	if !strings.Contains(hand, "J") {
		return hand
	}
	chars := map[rune]int{}
	for _, char := range hand {
		if char == 'J' {
			continue
		}
		chars[char]++
	}
	highestCharCount := 0
	var highestChar rune
	for char, count := range chars {
		if highestCharCount < count || (count == highestCharCount && getCardValue(highestChar) < getCardValue(char)) {
			highestCharCount = count
			highestChar = char
		}
	}
	return strings.ReplaceAll(hand, "J", string(highestChar))
}

func getCardValue(card rune) int {
	cardValues := map[rune]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
	return cardValues[card]
}

func getCardValueJ(card rune) int {
	if card == 'J' {
		return 1
	}
	return getCardValue(card)
}
