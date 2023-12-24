package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func p1(input []string, min, max float64) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			hailA := parseHail(input[i])
			hailB := parseHail(input[j])

			t, err := findIntersectionTime(hailA, hailB)
			if err != nil {
				log.Printf("%s: %v -> %v", err, hailA, hailB)
				continue
			}

			t2, err := findIntersectionTime(hailB, hailA)
			if err != nil {
				log.Printf("%s: %v -> %v", err, hailA, hailB)
				continue
			}

			if t < 0 || t2 < 0 {
				log.Printf("IN THE PAST: %v -> %v at (%f)", hailA, hailB, t)
				continue

			}

			intersectX := hailA.x + hailA.vx*t
			intersectY := hailA.y + hailA.vy*t

			if intersectX > max || intersectX < min || intersectY > max || intersectY < min {
				log.Printf("OUTSIDE of test box: %v -> %v at (%f, %f) ", hailA, hailB, intersectX, intersectY)
				continue
			}
			log.Printf("INTERSECT: %v -> %v at (%f, %f) ", hailA, hailB, intersectX, intersectY)
			sum++
		}
		log.Printf("--------")
	}

	return sum
}

func parseHail(hail string) vector {
	sep := strings.Split(hail, "@")

	coordinates := strings.Split(sep[0], ",")
	velocity := strings.Split(sep[1], ",")

	x, err := strconv.ParseFloat(strings.Trim(coordinates[0], " "), 64)
	if err != nil {
		panic(err)
	}
	y, err := strconv.ParseFloat(strings.Trim(coordinates[1], " "), 64)
	if err != nil {
		panic(err)
	}
	z, err := strconv.ParseFloat(strings.Trim(coordinates[2], " "), 64)
	if err != nil {
		panic(err)
	}
	vx, err := strconv.ParseFloat(strings.Trim(velocity[0], " "), 64)
	if err != nil {
		panic(err)
	}
	vy, err := strconv.ParseFloat(strings.Trim(velocity[1], " "), 64)
	if err != nil {
		panic(err)
	}
	vz, err := strconv.ParseFloat(strings.Trim(velocity[2], " "), 64)
	if err != nil {
		panic(err)
	}

	return vector{
		x:  x,
		y:  y,
		z:  z,
		vx: vx,
		vy: vy,
		vz: vz,
	}
}

// vector represents a 2D vector with package-private fields
type vector struct {
	x, y, z    float64
	vx, vy, vz float64
}

// findIntersectionTime calculates the time of intersection between two vectors
func findIntersectionTime(v1, v2 vector) (float64, error) {
	// Solve for the time of intersection
	denom := v1.vx*v2.vy - v1.vy*v2.vx
	if denom == 0 {
		return 0, fmt.Errorf("vectors are parallel, no intersection")
	}

	t := ((v2.x-v1.x)*v2.vy + (v1.y-v2.y)*v2.vx) / denom

	return t, nil
}

func p2(input []string) int {
	return 0
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(strings.Trim(string(content), "\n"), "\n")
	result := p1(input, 200_000_000_000_000, 400_000_000_000_000)
	println("p1", result)
	result = p2(input)
	println("p2", result)
}
