package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

type Input []Coord

func main() {
	//input := loadInput("day09/input.txt")
	//part1(input)
	//part2(input)

	// Surface trouvée à la main dans Gimp à partir du PNG dessiné
	fmt.Printf("Part 2 = %d\n", calcArea(Coord{94872, 50262}, Coord{5104, 67629}))
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
		input = append(input, Coord{x, y})
	}

	return input
}

func part1(input Input) {
	var result int

	for i, a := range input {
		for _, b := range input[i+1:] {
			area := calcArea(a, b)
			if area > result {
				result = area
			}
		}
	}

	fmt.Printf("Part 1 = %d\n", result)
}

func part2(input Input) {
	img := image.NewRGBA(image.Rect(0, 0, 1000, 1000))

	green := color.RGBA{G: 255, A: 255}
	red := color.RGBA{R: 255, A: 255}

	lastPoint := input[len(input)-1]
	for _, nextPoint := range input {

		if lastPoint.x == nextPoint.x {
			if lastPoint.y < nextPoint.y {
				for y := lastPoint.y; y <= nextPoint.y; y++ {
					img.Set(lastPoint.x/100, y/100, green)
				}
			} else {
				for y := nextPoint.y; y <= lastPoint.y; y++ {
					img.Set(lastPoint.x/100, y/100, green)
				}
			}
		} else if lastPoint.y == nextPoint.y {
			if lastPoint.x < nextPoint.x {
				for x := lastPoint.x; x <= nextPoint.x; x++ {
					img.Set(x/100, lastPoint.y/100, green)
				}
			} else {
				for x := nextPoint.x; x <= lastPoint.x; x++ {
					img.Set(x/100, lastPoint.y/100, green)
				}
			}
		} else {
			panic("not aligned points")
		}

		lastPoint = nextPoint
	}

	for _, point := range input {
		img.Set(point.x/100, point.y/100, red)
	}

	// Créer un fichier PNG
	outputFile, err := os.Create("day09/output.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Encoder l'image en PNG et l'écrire dans le fichier
	err = png.Encode(outputFile, img)
	if err != nil {
		panic(err)
	}
}

func calcArea(a Coord, b Coord) int {
	return (absInt(a.x-b.x) + 1) * (absInt(a.y-b.y) + 1)
}

func absInt(n int) int {
	return int(math.Abs(float64(n)))
}
