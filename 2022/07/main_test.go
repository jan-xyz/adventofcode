package main

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	input := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

	lines := strings.Split(strings.TrimSpace(input), "\n")
	got := parse(lines)

	want := map[string]directory{
		"/": {
			directories: map[string]directory{
				"a": {
					directories: map[string]directory{
						"e": {
							directories: map[string]directory{},
							files:       map[string]int{"i": 584},
						},
					},
					files: map[string]int{
						"f":     29116,
						"g":     2557,
						"h.lst": 62596,
					},
				},
				"d": {
					directories: map[string]directory{},
					files: map[string]int{
						"j":     4060174,
						"d.log": 8033020,
						"d.ext": 5626152,
						"k":     7214296,
					},
				},
			},
			files: map[string]int{
				"b.txt": 14848514,
				"c.dat": 8504156,
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got=%#v want=%#v", got, want)
	}
}

func TestFirstFilterDirectories(t *testing.T) {
	input := map[string]directory{
		"/": {
			directories: map[string]directory{
				"a": {
					directories: map[string]directory{
						"e": {
							directories: map[string]directory{},
							files:       map[string]int{"i": 584},
						},
					},
					files: map[string]int{
						"f":     29116,
						"g":     2557,
						"h.lst": 62596,
					},
				},
				"d": {
					directories: map[string]directory{},
					files: map[string]int{
						"j":     4060174,
						"d.log": 8033020,
						"d.ext": 5626152,
						"k":     7214296,
					},
				},
			},
			files: map[string]int{
				"b.txt": 14848514,
				"c.dat": 8504156,
			},
		},
	}

	got := filterDirectories(input, func(d directory) bool {
		return d.size() < 100000
	})

	want := []directory{
		{
			directories: map[string]directory{
				"e": {
					directories: map[string]directory{},
					files:       map[string]int{"i": 584},
				},
			},
			files: map[string]int{
				"f":     29116,
				"g":     2557,
				"h.lst": 62596,
			},
		},
		{
			directories: map[string]directory{},
			files:       map[string]int{"i": 584},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got=%#v want=%#v", got, want)
	}

	sum := 0
	for _, entry := range got {
		sum += entry.size()
	}
	wantSum := 95437
	if sum != wantSum {
		t.Errorf("sum=%#v wantSum=%#v", sum, wantSum)
	}
}

func TestSecondFilterDirectories(t *testing.T) {
	input := map[string]directory{
		"/": {
			directories: map[string]directory{
				"a": {
					directories: map[string]directory{
						"e": {
							directories: map[string]directory{},
							files:       map[string]int{"i": 584},
						},
					},
					files: map[string]int{
						"f":     29116,
						"g":     2557,
						"h.lst": 62596,
					},
				},
				"d": {
					directories: map[string]directory{},
					files: map[string]int{
						"j":     4060174,
						"d.log": 8033020,
						"d.ext": 5626152,
						"k":     7214296,
					},
				},
			},
			files: map[string]int{
				"b.txt": 14848514,
				"c.dat": 8504156,
			},
		},
	}

	got := filterDirectories(input, func(d directory) bool {
		return d.size() > 8381165
	})

	want := []directory{
		{
			directories: map[string]directory{
				"a": {
					directories: map[string]directory{
						"e": {
							directories: map[string]directory{},
							files:       map[string]int{"i": 584},
						},
					},
					files: map[string]int{
						"f":     29116,
						"g":     2557,
						"h.lst": 62596,
					},
				},
				"d": {
					directories: map[string]directory{},
					files: map[string]int{
						"j":     4060174,
						"d.log": 8033020,
						"d.ext": 5626152,
						"k":     7214296,
					},
				},
			},
			files: map[string]int{
				"b.txt": 14848514,
				"c.dat": 8504156,
			},
		},
		{
			directories: map[string]directory{},
			files: map[string]int{
				"j":     4060174,
				"d.log": 8033020,
				"d.ext": 5626152,
				"k":     7214296,
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got=%#v want=%#v", got, want)
	}

	sort.Slice(got, func(i, j int) bool { return got[i].size() < got[j].size() })

	size := got[0].size()
	wantSize := 24933642
	if size != wantSize {
		t.Errorf("sum=%#v wantSum=%#v", size, wantSize)
	}
}
