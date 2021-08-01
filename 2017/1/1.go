package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	for _, e := range []string{"1122", "1111", "1234", "91212129"} {
		fmt.Println(sumThatShizzle(e, 1))
	}

	thingy, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(sumThatShizzle(string(thingy), 1))

	for _, e := range []string{"1212", "1221", "123425", "123123", "12131415"} {
		fmt.Println(sumThatShizzle(e, len(e)/2))
	}
	fmt.Println(sumThatShizzle(string(thingy), len(string(thingy))/2))
}

func sumThatShizzle(input string, adder int) int {
	sum := 0
	for position, character := range input {
		next := position + adder
		if next > len(input)-1 {
			next = 0 + (next - len(input))
		}
		if input[position] == input[next] {
			sum += int(character - '0')
		}
	}
	return sum
}
