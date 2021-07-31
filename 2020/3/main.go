package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("2020/3/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")

	counterOne := 0
	counterTwo := 0
	counterThree := 0
	counterFour := 0
	counterFive := 0

	for i, line := range lines {
		colOne := i % len(line)
		colTwo := i * 3 % len(line)
		colThree := i * 5 % len(line)
		colFour := i * 7 % len(line)

		if string(line[colOne]) == "#" {
			counterOne++
		}
		if string(line[colTwo]) == "#" {
			counterTwo++
		}
		if string(line[colThree]) == "#" {
			counterThree++
		}
		if string(line[colFour]) == "#" {
			counterFour++
		}
	}

	for i := 0; i < len(lines); i = i + 2 {
		line := lines[i]
		col := i / 2 % len(line)
		if string(line[col]) == "#" {
			counterFive++
		}
	}

	println(counterOne * counterTwo * counterThree * counterFour * counterFive)
}
