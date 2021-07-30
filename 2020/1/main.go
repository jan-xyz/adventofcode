package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("2020/1/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")

	numbers := []int{}

	for _, l := range lines {
		i, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, i)
	}

	for _, i := range numbers {
		for _, j := range numbers {
			if i+j == 2020 {
				println(i * j)
			}
		}
	}
}
