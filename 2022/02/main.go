package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type score int

const (
	pointsRock score = iota + 1
	pointsPaper
	pointsScissors
)

const (
	pointsLose score = iota * 3
	pointsDraw
	pointsWin
)

type move string

const (
	rock     move = "rock"
	paper    move = "paper"
	scissors move = "scissors"
)

var points = map[move]map[outcome]score{
	rock: {
		lose: pointsRock + pointsLose,
		draw: pointsRock + pointsDraw,
		win:  pointsRock + pointsWin,
	},
	paper: {
		lose: pointsPaper + pointsLose,
		draw: pointsPaper + pointsDraw,
		win:  pointsPaper + pointsWin,
	},
	scissors: {
		lose: pointsScissors + pointsLose,
		draw: pointsScissors + pointsDraw,
		win:  pointsScissors + pointsWin,
	},
}

var firstColToMove = map[string]move{
	"A": rock,
	"B": paper,
	"C": scissors,
}

var secondColToMove = map[string]move{
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

type outcome string

const (
	lose outcome = "lose"
	draw outcome = "draw"
	win  outcome = "win"
)

var secondColToOutcome = map[string]outcome{
	"X": lose,
	"Y": draw,
	"Z": win,
}

var moveToOutcome = map[move]map[move]outcome{
	rock:     {rock: draw, paper: lose, scissors: win},
	paper:    {rock: win, paper: draw, scissors: lose},
	scissors: {rock: lose, paper: win, scissors: draw},
}

var outcomeToMove = map[move]map[outcome]move{
	rock:     {draw: rock, win: paper, lose: scissors},
	paper:    {lose: rock, draw: paper, win: scissors},
	scissors: {win: rock, lose: paper, draw: scissors},
}

func main() {
	content, err := ioutil.ReadFile("2022/2/input.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := readInput(content, secondTask)
	fmt.Println(calcPoints(input))
}

func calcPoints(input []map[move]outcome) int {
	total := 0
	for _, v := range input {
		for m, o := range v {
			total += int(points[m][o])
		}
	}
	return total
}

type mapper func([]string) (move, outcome)

// readInput parses a file and splits it into elves
func readInput(b []byte, m mapper) []map[move]outcome {
	rounds := []map[move]outcome{}

	lines := strings.Split(strings.TrimSpace(string(b)), "\n")

	for _, l := range lines {
		round := strings.Split(l, " ")
		m, o := m(round)
		rounds = append(rounds, map[move]outcome{m: o})
	}

	return rounds
}

func firstTask(round []string) (m move, o outcome) {
	if len(round) != 2 {
		fmt.Printf("%v\n", round)
		panic("")
	}
	opponent := firstColToMove[round[0]]
	m = secondColToMove[round[1]]
	o = moveToOutcome[opponent][m]

	return
}

func secondTask(round []string) (m move, o outcome) {
	if len(round) != 2 {
		fmt.Printf("%v\n", round)
		panic("")
	}
	opponent := firstColToMove[round[0]]
	o = secondColToOutcome[round[1]]
	m = outcomeToMove[opponent][o]

	return
}
