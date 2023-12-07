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

type RangeMap struct {
	start            int
	length           int
	transformedStart int
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
				sourceRange:      IdRange{source, source + length - 1},
				destinationRange: IdRange{dest, dest + length - 1},
			},
		)
	}
	return
}

func deserializeRangeMaps2(textLines []string) (maps map[string]map[int]RangeMap) {
	mapKey := ""
	maps = map[string]map[int]RangeMap{
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
		maps[mapKey][source] = RangeMap{
			start:            source,
			length:           length,
			transformedStart: dest,
		}

	}
	return
}
func deserializeSeedList(seedText string) (seeds []int) {
	for _, seed := range strings.Split(seedText, " ") {
		seedNumber, err := strconv.Atoi(seed)
		utils.CheckError(err)
		seeds = append(seeds, seedNumber)
	}
	return
}

func deserializeSeedRanges(seedText string) (seedRanges []IdRange) {
	seedList := strings.Split(seedText, " ")
	for i := 0; i < len(seedList); i += 2 {
		seedNumber1, err := strconv.Atoi(seedList[i])
		utils.CheckError(err)
		seedNumber2, err := strconv.Atoi(seedList[i+1])
		utils.CheckError(err)
		seedRanges = append(seedRanges, IdRange{seedNumber1, seedNumber1 + seedNumber2 - 1})
	}
	return
}

func Solve() {
	fmt.Println("Day 5")
	textLines := utils.ReadFile("day5/data.txt")
	firstLine := textLines[0]
	firstLineParts := strings.Split(firstLine, ": ")
	seeds := deserializeSeedList(firstLineParts[1])
	maps := deserializeRangeMaps(textLines[1:])
	fmt.Println(part1(maps, seeds))
	seedRanges := deserializeSeedRanges(firstLineParts[1])
	fmt.Println(part2(maps, seedRanges))
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
		return source
	}
	matchingRange := data[id]
	sourceOffset := source - matchingRange.sourceRange.from
	return matchingRange.destinationRange.from + sourceOffset
}

func evaluate2(data map[int]RangeMap, seed int) int {
	for _, rangeMap := range data {
		if seed >= rangeMap.start && seed < rangeMap.start+rangeMap.length {
			offset := seed - rangeMap.start
			return rangeMap.transformedStart + offset
		}
	}
	return seed
}

func getLocation(data map[string][]SourceDestinationRanges, seed int) int {
	soil := evaluate(data["seed-to-soil"], seed)
	fertilizer := evaluate(data["soil-to-fertilizer"], soil)
	water := evaluate(data["fertilizer-to-water"], fertilizer)
	light := evaluate(data["water-to-light"], water)
	temperature := evaluate(data["light-to-temperature"], light)
	humidity := evaluate(data["temperature-to-humidity"], temperature)
	return evaluate(data["humidity-to-location"], humidity)
}

func getLocationRanges(data map[string][]SourceDestinationRanges, seedRanges []IdRange) []IdRange {
	soil := evaluateRanges(data["seed-to-soil"], seedRanges)
	fertilizer := evaluateRanges(data["soil-to-fertilizer"], soil)
	water := evaluateRanges(data["fertilizer-to-water"], fertilizer)
	light := evaluateRanges(data["water-to-light"], water)
	temperature := evaluateRanges(data["light-to-temperature"], light)
	humidity := evaluateRanges(data["temperature-to-humidity"], temperature)
	return evaluateRanges(data["humidity-to-location"], humidity)
}

func evaluateRanges(data []SourceDestinationRanges, sourceRanges []IdRange) (result []IdRange) {
	for _, sourceRange := range sourceRanges {
		result = append(result, evaluateRange(data, sourceRange.from, sourceRange.to)...)
	}
	fmt.Println("result: ", result)
	return
}

func evaluateRange(data []SourceDestinationRanges, from, to int) (result []IdRange) {
	noMatchStart := -1
	for i := from; i <= to; i++ {
		id := slices.IndexFunc(data, func(element SourceDestinationRanges) bool {
			return element.sourceRange.from <= i && element.sourceRange.to >= i
		})
		if id < 0 {
			if noMatchStart < 0 {
				fmt.Println("Will set no match to true")
				noMatchStart = i
			}
			continue
		}
		if noMatchStart >= 0 {
			noMatchResult := IdRange{noMatchStart, i - 1}
			fmt.Println("No match result: ", noMatchResult)
			result = append(result, noMatchResult)
			noMatchStart = -1
		}
		matchingRange := data[id]
		resultTo := 0
		if matchingRange.destinationRange.to > to {
			resultTo = to
		} else {
			resultTo = matchingRange.destinationRange.to
		}
		result = append(result, IdRange{i, resultTo})
		i += matchingRange.destinationRange.to
	}
	return
}

func part1(data map[string][]SourceDestinationRanges, seeds []int) int {
	fmt.Println("Part 1")
	lowestLocation := 999999999999999999
	for _, seed := range seeds {
		location := getLocation(data, seed)
		if location < lowestLocation {
			lowestLocation = location
		}
	}
	return lowestLocation
}

func part2(data map[string][]SourceDestinationRanges, seedRanges []IdRange) int {
	fmt.Println("Part 2")
	lowestLocation := 99999999999999999
	locationRanges := getLocationRanges(data, seedRanges)
	fmt.Println(locationRanges)
	for _, locationRange := range locationRanges {
		if locationRange.from < lowestLocation {
			lowestLocation = locationRange.from
		}
	}
	return lowestLocation
}
