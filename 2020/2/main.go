package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var pattern = regexp.MustCompile(`(?P<lower>\d+)-(?P<upper>\d+) (?P<letter>\w): (?P<password>\w+)`)

func main() {
	b, err := ioutil.ReadFile("2020/2/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")

	count := 0

	for _, line := range lines {
		matches := pattern.FindAllStringSubmatch(line, -1)[0]
		lower, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		lower--
		upper, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		upper--
		letter := matches[3]
		str := matches[4]

		if str[upper] != str[lower] && (string(str[upper]) == letter || string(str[lower]) == letter) {
			println(line)
			count++
		}
	}
	println(count)
}
