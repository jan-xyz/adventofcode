package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	input := `30373
25512
65332
33549
35390`

	lines := strings.Split(strings.TrimSpace(input), "\n")
	got := parse(lines)

	want := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got=%#v want=%#v", got, want)
	}
}

func TestCount(t *testing.T) {
	input := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}

	got := countVisibleTrees(input)

	want := 21

	if got != want {
		t.Errorf("got=%#v want=%#v", got, want)
	}
}

func TestMaxScenicValue(t *testing.T) {
	input := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}

	got := maxScenicValue(input)

	want := 8

	if got != want {
		t.Errorf("got=%#v want=%#v", got, want)
	}
}

func TestScenicValue(t *testing.T) {
	input := patch{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}

	got := input.scenicValue(1, 2)
	want := 4
	if got != want {
		t.Errorf("got=%#v want=%#v", got, want)
	}

	got = input.scenicValue(3, 2)
	want = 8
	if got != want {
		t.Errorf("got=%#v want=%#v", got, want)
	}
}

func TestPatch(t *testing.T) {
	testCases := []struct {
		desc   string
		inputI int
		inputJ int
		left   []int
		right  []int
		up     []int
		down   []int
	}{
		{
			desc:   "2,2",
			inputI: 2,
			inputJ: 2,
			left:   []int{5, 6},
			right:  []int{3, 2},
			up:     []int{5, 3},
			down:   []int{5, 3},
		},
		{
			desc:   "3,2",
			inputI: 3,
			inputJ: 2,
			left:   []int{3, 3},
			right:  []int{4, 9},
			up:     []int{3, 5, 3},
			down:   []int{3},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			input := patch{
				{3, 0, 3, 7, 3},
				{2, 5, 5, 1, 2},
				{6, 5, 3, 3, 2},
				{3, 3, 5, 4, 9},
				{3, 5, 3, 9, 0},
			}
			got := input.left(tC.inputI, tC.inputJ)
			if !reflect.DeepEqual(got, tC.left) {
				t.Errorf("got=%#v want=%#v", got, tC.left)
			}

			got = input.right(tC.inputI, tC.inputJ)
			if !reflect.DeepEqual(got, tC.right) {
				t.Errorf("got=%#v want=%#v", got, tC.right)
			}

			got = input.up(tC.inputI, tC.inputJ)
			if !reflect.DeepEqual(got, tC.up) {
				t.Errorf("got=%#v want=%#v", got, tC.up)
			}

			got = input.down(tC.inputI, tC.inputJ)
			if !reflect.DeepEqual(got, tC.down) {
				t.Errorf("got=%#v want=%#v", got, tC.down)
			}
		})
	}
}
