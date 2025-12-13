package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Machine struct {
	lights  int64
	length  int
	buttons []int64
	joltage []int
}

type Input []Machine

func main() {
	input := loadInput("day10/input.txt")
	part1(input)
	//part2(input)
}

func loadInput(filepath string) Input {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input Input

	re := regexp.MustCompile(`\[(.*)] (.*) {(.*)}`)

	for scanner.Scan() {
		line := scanner.Text()

		machine := Machine{}

		segments := re.FindStringSubmatch(line)

		var lights = strings.ReplaceAll(segments[1], ".", "0")
		lights = strings.ReplaceAll(lights, "#", "1")
		machine.lights, _ = strconv.ParseInt(lights, 2, 64)
		machine.length = len(lights)

		parts := strings.Fields(segments[2])
		for _, part := range parts {

			bs := strings.Split(part[1:len(part)-1], ",")
			var button int64
			for _, b := range bs {
				num, _ := strconv.Atoi(b)

				button += int64(math.Pow(2, float64(machine.length-num-1)))
			}
			machine.buttons = append(machine.buttons, button)
		}

		for _, b := range strings.Split(segments[3], ",") {
			num, _ := strconv.Atoi(b)
			machine.joltage = append(machine.joltage, num)
		}

		input = append(input, machine)
	}

	return input
}

var cache map[string]int

func part1(input Input) {
	var result int

	for i, machine := range input {

		cache = map[string]int{}

		var count int
		var state int64
		best := -1

		for idx := range machine.buttons {
			res := press(&machine, count+1, state, idx, best)
			if res != -1 {
				if best == -1 || res < best {
					best = res
				}
			}
		}

		if best == -1 {
			log.Fatalf("Cannot resolve machine %d", i)
		}

		result += best
	}

	fmt.Printf("Part 1 = %d\n", result)
}

func part2(input Input) {
	var result int

	fmt.Printf("Part 2 = %d\n", result)
}

func press(machine *Machine, count int, state int64, buttonIdx int, previousBest int) int {
	var key = fmt.Sprintf("%d-%d-%d", count, state, buttonIdx)

	var res int
	var ok bool

	res, ok = cache[key]
	if !ok {
		res = doPress(machine, count, state, buttonIdx, previousBest)
		cache[key] = res
	}

	return res
}

func doPress(machine *Machine, count int, state int64, buttonIdx int, previousBest int) int {
	if count > 7 || (previousBest != -1 && count >= previousBest) {
		return -1
	}

	state = state ^ machine.buttons[buttonIdx]
	if state == machine.lights {
		return count
	}

	var best = -1
	for idx := range machine.buttons {
		if idx != buttonIdx {
			res := press(machine, count+1, state, idx, best)
			if res != -1 {
				if best == -1 || res < best {
					best = res
				}
			}
		}
	}
	return best
}
