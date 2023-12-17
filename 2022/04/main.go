package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type (
	section   int
	shift     [2]int
	shiftPair [2]shift
)

func main() {
	content, err := ioutil.ReadFile("2022/4/input.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := parse(content)
	fmt.Println(countOverlap(input))
}

func parse(b []byte) (sps []shiftPair) {
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	for _, l := range lines {
		shifts := strings.Split(l, ",")
		firstShift := shiftFromString(shifts[0])
		secondShift := shiftFromString(shifts[1])
		sps = append(sps, shiftPair{firstShift, secondShift})
	}
	return
}

func shiftFromString(s string) shift {
	sections := strings.Split(strings.TrimSpace(s), "-")
	start, err := strconv.Atoi(sections[0])
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(sections[1])
	if err != nil {
		panic(err)
	}
	return shift{start, end}
}

func countContaining(s []shiftPair) (c int) {
	for _, sp := range s {
		start1 := sp[0][0]
		end1 := sp[0][1]
		start2 := sp[1][0]
		end2 := sp[1][1]
		if start1 <= start2 && end1 >= end2 {
			c++
		} else if start1 >= start2 && end1 <= end2 {
			c++
		}
	}
	return
}

func countOverlap(s []shiftPair) (c int) {
	for _, sp := range s {
		start1 := sp[0][0]
		end1 := sp[0][1]
		start2 := sp[1][0]
		end2 := sp[1][1]
		if start1 <= end2 && end1 >= start2 {
			c++
		}
	}
	return
}
