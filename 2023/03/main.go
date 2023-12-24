package main

import (
	"os"
	"strconv"
	"strings"
)

func p1(input []string) int {
	sum := 0
	for i, line := range input {
		for j, char := range line {
			if !strings.ContainsRune("0123456789.", char) {
				println(string(char), "at(", i, j, ")")
				val := sumAdjacentNumbers(i, j, input)
				sum += val
				println("----------------------")
			}
		}
	}

	return sum
}

func sumAdjacentNumbers(i, j int, input []string) int {
	sum := 0
	// left to the special character
	if j != 0 && input[i][j-1] != '.' {
		num := ""
		rev := reverse(input[i][:j])
		for _, char := range rev {
			if !strings.ContainsRune("0123456789", char) {
				break
			}
			num += string(char)
		}
		num = reverse(num)
		println("leftSum", num)
		leftSum, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		sum += leftSum
	}
	// right to the special character
	if j != len(input[i]) && input[i][j+1] != '.' {
		num := ""
		rev := input[i][j+1:]
		for _, char := range rev {
			if !strings.ContainsRune("0123456789", char) {
				break
			}
			num += string(char)
		}
		println("rightSum", num)
		leftSum, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		sum += leftSum
	}
	// above the special character
	if i != 0 && input[i-1][j] != '.' {
		num := string(input[i-1][j])
		for _, char := range input[i-1][j+1:] {
			if !strings.ContainsRune("0123456789", char) {
				break
			}
			num += string(char)
		}
		num = reverse(num)
		for _, char := range reverse(input[i-1][:j]) {
			if !strings.ContainsRune("0123456789", char) {
				break
			}
			num += string(char)
		}
		num = reverse(num)
		println("topSum", num)
		topSum, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		sum += topSum
	} else {
		// top left split
		if i != 0 && j != 0 && input[i-1][j-1] != '.' {
			num := reverse(string(input[i-1][j-1]))
			for _, char := range reverse(input[i-1][:j-1]) {
				if !strings.ContainsRune("0123456789", char) {
					break
				}
				num += string(char)
			}
			num = reverse(num)
			println("topLeftSum", num)
			topSum, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			sum += topSum
		}
		// top right split
		if i != 0 && j != len(input[i]) && input[i-1][j+1] != '.' {
			num := string(input[i-1][j+1])
			for _, char := range input[i-1][j+2:] {
				if !strings.ContainsRune("0123456789", char) {
					break
				}
				num += string(char)
			}
			println("topRightSum", num)
			topSum, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			sum += topSum
		}
	}
	// below the special character
	if i != len(input) && input[i+1][j] != '.' {
		num := string(input[i+1][j])
		for _, char := range input[i+1][j+1:] {
			if !strings.ContainsRune("0123456789", char) {
				break
			}
			num += string(char)
		}
		num = reverse(num)
		for _, char := range reverse(input[i+1][:j]) {
			if !strings.ContainsRune("0123456789", char) {
				break
			}
			num += string(char)
		}
		num = reverse(num)
		println("botSum", num)
		botSum, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		sum += botSum
	} else {
		// bot left split
		if i != len(input) && j != 0 && input[i+1][j-1] != '.' {
			num := reverse(string(input[i+1][j-1]))
			for _, char := range reverse(input[i+1][:j-1]) {
				if !strings.ContainsRune("0123456789", char) {
					break
				}
				num += string(char)
			}
			num = reverse(num)
			println("botLeftSum", num)
			botSum, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			sum += botSum
		}
		// bot right split
		if i != len(input) && j != len(input[i]) && input[i+1][j+1] != '.' {
			num := string(input[i+1][j+1])
			for _, char := range input[i+1][j+2:] {
				if !strings.ContainsRune("0123456789", char) {
					break
				}
				num += string(char)
			}
			println("botRightSum", num)
			botSum, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			sum += botSum
		}
	}

	return sum
}

// function, which takes a string as
// argument and return the reverse of string.
func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func p2(input []string) int {
	sum := 0
	for i, line := range input {
		for j, char := range line {
			if char == '*' {
				println(string(char), "at(", i, j, ")")
				val := gearRatio(i, j, input)
				sum += val
				println("----------------------")
			}
		}
	}

	return sum
}

