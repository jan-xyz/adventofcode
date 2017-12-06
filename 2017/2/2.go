package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	// test data
	input1 := [][]int{{5, 1, 9, 5}, {7, 5, 3}, {2, 4, 6, 8}}

	fmt.Println(hashThatShizzle(input1))

	// riddle data
	thingy, err := ioutil.ReadFile("puzzle_input")
	if err != nil {
		fmt.Print(err)
	}
	puzzleInput := parseThatShizzle(thingy)
	fmt.Println(hashThatShizzle(puzzleInput))
}

func hashThatShizzle(array [][]int) int {
	sum := 0
	for _, row := range array {
		lowest := -1
		highest := -1
		for _, element := range row {
			if element < lowest || lowest == -1 {
				lowest = element
			}
			if element > highest {
				highest = element
			}
		}
		sum += highest - lowest
	}
	return sum
}

func parseThatShizzle(byteMe []byte) [][]int {
	var intArray [][]int
	byteMeLines := strings.Split(strings.TrimSpace(string(byteMe)), "\n")

	for index, row := range byteMeLines {
		intArray = append(intArray, []int{})
		for _, element := range strings.Split(row, "\t") {
			num, _ := strconv.Atoi(element)
			intArray[index] = append(intArray[index], num)
		}
	}
	return intArray
}
