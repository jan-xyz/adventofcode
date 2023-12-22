package main

import "testing"

func TestP1(t *testing.T) {
	input := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	got := p1(input)

	want := 142

	if got != want {
		t.Errorf("got=%d; want=%d", got, want)
	}
}

func TestP2(t *testing.T) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	got := p2(input)

	want := 281

	if got != want {
		t.Errorf("got=%d; want=%d", got, want)
	}
}

func TestP2Extended(t *testing.T) {
	input := []string{
		"sevenine",
		"eighthree",
		"oneight",
		"4",
	}
	got := p2(input)

	want := 224

	if got != want {
		t.Errorf("got=%d; want=%d", got, want)
	}
}
