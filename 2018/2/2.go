package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	puzzleInputBytes, _ := ioutil.ReadFile("2/1.input.txt")
	puzzleInput := string(puzzleInputBytes)

	puzzleLines := strings.Split(strings.TrimSpace(puzzleInput), "\n")

	puzzle1(puzzleLines)
}

func puzzle1(puzzleLines []string) {
	twoLetters := 0
	threeLetters := 0
	for _, line := range puzzleLines {
		fmt.Printf("%s", line)
		letterCollection := make(map[rune]int)
		for _, letter := range line {
			count := letterCollection[letter]
			letterCollection[letter] = count + 1
		}
		addTwo := 0
		addThree := 0
		for letter, count := range letterCollection {
			if count == 2 {
				fmt.Printf(" has 2 %c", letter)
				addTwo = 1
			} else if count == 3 {
				fmt.Printf(" has 3 %c", letter)
				addThree = 1
			}
		}
		twoLetters = twoLetters + addTwo
		threeLetters = threeLetters + addThree
		fmt.Println()
	}
	fmt.Printf("%d * %d = %d\n", twoLetters, threeLetters, twoLetters*threeLetters)

}

func puzzle2(puzzleLines [][]byte) {
}
