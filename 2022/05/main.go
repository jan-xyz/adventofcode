package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type move [3]int

func main() {
	content, err := ioutil.ReadFile("2022/5/input.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := parse(content)

	// 	       [M]     [B]             [N]
	// [T]     [H]     [V] [Q]         [H]
	// [Q]     [N]     [H] [W] [T]     [Q]
	// [V]     [P] [F] [Q] [P] [C]     [R]
	// [C]     [D] [T] [N] [N] [L] [S] [J]
	// [D] [V] [W] [R] [M] [G] [R] [N] [D]
	// [S] [F] [Q] [Q] [F] [F] [F] [Z] [S]
	// [N] [M] [F] [D] [R] [C] [W] [T] [M]
	//  1   2   3   4   5   6   7   8   9
	stacks := [][]string{
		{"N", "S", "D", "C", "V", "Q", "T"},
		{"M", "F", "V"},
		{"F", "Q", "W", "D", "P", "N", "H", "M"},
		{"D", "Q", "R", "T", "F"},
		{"R", "F", "M", "N", "Q", "H", "V", "B"},
		{"C", "F", "G", "N", "P", "W", "Q"},
		{"W", "F", "R", "L", "C", "T"},
		{"T", "Z", "N", "S"},
		{"M", "S", "D", "J", "R", "Q", "H", "N"},
	}

	fmt.Println(process(input, stacks, processAll))
}

var parser = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func parse(b []byte) (moves []move) {
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	for _, l := range lines {
		matches := parser.FindStringSubmatch(l)
		amount, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		from, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(matches[3])
		if err != nil {
			panic(err)
		}
		moves = append(moves, move{amount, from - 1, to - 1})
	}
	return
}

type processor = func(amount int, from, to []string) ([]string, []string)

func process(moves []move, stacks [][]string, proc processor) (s string) {
	for _, m := range moves {
		from, to := proc(m[0], stacks[m[1]], stacks[m[2]])
		stacks[m[1]] = from
		stacks[m[2]] = to
	}

	for _, v := range stacks {
		s += v[len(v)-1]
	}
	return
}

func processRecurse(amount int, from, to []string) ([]string, []string) {
	if amount == 0 {
		return from, to
	}
	newFrom := from[:len(from)-1]
	newTo := append(to, from[len(from)-1])
	return processRecurse(amount-1, newFrom, newTo)
}

func processAll(amount int, from, to []string) ([]string, []string) {
	newFrom := from[:len(from)-amount]
	newTo := append(to, from[len(from)-amount:]...)
	return newFrom, newTo
}
