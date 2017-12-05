package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	for _, e := range []string{"1122", "1111", "1234", "91212129"} {
		fmt.Println(sumThatShizzle(e))
	}

	thingy, err := ioutil.ReadFile("puzzle_input")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(sumThatShizzle(string(thingy)))

}


func sumThatShizzle(input string) int {
	sum := 0
	for position, character := range(input) {
		next := position + 1
		if next >= len(input) - 1 {
			next = 0
		}
		if input[position] == input[next] {
			sum += int(character - '0')
		}
	}
	return sum
}