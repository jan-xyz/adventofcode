package main

import (
	"testing"
)

func TestCalcPoints(t *testing.T) {
	testCases := []struct {
		desc string
		want int
		m    mapper
	}{
		{
			desc: "first",
			want: 15,
			m:    firstTask,
		},
		{
			desc: "second",
			want: 12,
			m:    secondTask,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			input := []byte(`A Y
B X
C Z`)
			rounds := readInput(input, tC.m)

			got := calcPoints(rounds)
			if got != tC.want {
				t.Errorf("want=%d, got=%d", tC.want, got)
			}
		})
	}
}
