package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type directory struct {
	directories map[string]directory
	files       map[string]int
}

func (d directory) size() int {
	s := 0
	for _, dir := range d.directories {
		s = s + dir.size()
	}
	for _, fileSize := range d.files {
		s += fileSize
	}
	return s
}

func main() {
	content, err := ioutil.ReadFile("2022/7/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	fs := parse(lines)

	// first
	dirs := filterDirectories(fs, func(d directory) bool {
		return d.size() < 100000
	})
	sum := 0
	for _, d := range dirs {
		sum += d.size()
	}
	fmt.Println(sum)

	// second
	freeSpace := 70000000 - fs["/"].size()
	dirs = filterDirectories(fs, func(d directory) bool {
		return d.size() >= 30000000-freeSpace
	})
	sort.Slice(dirs, func(i, j int) bool { return dirs[i].size() < dirs[j].size() })

	size := dirs[0].size()
	fmt.Println(size)
}

func parse(lines []string) map[string]directory {
	d := map[string]directory{"/": {directories: map[string]directory{}, files: map[string]int{}}}
	currentPath := []string{}

	for _, l := range lines {
		if l == "$ cd .." {
			currentPath = currentPath[:len(currentPath)-1]
			continue
		} else if strings.HasPrefix(l, "$ cd") {
			nextDir := strings.Split(l, " ")[2]
			currentPath = append(currentPath, nextDir)
			continue
		} else if l == "$ ls" {
			continue
		} else if strings.HasPrefix(l, "dir ") {
			newDir := strings.Split(l, " ")[1]
			currentDir := d
			for _, dir := range currentPath {
				currentDir = currentDir[dir].directories
			}
			currentDir[newDir] = directory{directories: map[string]directory{}, files: map[string]int{}}
			continue
		}
		fileEntry := strings.Split(l, " ")
		size, err := strconv.Atoi(fileEntry[0])
		if err != nil {
			panic(err)
		}
		name := fileEntry[1]
		currentDir := d["/"]
		for _, dir := range currentPath[1:] {
			currentDir = currentDir.directories[dir]
		}
		currentDir.files[name] = size
	}

	return d
}

type predicate = func(d directory) bool

func filterDirectories(d map[string]directory, pred predicate) []directory {
	dirs := []directory{}
	for _, dir := range d {
		if pred(dir) {
			dirs = append(dirs, dir)
		}
		dirs = append(dirs, filterDirectories(dir.directories, pred)...)
	}
	return dirs
}
