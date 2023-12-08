package day6

import (
	"AdventOfCode2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func Solve() {
	fmt.Println("Day 6")
	lines := utils.ReadFile("day6/data.txt")
	fmt.Println(part1(deserialize(lines)))
	fmt.Println(part2(deserializeToOneRace(lines)))
}

// deserialize

func deserialize(lines []string) (races []Race) {
	timeRow := lines[0]
	timeRowParts := strings.Split(timeRow, ":")
	distanceRow := lines[1]
	distanceRowParts := strings.Split(distanceRow, ":")
	times := []int{}
	distances := []int{}
	for _, time := range strings.Split(timeRowParts[1], " ") {
		if time == "" {
			continue
		}
		ms, err := strconv.Atoi(strings.TrimSpace(time))
		utils.CheckError(err)
		times = append(times, ms)
	}
	for _, distance := range strings.Split(distanceRowParts[1], " ") {
		if distance == "" {
			continue
		}
		dist, err := strconv.Atoi(distance)
		utils.CheckError(err)
		distances = append(distances, dist)
	}
	for id, time := range times {
		races = append(races, Race{time, distances[id]})
	}
	return
}

func deserializeToOneRace(lines []string) Race {
	timeRow := lines[0]
	timeRowParts := strings.Split(timeRow, ":")
	distanceRow := lines[1]
	distanceRowParts := strings.Split(distanceRow, ":")
	timeString := ""
	distanceString := ""
	for _, time := range strings.Split(timeRowParts[1], " ") {
		if time == "" {
			continue
		}
		timeString += time
	}
	for _, distance := range strings.Split(distanceRowParts[1], " ") {
		if distance == "" {
			continue
		}
		distanceString += distance
	}
	time, err := strconv.Atoi(timeString)
	utils.CheckError(err)
	distance, err := strconv.Atoi(distanceString)
	utils.CheckError(err)
	return Race{time, distance}
}

// parts

func part1(races []Race) (result int) {
	fmt.Println("Part 1")
	result = 1
	for _, race := range races {
		result *= drive(race)
	}
	return
}

func part2(race Race) int {
	fmt.Println("Part 2")
	return drive(race)
}

func drive(race Race) (wins int) {
	for i := 0; i <= race.Time; i++ {
		if race.Distance < getDistance(i, race.Time) {
			wins++
		}
	}
	return
}

// helping methods

func getDistance(pressed, total int) int {
	return pressed * (total - pressed)
}
