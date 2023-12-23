package main

import (
	"os"
	"strconv"
	"strings"
)

func p1(input []string) int {
	sum := 0
	for _, line := range input {
		line, _ = strings.CutPrefix(line, "Game ")
		game, err := strconv.Atoi(line[:strings.Index(line, ":")])
		if err != nil {
			panic(err)
		}
		draws := line[strings.Index(line, ":")+2:]
		possible := isPossible(draws)
		if !possible {
			continue
		}
		sum += game
	}

	return sum
}

func isPossible(draws string) bool {
	capacity := map[string]int{
		"red": 12, "green": 13, "blue": 14,
	}
	for _, game := range strings.Split(draws, ";") {
		for _, colorCount := range strings.Split(game, ",") {
			colorCount = strings.Trim(colorCount, " ")
			split := strings.Split(colorCount, " ")
			count, err := strconv.Atoi(split[0])
			if err != nil {
				panic(err)
			}
			color := split[1]
			if count > capacity[color] {
				return false
			}
		}
	}

	return true
}

func p2(input []string) int {
	sum := 0
	for _, line := range input {
		line, _ = strings.CutPrefix(line, "Game ")
		draws := line[strings.Index(line, ":")+2:]
		power := calculatePower(draws)
		sum += power
	}

	return sum
}

func calculatePower(draws string) int {
	minSet := map[string]int{
		"red": 0, "green": 0, "blue": 0,
	}
	for _, game := range strings.Split(draws, ";") {
		for _, colorCount := range strings.Split(game, ",") {
			colorCount = strings.Trim(colorCount, " ")
			split := strings.Split(colorCount, " ")
			count, err := strconv.Atoi(split[0])
			if err != nil {
				panic(err)
			}
			color := split[1]
			if count > minSet[color] {
				minSet[color] = count
			}
		}
	}

	return minSet["red"] * minSet["blue"] * minSet["green"]
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(strings.Trim(string(content), "\n"), "\n")
	result := p1(input)
	println("p1", result)
	result = p2(input)
	println("p2", result)
}