func gearRatio(i, j int, input []string) int {
	var gears, left, right, bot, top, topLeft, topRight, botLeft, botRight int
	var err error
	// left to the special character
	if j != 0 && input[i][j-1] != '.' {
		gears++
		num := ""
		rev := reverse(input[i][:j])
		for _, char := range rev {
			if !strings.ContainsRune("0123456789", char) {
				break
			}
			num += string(char)
		}
		num = reverse(num)
		println("leftSum", num)
		left, err = strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
	}
	// right to the special character
	if j != len(input[i]) && input[i][j+1] != '.' {
		gears++
		num := ""
		rev := input[i][j+1:]
		for _, char := range rev {
			if !strings.ContainsRune("0123456789", char) {
				break
			}
			num += string(char)
		}
		println("rightSum", num)
		right, err = strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
	}
	// above the special character
	if i != 0 && input[i-1][j] != '.' {
		gears++
		num := string(input[i-1][j])
		for _, char := range input[i-1][j+1:] {
			if !strings.ContainsRune("0123456789", char) {
				break
			}
			num += string(char)
		}
		num = reverse(num)
		for _, char := range reverse(input[i-1][:j]) {
			if !strings.ContainsRune("0123456789", char) {
				break
			}
			num += string(char)
		}
		num = reverse(num)
		println("topSum", num)
		top, err = strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
	} else {
		// top left split
		if i != 0 && j != 0 && input[i-1][j-1] != '.' {
			gears++
			num := reverse(string(input[i-1][j-1]))
			for _, char := range reverse(input[i-1][:j-1]) {
				if !strings.ContainsRune("0123456789", char) {
					break
				}
				num += string(char)
			}
			num = reverse(num)
			println("topLeftSum", num)
			topLeft, err = strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
		}
		// top right split
		if i != 0 && j != len(input[i]) && input[i-1][j+1] != '.' {
			gears++
			num := string(input[i-1][j+1])
			for _, char := range input[i-1][j+2:] {
				if !strings.ContainsRune("0123456789", char) {
					break
				}
				num += string(char)
			}
			println("topRightSum", num)
			topRight, err = strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
		}
	}
	// below the special character
	if i != len(input) && input[i+1][j] != '.' {
		gears++
		num := string(input[i+1][j])
		for _, char := range input[i+1][j+1:] {
			if !strings.ContainsRune("0123456789", char) {
				break
			}
			num += string(char)
		}
		num = reverse(num)
		for _, char := range reverse(input[i+1][:j]) {
			if !strings.ContainsRune("0123456789", char) {
				break
			}
			num += string(char)
		}
		num = reverse(num)
		println("botSum", num)
		bot, err = strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
	} else {
		// bot left split
		if i != len(input) && j != 0 && input[i+1][j-1] != '.' {
			gears++
			num := reverse(string(input[i+1][j-1]))
			for _, char := range reverse(input[i+1][:j-1]) {
				if !strings.ContainsRune("0123456789", char) {
					break
				}
				num += string(char)
			}
			num = reverse(num)
			println("botLeftSum", num)
			botLeft, err = strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
		}
		// bot right split
		if i != len(input) && j != len(input[i]) && input[i+1][j+1] != '.' {
			gears++
			num := string(input[i+1][j+1])
			for _, char := range input[i+1][j+2:] {
				if !strings.ContainsRune("0123456789", char) {
					break
				}
				num += string(char)
			}
			println("botRightSum", num)
			botRight, err = strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
		}
	}
	if gears != 2 {
		return 0
	}
	println(left, right, top, bot, botLeft, botRight, topLeft, topRight)

	ret := product(filter([]int{left, right, top, bot, botLeft, botRight, topLeft, topRight}, func(num int) bool { return num != 0 }))

	println(ret)
	return ret
}

func filter[T any](slice []T, f func(T) bool) []T {
	var n []T
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}

func product(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	n := slice[0]
	for _, e := range slice[1:] {
		n = n * e
	}
	return n
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(strings.Trim(string(content), "\n"), "\n")
	result := p1(input)
	println("p1", result)
	result = p2(input)
	println("p2", result)
}
