package main

import (
	"bufio"
	"fmt"
	"log"
	"maps"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Junction struct {
	coordA   *Coordinates
	coordB   *Coordinates
	distance float64
}

type Coordinates struct {
	x       int
	y       int
	z       int
	circuit int
}

type Input []Coordinates

func main() {
	input := loadInput("day08/input.txt")
	//part1(input, 1000)
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

		coord := strings.Split(line, ",")
		x, err := strconv.Atoi(coord[0])
		if err != nil {
			log.Fatalf("Cannot convert %s to int", coord[0])
		}
		y, err := strconv.Atoi(coord[1])
		if err != nil {
			log.Fatalf("Cannot convert %s to int", coord[1])
		}
		z, err := strconv.Atoi(coord[2])
		if err != nil {
			log.Fatalf("Cannot convert %s to int", coord[2])
		}

		input = append(input, Coordinates{x, y, z, 0})
	}

	return input
}

func part1(input Input, nbPairs int) {
	var result int

	var junctions []Junction
	var circuits = map[int][]*Coordinates{}

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			dist := distance(input[i], input[j])
			junctions = append(junctions, Junction{
				coordA:   &input[i],
				coordB:   &input[j],
				distance: dist,
			})
		}
	}

	slices.SortFunc(junctions, func(a, b Junction) int {
		if a.distance > b.distance {
			return 1
		}
		if a.distance < b.distance {
			return -1
		}
		return 0
	})

	nextCircuitId := 1
	countPairs := 0

	for i := 0; i < nbPairs; i++ {
		coordA := junctions[i].coordA
		coordB := junctions[i].coordB

		if coordA.circuit != 0 && coordB.circuit == coordA.circuit {
			continue
		}

		if coordA.circuit != 0 && coordB.circuit == 0 {
			coordB.circuit = coordA.circuit
			circuits[coordA.circuit] = append(circuits[coordA.circuit], coordB)
		} else if coordB.circuit != 0 && coordA.circuit == 0 {
			coordA.circuit = coordB.circuit
			circuits[coordB.circuit] = append(circuits[coordB.circuit], coordA)
		} else if coordA.circuit != 0 && coordB.circuit != 0 {
			circuitBId := coordB.circuit
			for _, coord := range circuits[coordB.circuit] {
				coord.circuit = coordA.circuit
				circuits[coordA.circuit] = append(circuits[coordA.circuit], coord)
			}
			delete(circuits, circuitBId)
		} else if coordA.circuit == 0 && coordB.circuit == 0 {
			coordA.circuit = nextCircuitId
			coordB.circuit = nextCircuitId
			circuits[nextCircuitId] = append(circuits[nextCircuitId], coordA, coordB)
			nextCircuitId++
		}

		countPairs++
		if countPairs >= nbPairs {
			break
		}
	}

	orderedCircuits := slices.Collect(maps.Values(circuits))

	slices.SortFunc(orderedCircuits, func(a, b []*Coordinates) int {
		if len(a) > len(b) {
			return -1
		}
		if len(a) < len(b) {
			return 1
		}
		return 0
	})

	result = 1
	for i := 0; i < 3; i++ {
		result *= len(orderedCircuits[i])
	}

	fmt.Printf("Part 1 = %d\n", result)
}

func part2(input Input) {
	var result int

	var junctions []Junction
	var circuits = map[int][]*Coordinates{}

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			dist := distance(input[i], input[j])
			junctions = append(junctions, Junction{
				coordA:   &input[i],
				coordB:   &input[j],
				distance: dist,
			})
		}
	}

	slices.SortFunc(junctions, func(a, b Junction) int {
		if a.distance > b.distance {
			return 1
		}
		if a.distance < b.distance {
			return -1
		}
		return 0
	})

	nextCircuitId := 1

	for _, junction := range junctions {
		coordA := junction.coordA
		coordB := junction.coordB

		if coordA.circuit != 0 && coordB.circuit == coordA.circuit {
			continue
		}

		if coordA.circuit != 0 && coordB.circuit == 0 {
			coordB.circuit = coordA.circuit
			circuits[coordA.circuit] = append(circuits[coordA.circuit], coordB)
		} else if coordB.circuit != 0 && coordA.circuit == 0 {
			coordA.circuit = coordB.circuit
			circuits[coordB.circuit] = append(circuits[coordB.circuit], coordA)
		} else if coordA.circuit != 0 && coordB.circuit != 0 {
			circuitBId := coordB.circuit
			for _, coord := range circuits[coordB.circuit] {
				coord.circuit = coordA.circuit
				circuits[coordA.circuit] = append(circuits[coordA.circuit], coord)
			}
			delete(circuits, circuitBId)
		} else if coordA.circuit == 0 && coordB.circuit == 0 {
			coordA.circuit = nextCircuitId
			coordB.circuit = nextCircuitId
			circuits[nextCircuitId] = append(circuits[nextCircuitId], coordA, coordB)
			nextCircuitId++
		}

		if len(circuits) == 1 && len(slices.Collect(maps.Values(circuits))[0]) == len(input) {
			result = coordA.x * coordB.x
			break
		}
	}

	fmt.Printf("Part 2 = %d\n", result)
}

func distance(coordA Coordinates, coordB Coordinates) float64 {
	// Source - https://stackoverflow.com/a
	// Posted by Ron Warholic, modified by community. See post 'Timeline' for change history
	// Retrieved 2025-12-08, License - CC BY-SA 3.0

	deltaX := float64(coordB.x - coordA.x)
	deltaY := float64(coordB.y - coordA.y)
	deltaZ := float64(coordB.z - coordA.z)

	return math.Sqrt(deltaX*deltaX + deltaY*deltaY + deltaZ*deltaZ)
}
