package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const items = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type rucksack []string

type parser = func([]byte) []rucksack

func main() {
	content, err := ioutil.ReadFile("2022/3/input.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := secondTask(content)
	fmt.Println(calcPrio(input))
}

func firstTask(b []byte) []rucksack {
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	rucksacks := []rucksack{}
	for _, l := range lines {
		size := len(l) / 2
		rucksacks = append(rucksacks, rucksack{l[:size], l[size:]})
	}
	return rucksacks
}

func secondTask(b []byte) []rucksack {
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	rucksacks := []rucksack{}
	for i := 0; i < len(lines); i += 3 {
		rucksacks = append(rucksacks, rucksack{lines[i], lines[i+1], lines[i+2]})
	}
	return rucksacks
}

func calcPrio(rucksacks []rucksack) int {
	total := 0
	for _, r := range rucksacks {
		if k, ok := findCommon(r); ok {
			prio := strings.IndexRune(items, k) + 1
			total += prio
		}
	}

	return total
}

func findCommon(ruck rucksack) (r rune, found bool) {
	for _, r := range ruck[0] {
		found = contains(r, ruck[1], ruck[2:]...)
		if found {
			return r, true
		}
	}
	return r, false
}

func contains(r rune, rucks string, others ...string) bool {
	if strings.ContainsRune(rucks, r) {
		if len(others) > 0 {
			return contains(r, others[0], others[1:]...)
		}
		return true
	}
	return false
}
