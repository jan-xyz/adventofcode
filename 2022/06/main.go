package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("2022/6/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(process(content, 14))
}

func process(content []byte, numDistinctChars int) int {
	for i := numDistinctChars; i < len(content); i++ {
		sub := content[i-numDistinctChars : i]
		if uniqueContent(sub) {
			return i
		}
	}
	return -1
}

func uniqueContent(sub []byte) bool {
	set := map[byte]struct{}{}
	for _, v := range sub {
		set[v] = struct{}{}
	}
	return len(sub) == len(set)
}
