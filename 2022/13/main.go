package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("2022/13/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	packetPairs := strings.Split(strings.TrimSpace(string(content)), "\n\n")

	pairs := parse(packetPairs)

	fmt.Println(countCorrectOrder(pairs))
}

func parse(packetPairs []string) [][2][]any {
	ret := [][2][]any{}
	for _, pair := range packetPairs {
		l := strings.Split(strings.TrimSpace(pair), "\n")
		parsedPair := [2][]any{}
		for i, n := range l {
			b := strings.NewReader(n)
			_, _, err := b.ReadRune()
			if err != nil {
				panic("oops")
			}
			parsedPair[i] = parseLine(b)
		}
		ret = append(ret, parsedPair)

	}
	return ret
}

func parseLine(n *strings.Reader) []any {
	parsed := []any{}

	for {
		r, _, err := n.ReadRune()
		if err != nil {
			return parsed
		}
		switch r {
		case ']':
			return parsed
		case '[':
			parsed = append(parsed, parseLine(n))
		case ',':
			continue
		default:
			parsed = append(parsed, int(r-'0'))
		}
	}
}

func countCorrectOrder(pairs [][2][]any) int {
	count := 0
	for _, pair := range pairs {
		if isOrderCorrect(pair) {
			count++
		}
	}
	return count
}

func isOrderCorrect(pair [2][]any) bool {
	if len(pair[0]) == len(pair[1]) {
		for i := 0; i < len(pair[0]); i++ {
			_ = pair[0][i]
			_ = pair[1][i]
		}
	}
	return len(pair[0]) > len(pair[1])
}
