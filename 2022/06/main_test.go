package main

import (
	"testing"
)

func TestProcess(t *testing.T) {
	testCases := []struct {
		desc            string
		input           []byte
		numDistinctChar int
		want            int
	}{
		// first
		{
			desc:            "first - 1",
			input:           []byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"),
			numDistinctChar: 4,
			want:            7,
		},
		{
			desc:            "first - 2",
			input:           []byte("bvwbjplbgvbhsrlpgdmjqwftvncz"),
			numDistinctChar: 4,
			want:            5,
		},
		{
			desc:            "first - 3",
			input:           []byte("nppdvjthqldpwncqszvftbrmjlhg"),
			numDistinctChar: 4,
			want:            6,
		},
		{
			desc:            "first - 4",
			input:           []byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"),
			numDistinctChar: 4,
			want:            10,
		},
		{
			desc:            "first - 5",
			input:           []byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
			numDistinctChar: 4,
			want:            11,
		},

		// second
		{
			desc:            "second - 1",
			input:           []byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"),
			numDistinctChar: 14,
			want:            19,
		},
		{
			desc:            "second - 2",
			input:           []byte("bvwbjplbgvbhsrlpgdmjqwftvncz"),
			numDistinctChar: 14,
			want:            23,
		},
		{
			desc:            "second - 3",
			input:           []byte("nppdvjthqldpwncqszvftbrmjlhg"),
			numDistinctChar: 14,
			want:            23,
		},
		{
			desc:            "second - 4",
			input:           []byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"),
			numDistinctChar: 14,
			want:            29,
		},
		{
			desc:            "second - 5",
			input:           []byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
			numDistinctChar: 14,
			want:            26,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := process(tC.input, tC.numDistinctChar)

			if got != tC.want {
				t.Errorf("got=%d want=%d", got, tC.want)
			}
		})
	}
}
