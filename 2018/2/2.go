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

	puzzle2(puzzleLines)
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

func puzzle2(puzzleLines []string) {

	var candidates [2]string
	for index, line := range puzzleLines {
		fmt.Printf("%s", line)
		for _, newLine := range puzzleLines[index+1:] {
			fmt.Printf(" ; %s", newLine)
			diff := 0
			for i, char := range line {
				if char != rune(newLine[i]) {
					diff = diff + 1
				}
			}
			if diff < 2 {
				candidates[0] = line
				candidates[1] = newLine
			}
		}
		fmt.Println()
	}
	fmt.Println(candidates)
	output := ""
	for index, char := range candidates[0] {
		if char == rune(candidates[1][index]) {
			output = output + string(char)
		}
	}
	fmt.Println(output)
}
