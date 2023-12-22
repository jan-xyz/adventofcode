package main

import (
	"log"
	"os"
	"strings"
)

func p1(input []string) int {
	replacements := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
	sum := 0
	for _, line := range input {
		println("line: ", line)
		a := findFirst(line, replacements)
		b := findLast(line, replacements)
		x := a + b
		println("sum: ", x, "\n----")
		sum += x
	}
	return sum
}

func p2(input []string) int {
	replacements := map[string]int{
		"zero": 0, "0": 0,
		"one": 1, "1": 1,
		"two": 2, "2": 2,
		"three": 3, "3": 3,
		"four": 4, "4": 4,
		"five": 5, "5": 5,
		"six": 6, "6": 6,
		"seven": 7, "7": 7,
		"eight": 8, "8": 8,
		"nine": 9, "9": 9,
	}
	sum := 0
	for _, line := range input {
		a := findFirst(line, replacements)
		b := findLast(line, replacements)
		x := a + b
		sum += x
	}
	return sum
}

func findFirst(line string, lookup map[string]int) int {
	for i := 0; i <= len(line); i++ {
		for k, v := range lookup {
			x := line[:i]
			if strings.Contains(x, k) {
				return v * 10
			}
		}
	}
	panic("cannot be")
}

func findLast(line string, lookup map[string]int) int {
	for i := 0; i <= len(line); i++ {
		for k, v := range lookup {
			x := line[len(line)-i:]
			if strings.Contains(x, k) {
				return v
			}
		}
	}
	panic("cannot be")
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(strings.Trim(string(content), "\n"), "\n")
	result := p2(input)
	println(result)
}
