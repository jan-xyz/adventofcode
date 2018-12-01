package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {

	puzzleInput, _ := ioutil.ReadFile("1.puzzle_input.txt")

	puzzleLines := bytes.Split(puzzleInput, []byte("\n"))

	sum := 0
	for _, element := range puzzleLines {
		summand, _ := strconv.Atoi(string(bytes.TrimSpace(element)))
		sum = sum + summand
	}
	fmt.Println(sum)
}
