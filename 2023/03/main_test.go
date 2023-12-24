package main

import "testing"

func TestP1(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	got := p1(input)

	want := 4361
	if got != want {
		t.Errorf("got=%d; want=%d", got, want)
	}
}

func TestP2(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	got := p2(input)

	want := 467835
	if got != want {
		t.Errorf("got=%d; want=%d", got, want)
	}
}
