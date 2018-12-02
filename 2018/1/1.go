package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	puzzle2()
}

func puzzle2() {
	puzzleInput, _ := ioutil.ReadFile("1/1.puzzle_input2.txt")

	puzzleLines := bytes.Split(bytes.TrimSpace(puzzleInput), []byte("\n"))

	listOfFrequencies := make(map[int]int)
	sum := 0
	index := 0
	for {
		fmt.Printf("length: %d ; index: %d\n", len(puzzleLines), index)
		element := puzzleLines[index]
		summand, _ := strconv.Atoi(string(bytes.TrimSpace(element)))
		sum = sum + summand
		//fmt.Printf("summand: %d ; new sum: %d\n", summand, sum)
		id, _ := listOfFrequencies[sum]
		if id == -1 {
			fmt.Println(sum)
			break
		} else {
			listOfFrequencies[sum] = -1
		}
		index = (index + 1) % len(puzzleLines)
	}
}

func puzzle1() {
	puzzleInput, _ := ioutil.ReadFile("1.puzzle_input.txt")

	puzzleLines := bytes.Split(puzzleInput, []byte("\n"))

	sum := 0
	for _, element := range puzzleLines {
		summand, _ := strconv.Atoi(string(bytes.TrimSpace(element)))
		sum = sum + summand
	}
	fmt.Println(sum)
}
