package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Input []Point

func main() {
	input := loadInput("day09/input.txt")
	//part1(input)
	part2(input)
}

func loadInput(filepath string) Input {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input Input

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Cannot convert value %s to int", parts[0])
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Cannot convert value %s to int", parts[1])
		}
		input = append(input, Point{x, y})
	}

	return input
}

func part1(input Input) {
	var result int

	for i, a := range input {
		for _, b := range input[i+1:] {
			area := newRect(a, b).getArea()
			if area > result {
				result = area
			}
		}
	}

	fmt.Printf("Part 1 = %d\n", result)
}

func part2(input Input) {
	var result int

	for i, a := range input {
		for _, b := range input[i+1:] {
			rect := newRect(a, b)
			area := rect.getArea()
			if area > result && rect.isValid(&input) {
				result = area
			}
		}
	}

	fmt.Printf("Part 2 = %d\n", result)
}

type Rect struct {
	top    int
	right  int
	bottom int
	left   int
}

func newRect(a Point, b Point) *Rect {
	var rect Rect

	if a.x < b.x {
		rect.left = a.x
		rect.right = b.x
	} else {
		rect.left = b.x
		rect.right = a.x
	}

	if a.y < b.y {
		rect.bottom = a.y
		rect.top = b.y
	} else {
		rect.bottom = b.y
		rect.top = a.y
	}

	return &rect
}

func (rect *Rect) getArea() int {
	return ((rect.right - rect.left) + 1) * ((rect.top - rect.bottom) + 1)
}

func (rect *Rect) intersectWith(otherRect *Rect) bool {
	return rect.left < otherRect.right && rect.right > otherRect.left && rect.top > otherRect.bottom && rect.bottom < otherRect.top
}

func (rect *Rect) isValid(input *Input) bool {

	lastPoint := (*input)[len(*input)-1]
	for _, nextPoint := range *input {

		otherRect := newRect(lastPoint, nextPoint)

		if rect.intersectWith(otherRect) {
			return false
		}

		lastPoint = nextPoint
	}

	return true
}
