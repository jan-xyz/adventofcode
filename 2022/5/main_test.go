package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	input := []byte(`move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`)

	got := parse(input)

	want := []move{
		{1, 1, 0},
		{3, 0, 2},
		{2, 1, 0},
		{1, 0, 1},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got=%v, want=%v", got, want)
	}
}

func TestMoveOne(t *testing.T) {
	input := []move{
		{1, 1, 0},
		{3, 0, 2},
		{2, 1, 0},
		{1, 0, 1},
	}

	//     [D]
	// [N] [C]
	// [Z] [M] [P]
	//  1   2   3
	stack := [][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}
	got := process(input, stack, processRecurse)

	want := "CMZ"

	if got != want {
		t.Errorf("got=%v, want=%v", got, want)
	}
}

func TestMoveAll(t *testing.T) {
	input := []move{
		{1, 1, 0},
		{3, 0, 2},
		{2, 1, 0},
		{1, 0, 1},
	}

	//     [D]
	// [N] [C]
	// [Z] [M] [P]
	//  1   2   3
	stack := [][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}
	got := process(input, stack, processAll)

	want := "MCD"

	if got != want {
		t.Errorf("got=%v, want=%v", got, want)
	}
}
