package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("2022/1/input.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := readInput(content)
	fmt.Println(maxCalories(input, 3))
}

// readInput parses a file and splits it into elves
func readInput(b []byte) [][]int {
	elves := [][]int{}

	lines := strings.Split(strings.TrimSpace(string(b)), "\n")

	elf := []int{}
	for _, l := range lines {
		if l == "" {
			elves = append(elves, elf)
			elf = []int{}
			continue
		}
		i, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		elf = append(elf, i)
	}
	elves = append(elves, elf)

	return elves
}

// maxCalories returns the elf with the highest calories
func maxCalories(list [][]int, num int) int {
	totalCalories := []int{}
	for _, v := range list {
		s := sum(v)
		totalCalories = append(totalCalories, s)
	}
	sort.Slice(totalCalories, func(i, j int) bool {
		return totalCalories[i] > totalCalories[j]
	})
	return sum(totalCalories[:num])
}

func sum(ints []int) int {
	total := 0
	for _, v := range ints {
		total += v
	}
	return total
}
