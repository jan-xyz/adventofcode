package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	input := []byte(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`)

	got := parse(input)

	want := []shiftPair{
		{{2, 4}, {6, 8}},
		{{2, 3}, {4, 5}},
		{{5, 7}, {7, 9}},
		{{2, 8}, {3, 7}},
		{{6, 6}, {4, 6}},
		{{2, 6}, {4, 8}},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got=%v, want=%v", got, want)
	}
}

func TestCountContaining(t *testing.T) {
	input := []shiftPair{
		{{2, 4}, {6, 8}},
		{{2, 3}, {4, 5}},
		{{5, 7}, {7, 9}},
		{{2, 8}, {3, 7}},
		{{6, 6}, {4, 6}},
		{{2, 6}, {4, 8}},
	}

	got := countContaining(input)

	want := 2

	if got != want {
		t.Errorf("got=%v, want=%v", got, want)
	}
}

func TestCountOverlap(t *testing.T) {
	input := []shiftPair{
		{{2, 4}, {6, 8}},
		{{2, 3}, {4, 5}},
		{{5, 7}, {7, 9}},
		{{2, 8}, {3, 7}},
		{{6, 6}, {4, 6}},
		{{2, 6}, {4, 8}},
	}

	got := countOverlap(input)

	want := 4

	if got != want {
		t.Errorf("got=%v, want=%v", got, want)
	}
}
