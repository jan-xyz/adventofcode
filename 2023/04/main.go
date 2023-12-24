package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func p1(input []string) int {
	sum := 0
	for _, line := range input {
		println("card:", line)
		line, _ = strings.CutPrefix(line, "Card ")
		cardSep := strings.Index(line, ":")
		game := line[cardSep+1:]
		winSep := strings.Index(game, "|")
		wins := strings.Split(strings.ReplaceAll(strings.Trim(game[:winSep], " "), "  ", " "), " ")
		draws := strings.Split(strings.ReplaceAll(strings.Trim(game[winSep:], " "), "  ", " "), " ")
		score := 0
		println("wins:", game[:winSep])
		println("draws:", game[winSep:])
		for _, draw := range draws {
			if slices.Contains(wins, draw) {
				println("draw:", draw)
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		println("score:", score)
		println("----------")
		sum += score
	}

	return sum
}

func p2(input []string) int {
	sum := 0
	myCards := []struct {
		game  string
		count int
	}{}
	for _, line := range input {
		line, _ = strings.CutPrefix(line, "Card ")
		cardSep := strings.Index(line, ":")
		game := line[cardSep+1:]
		myCards = append(myCards, struct {
			game  string
			count int
		}{game: game, count: 1})
		sum++
	}

	for i := 0; i < len(myCards); i++ {
		game := myCards[i].game
		count := myCards[i].count
		fmt.Printf("game: %v, count: %d\n", game, count)
		winSep := strings.Index(game, "|")
		wins := strings.Split(strings.ReplaceAll(strings.Trim(game[:winSep], " "), "  ", " "), " ")
		draws := strings.Split(strings.ReplaceAll(strings.Trim(game[winSep:], " "), "  ", " "), " ")
		winCount := 0
		for _, draw := range draws {
			if slices.Contains(wins, draw) {
				println(draw)

				winCount++
			}
		}
		sum += winCount * count
		for j := 1; j <= winCount && j+i < len(myCards); j++ {
			myCards[i+j].count += count
		}
		println("-------------")
	}

	return sum
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
