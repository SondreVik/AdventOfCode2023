package day5

import (
	"AdventOfCode2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Mapping struct {
	destination int
	source      int
}

type IdRange struct {
	from int
	to   int
}

type SourceDestinationRanges struct {
	destinationRange IdRange
	sourceRange      IdRange
}

func deserializeRangeMaps(textLines []string) (maps map[string][]SourceDestinationRanges) {
	mapKey := ""
	maps = map[string][]SourceDestinationRanges{
		"seed-to-soil":            {},
		"soil-to-fertilizer":      {},
		"fertilizer-to-water":     {},
		"water-to-light":          {},
		"light-to-temperature":    {},
		"temperature-to-humidity": {},
		"humidity-to-location":    {},
	}
	for _, line := range textLines {
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "map") {
			mapKey = strings.Split(line, " ")[0]
			continue
		}
		parts := strings.Split(line, " ")
		dest, err := strconv.Atoi(parts[0])
		utils.CheckError(err)
		source, err := strconv.Atoi(parts[1])
		utils.CheckError(err)
		length, err := strconv.Atoi(parts[2])
		utils.CheckError(err)
		maps[mapKey] = append(
			maps[mapKey],
			SourceDestinationRanges{
				sourceRange:      IdRange{source, source + length},
				destinationRange: IdRange{dest, dest + length},
			},
		)
	}
	return
}

func Solve() {
	fmt.Println("Day 5")
	textLines := utils.ReadFile("day5/data.txt")
	firstLine := textLines[0]
	firstLineParts := strings.Split(firstLine, ": ")
	seeds := []int{}
	for _, seed := range strings.Split(firstLineParts[1], " ") {
		seedNumber, err := strconv.Atoi(seed)
		utils.CheckError(err)
		seeds = append(seeds, seedNumber)
	}
	maps := deserializeRangeMaps(textLines[1:])
	fmt.Println(part1(maps, seeds))
}

func getValueOrDefault(value int, fallBack int) int {
	if value <= 0 {
		return fallBack
	}
	return value
}

func evaluate(data []SourceDestinationRanges, source int) int {
	id := slices.IndexFunc(data, func(element SourceDestinationRanges) bool {
		return element.sourceRange.from <= source && element.sourceRange.to >= source
	})
	if id < 0 {
		fmt.Println("no match")
		return source
	}
	matchingRange := data[id]
	fmt.Println(" match: ", matchingRange)
	sourceOffset := source - matchingRange.sourceRange.from
	return matchingRange.destinationRange.from + sourceOffset
}

func part1(data map[string][]SourceDestinationRanges, seeds []int) int {
	fmt.Println("Part 1")
	lowestLocation := 999999999999999999
	//fmt.Println("seed-to-soil: ", data["seed-to-soil"])
	for _, seed := range seeds {
		fmt.Println("seed: ", seed)
		soil := evaluate(data["seed-to-soil"], seed)
		fmt.Println("soil: ", soil)
		fertilizer := evaluate(data["soil-to-ferilizer"], soil)
		fmt.Println("fertilizer: ", fertilizer)
		water := evaluate(data["ferilizer-to-water"], fertilizer)
		fmt.Println("water: ", water)
		light := evaluate(data["water-to-light"], water)
		fmt.Println("light", light)
		temperature := evaluate(data["light-to-temperature"], light)
		fmt.Println("temperature: ", temperature)
		humidity := evaluate(data["temperature-to-light"], temperature)
		fmt.Println("humidity: ", humidity)
		location := evaluate(data["humidity-to-location"], humidity)
		fmt.Println("location: ", location)
		if location < lowestLocation {
			lowestLocation = location
		}
	}
	return lowestLocation
}
