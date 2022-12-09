package main

import (
	"reflect"
	"testing"
)

func TestReadInput(t *testing.T) {
	input := []byte(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`)

	got := readInput(input)

	want := [][]int{
		{1000, 2000, 3000},
		{4000},
		{5000, 6000},
		{7000, 8000, 9000},
		{10000},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want=%v, got=%v", want, got)
	}
}

func TestMaxCal(t *testing.T) {
	testCases := []struct {
		desc        string
		numElements int
		want        int
	}{
		{
			desc:        "top element",
			numElements: 1,
			want:        24000,
		},
		{
			desc:        "top three element",
			numElements: 3,
			want:        45000,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			input := [][]int{
				{1000, 2000, 3000},
				{4000},
				{5000, 6000},
				{7000, 8000, 9000},
				{10000},
			}
			got := maxCalories(input, tC.numElements)

			if got != tC.want {
				t.Errorf("want=%d, got=%d", tC.want, got)
			}
		})
	}
}
