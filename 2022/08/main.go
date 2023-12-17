package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type patch [][]int

func (p patch) left(i, j int) []int {
	s := []int{}
	for k := j - 1; k >= 0; k-- {
		s = append(s, p[i][k])
	}
	return s
}

func (p patch) right(i, j int) []int {
	return p[i][j+1:]
}

func (p patch) up(i, j int) []int {
	t := []int{}

	for k := i - 1; k >= 0; k-- {
		t = append(t, p[k][j])
	}
	return t
}

func (p patch) down(i, j int) []int {
	t := []int{}

	for _, line := range p[i+1:] {
		t = append(t, line[j])
	}
	return t
}

func (p patch) isVisible(i, j int) bool {
	currentTree := p[i][j]
	isVisible := func(this int, others []int) bool {
		for _, other := range others {
			if other >= this {
				return false
			}
		}
		return true
	}
	return isVisible(currentTree, p.up(i, j)) ||
		isVisible(currentTree, p.left(i, j)) ||
		isVisible(currentTree, p.right(i, j)) ||
		isVisible(currentTree, p.down(i, j))
}

func (p patch) scenicValue(i, j int) int {
	currentTree := p[i][j]
	scenicValue := func(this int, others []int) int {
		sv := 0
		for _, other := range others {
			sv++
			if other >= this {
				return sv
			}
		}
		return sv
	}
	return scenicValue(currentTree, p.up(i, j)) *
		scenicValue(currentTree, p.left(i, j)) *
		scenicValue(currentTree, p.right(i, j)) *
		scenicValue(currentTree, p.down(i, j))
}

func main() {
	content, err := ioutil.ReadFile("2022/8/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	p := parse(lines)

	fmt.Println(countVisibleTrees(p))
	fmt.Println(maxScenicValue(p))
}

func parse(lines []string) [][]int {
	p := [][]int{}
	for _, l := range lines {
		row := []int{}

		for _, r := range l {
			row = append(row, int(r-'0'))
		}
		p = append(p, row)
	}
	return p
}

func countVisibleTrees(p patch) int {
	height := len(p)
	width := len(p[0])
	visibleTrees := 2*(height+width) - 4 // top, bottom, left, right edges

	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			if p.isVisible(i, j) {
				visibleTrees++
			}
		}
	}

	return visibleTrees
}

func maxScenicValue(p patch) int {
	height := len(p)
	width := len(p[0])

	max := 0
	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			curr := p.scenicValue(i, j)
			if curr > max {
				max = curr
			}
		}
	}

	return max
}
