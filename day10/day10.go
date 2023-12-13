package day10

import (
	"AdventOfCode2023/utils"
	"fmt"
	"slices"
	"strings"
)

type Direction int
type Pipe string

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
	NULL
)

const (
	NORTHSOUTH Pipe = "|"
	EASTWEST   Pipe = "-"
	NORTHEAST  Pipe = "L"
	NORTHWEST  Pipe = "J"
	SOUTHWEST  Pipe = "7"
	SOUTHEAST  Pipe = "F"
	START      Pipe = "S"
	NONE       Pipe = "."
)

type coordinate struct {
	x int
	y int
}

type traversal struct {
	x    int
	y    int
	from Direction
}

func Solve() {
	fmt.Println("Day 10")
	lines := utils.ReadFile("day10/example.txt")
	pipes := [][]Pipe{}
	for id, line := range lines {
		pipes = append(pipes, []Pipe{})
		for _, pipe := range strings.Split(line, "") {
			pipes[id] = append(pipes[id], Pipe(pipe))
		}
	}
	fmt.Println(part1(pipes))
}

func part1(pipes [][]Pipe) int {
	fmt.Println("Part 1")
	return getStepsThroughLoop(pipes) / 2
}
func part2() {

}

// helpers

func getStepsThroughLoop(pipes [][]Pipe) int {
	start := findStart(pipes)
	validDirections := getValidDirections(pipes, start)
	fmt.Println("Valid directions: ", validDirections)
	fmt.Println("Start: ", start)
	startTraversal := traversal{start.x, start.y, getOposite(validDirections[0])}
	fmt.Println("Start traversal: ", startTraversal)
	firstNext := nextStep(pipes, startTraversal)
	fmt.Println("First next: ", firstNext)
	return traverse(pipes, firstNext)
}

func traverse(pipes [][]Pipe, currentHeading traversal) int {
	stepNumber := 1
	fmt.Println("current heading: ", currentHeading)
	if pipes[currentHeading.y][currentHeading.x] == START {
		return stepNumber
	}
	next := nextStep(pipes, currentHeading)
	return stepNumber + traverse(pipes, next)
}

func nextStep(pipes [][]Pipe, currentHeading traversal) traversal {
	direction := getDirection(pipes, currentHeading)
	// fmt.Println("Next step direction: ", direction)
	return traversal{
		currentHeading.x + nextXOffset(direction),
		currentHeading.y + nextYOffset(direction),
		getOposite(direction),
	}
}

func nextXOffset(newDirection Direction) int {
	switch newDirection {
	case EAST:
		return +1
	case WEST:
		return -1
	}
	return 0
}

func nextYOffset(newDirection Direction) int {
	switch newDirection {
	case SOUTH:
		return +1
	case NORTH:
		return -1
	}
	return 0
}

func findStart(pipes [][]Pipe) coordinate {
	for y, line := range pipes {
		x := slices.Index(line, START)
		if x >= 0 {
			return coordinate{x, y}
		}
	}
	return coordinate{}
}

func getDirections(pipe Pipe) []Direction {
	switch pipe {
	case NORTHSOUTH:
		return []Direction{NORTH, SOUTH}
	case EASTWEST:
		return []Direction{EAST, WEST}
	case NORTHEAST:
		return []Direction{NORTH, EAST}
	case NORTHWEST:
		return []Direction{NORTH, WEST}
	case SOUTHWEST:
		return []Direction{SOUTH, WEST}
	case SOUTHEAST:
		return []Direction{SOUTH, EAST}
	case START:
		return []Direction{NORTH, EAST, SOUTH, WEST}
	}

	return []Direction{}
}

func getValidDirections(pipes [][]Pipe, coord coordinate) (validDirections []Direction) {
	north := pipes[coord.y-1][coord.x]
	east := pipes[coord.y][coord.x+1]
	south := pipes[coord.y+1][coord.x]
	west := NONE
	if coord.x > 0 {
		west = pipes[coord.y][coord.x-1]
	}
	fmt.Println("get valid directions from ", north, east, south, west)
	if north != NONE {
		validDirections = append(validDirections, NORTH)
	}
	if east != NONE {
		validDirections = append(validDirections, EAST)
	}
	if south != NONE {
		validDirections = append(validDirections, SOUTH)
	}
	if west != NONE {
		validDirections = append(validDirections, WEST)
	}
	fmt.Println("Valid directions: ", validDirections)
	return
}
func getOposite(direction Direction) Direction {
	switch direction {
	case NORTH:
		return SOUTH
	case EAST:
		return WEST
	case SOUTH:
		return NORTH
	case WEST:
		return EAST
	}
	return direction
}

func getDirection(pipes [][]Pipe, currentHeading traversal) Direction {
	directions := getDirections(pipes[currentHeading.y][currentHeading.x])
	for _, direction := range directions {
		if currentHeading.from != direction {
			return direction
		}
	}
	return NULL
}
