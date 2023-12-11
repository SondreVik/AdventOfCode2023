package day8

import (
	"AdventOfCode2023/utils"
	"fmt"
	"strings"
)

func Solve() {
	fmt.Println("Day 8")
	lines := utils.ReadFile("day8/data.txt")
	instructions := deserializeInstructions(lines[0])
	nodes := lines[1:]
	maps := deserializeMap(nodes)
	fmt.Println(Part1(maps, instructions, "AAA", "ZZZ"))
	fmt.Println(Part2(maps, instructions))
}

func Part1(nodes map[string]map[string]string, instructions []string, startNode string, endNode string) (steps int) {
	fmt.Println("Part 1")
	end := false
	instructionLength := len(instructions)
	currentNode := nodes[startNode]
	for !end {
		instruction := instructions[steps%instructionLength]
		steps++
		key := currentNode[instruction]
		currentNode = nodes[key]
		end = key == endNode
	}
	return
}

func Part2(nodes map[string]map[string]string, instructions []string) int64 {
	fmt.Println("Part 2")
	end := false
	currentNodeKeys := map[string]string{}
	instructionLength := len(instructions)
	for key := range nodes {
		keyEnding := key[len(key)-1]
		if keyEnding == 'A' {
			currentNodeKeys[key] = key
		}
	}
	startKeySteps := map[string]int{}
	i := 0
	for !end {
		stepKeyes := map[string]string{}
		instruction := instructions[i%instructionLength]
		i++
		for startKey, key := range currentNodeKeys {
			stepKeyes[startKey] = nodes[key][instruction]
		}
		currentNodeKeys = stepKeyes
		for startKey, key := range currentNodeKeys {
			if startKeySteps[startKey] != 0 {
				continue
			}
			lastKey := key[len(key)-1]
			if lastKey == 'Z' {
				startKeySteps[startKey] = i
			}
		}
		end = len(startKeySteps) == len(currentNodeKeys)
	}
	endSteps := []int{}
	for _, steps := range startKeySteps {
		endSteps = append(endSteps, steps)
	}

	return int64(getLeastCommonMultiple(endSteps))
}

// Helper methods

func getLeastCommonMultiple(list []int) int {
	multiples := [][]int{}
	for _, item := range list {
		multiples = append(multiples, []int{item})
	}
	common := -1
	for step := 2; ; step++ {
		if common >= 0 {
			break
		}
		for _, multiple := range multiples {
			factor := multiple[0] * step
			multiple = append(multiple, factor)
			if isDivisibleByAll(factor, multiples) {
				common = factor
			}
		}
	}
	return common
}

func isDivisibleByAll(factor int, list [][]int) bool {
	isDivisible := true
	for _, multiples := range list {
		if !isDivisible {
			break
		}
		someDivisible := false
		for _, multiple := range multiples {
			if factor%multiple == 0 {
				someDivisible = true
				break
			}
		}
		isDivisible = isDivisible && someDivisible
	}
	return isDivisible
}

// Deserialize
func deserializeInstructions(text string) []string {
	return strings.Split(text, "")
}

func deserializeMap(lines []string) map[string]map[string]string {
	result := map[string]map[string]string{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "=")
		node := strings.TrimSpace(parts[0])
		crossRoad := strings.Split(parts[1], ", ")
		result[node] = map[string]string{
			"L": strings.TrimSpace(crossRoad[0])[1:],
			"R": crossRoad[1][:len(crossRoad[1])-1],
		}
	}
	return result
}
