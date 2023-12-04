package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id   int
	sets []CubeSet
}

type CubeSet struct {
	red   int
	green int
	blue  int
}

type ColorValue struct {
	color string
	value int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func Solve() {
	fmt.Println("Solving day 2")
	dat, err := os.ReadFile("day2/data.txt")
	check(err)
	text := string(dat)
	games := deserializeGames(text)
	fmt.Println(part1(games, CubeSet{12, 13, 14}))
}

func deserializeGames(text string) (games []Game) {
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		gameParts := strings.Split(parts[0], " ")
		gameId, err := strconv.Atoi(gameParts[1])
		check(err)
		games = append(games, Game{gameId, deserializeSets(parts[1])})
	}
	return
}

func deserializeSets(text string) (cubeSets []CubeSet) {
	sets := strings.Split(text, "; ")

	for _, set := range sets {
		colors := strings.Split(strings.Trim(set, " "), ", ")

		colorMaps := make(map[string]int)
		for _, color := range colors {
			colorParts := strings.Split(color, " ")
			colorName := strings.TrimSpace(strings.ToLower(colorParts[1]))
			colorValue, err := strconv.Atoi(strings.TrimSpace(colorParts[0]))
			check(err)
			colorMaps[colorName] = colorValue
		}
		cubeSets = append(cubeSets, CubeSet{colorMaps["red"], colorMaps["green"], colorMaps["blue"]})
	}
	return
}

func part1(games []Game, validSet CubeSet) (sum int) {
	fmt.Println("Part 1")
	for _, game := range games {
		impossible := false
		for _, set := range game.sets {
			if set.blue > validSet.blue || set.green > validSet.green || set.red > validSet.red {
				impossible = true
				break
			}
		}

		if !impossible {
			sum += game.id
		}
	}
	return
}
