package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	input := `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

	lines := strings.Split(strings.TrimSpace(input), "\n\n")
	got := parse(lines)

	want := [][2][]any{
		{
			[]any{1, 1, 3, 1, 1},
			[]any{1, 1, 5, 1, 1},
		},
		{
			[]any{[]any{1}, []any{2, 3, 4}},
			[]any{[]any{1}, 4},
		},
		{
			[]any{9},
			[]any{[]any{8, 7, 6}},
		},
		{
			[]any{[]any{4, 4}, 4, 4},
			[]any{[]any{4, 4}, 4, 4, 4},
		},
		{
			[]any{7, 7, 7, 7},
			[]any{7, 7, 7},
		},
		{
			[]any{},
			[]any{3},
		},
		{
			[]any{[]any{[]any{}}},
			[]any{[]any{}},
		},
		{
			[]any{1, []any{2, []any{3, []any{4, []any{5, 6, 7}}}}, 8, 9},
			[]any{1, []any{2, []any{3, []any{4, []any{5, 6, 0}}}}, 8, 9},
		},
	}

	for i, v := range want {
		if !reflect.DeepEqual(got[i], v) {
			t.Errorf("got=%v want=%v", got[i], v)
		}
	}
}
