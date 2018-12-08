package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var testClaims = []string{
	"#1 @ 1,3: 4x4",
	"#2 @ 3,1: 4x4",
	"#3 @ 5,5: 2x2",
}

var claimStringMatcher = regexp.MustCompile("#\\d+ @ (\\d+),(\\d+): (\\d+)x(\\d+)")

//Claim is a representation of an elves claim on the fabric
type Claim struct {
	x      int
	y      int
	width  int
	height int
}

func (c *Claim) parseFromString(claimString string) {
	matches := claimStringMatcher.FindStringSubmatch(claimString)
	c.x, _ = strconv.Atoi(matches[1])
	c.y, _ = strconv.Atoi(matches[2])
	c.width, _ = strconv.Atoi(matches[3])
	c.height, _ = strconv.Atoi(matches[4])
}

func main() {
	file, err := os.Open("2018/3/1.input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	claims := []string{}
	for scanner.Scan() {
		claims = append(claims, scanner.Text())
	}

	one(claims)
}

func one(claims []string) {

	overlap := 0

	fabric := make(map[string]int)

	for _, claimString := range claims {
		claim := new(Claim)
		claim.parseFromString(claimString)
		for currentHeight := 0; currentHeight < claim.height; currentHeight++ {
			for currentWidth := 0; currentWidth < claim.width; currentWidth++ {
				row := claim.y + currentHeight
				column := claim.x + currentWidth
				index := fabric[fmt.Sprintf("%d:%d", row, column)]
				if index == 1 {
					overlap = overlap + 1
				}
				fabric[fmt.Sprintf("%d:%d", row, column)] = index + 1
			}
		}

	}
	fmt.Printf("Overlap: %d", overlap)

}
