package main

import "testing"

func TestP1(t *testing.T) {
	input := []string{
		"19, 13, 30 @ -2,  1, -2",
		"18, 19, 22 @ -1, -1, -2",
		"20, 25, 34 @ -2, -2, -4",
		"12, 31, 28 @ -1, -2, -1",
		"20, 19, 15 @  1, -5, -3",
	}

	got := p1(input, 7, 27)

	want := 2
	if got != want {
		t.Errorf("got=%d; want=%d", got, want)
	}
}

func TestP2(t *testing.T) {
	input := []string{
		"19, 13, 30 @ -2,  1, -2",
		"18, 19, 22 @ -1, -1, -2",
		"20, 25, 34 @ -2, -2, -4",
		"12, 31, 28 @ -1, -2, -1",
		"20, 19, 15 @  1, -5, -3",
	}

	got := p2(input)

	want := 47
	if got != want {
		t.Errorf("got=%d; want=%d", got, want)
	}
}
