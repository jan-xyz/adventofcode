package main

import (
	"reflect"
	"testing"
)

func TestFirstTask(t *testing.T) {
	input := []byte(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg`)

	got := firstTask(input)

	want := []rucksack{
		{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
		{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
		{"PmmdzqPrV", "vPwwTWBwg"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got=%v, want=%v", got, want)
	}
}

func TestSecondTask(t *testing.T) {
	input := []byte(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg`)

	got := secondTask(input)

	want := []rucksack{
		{
			"vJrwpWtwJgWrhcsFMMfFFhFp",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			"PmmdzqPrVvPwwTWBwg",
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got=%v, want=%v", got, want)
	}
}

func TestCalcPoints(t *testing.T) {
	testCases := []struct {
		desc string
		p    parser
		want int
	}{
		{
			desc: "first",
			p:    firstTask,
			want: 157,
		},
		{
			desc: "second",
			p:    secondTask,
			want: 70,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			input := []byte(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`)

			rucksacks := tC.p(input)
			got := calcPrio(rucksacks)
			if got != tC.want {
				t.Errorf("want=%d, got=%d", tC.want, got)
			}
		})
	}
}
