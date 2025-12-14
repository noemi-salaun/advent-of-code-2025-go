package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Button []int

type Machine struct {
	lights        int64
	length        int
	buttonsBinary []int64
	buttonsInt    []Button
	joltage       []int
}

type Input []Machine

func main() {
	input := loadInput("day10/input.txt")
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
			var buttonBinary int64
			var button Button
			for _, b := range bs {
				num, _ := strconv.Atoi(b)
				button = append(button, num)
				buttonBinary += int64(math.Pow(2, float64(machine.length-num-1)))
			}
			machine.buttonsInt = append(machine.buttonsInt, button)
			machine.buttonsBinary = append(machine.buttonsBinary, buttonBinary)
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

		for idx := range machine.buttonsBinary {
			res := pressLight(&machine, count+1, state, idx, best)
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

func pressLight(machine *Machine, count int, state int64, buttonIdx int, previousBest int) int {
	var key = fmt.Sprintf("%d-%d-%d", count, state, buttonIdx)

	var res int
	var ok bool

	res, ok = cache[key]
	if !ok {
		res = doPressLight(machine, count, state, buttonIdx, previousBest)
		cache[key] = res
	}

	return res
}

func doPressLight(machine *Machine, count int, state int64, buttonIdx int, previousBest int) int {
	if count > 7 || (previousBest != -1 && count >= previousBest) {
		return -1
	}

	state = state ^ machine.buttonsBinary[buttonIdx]
	if state == machine.lights {
		return count
	}

	var best = -1
	for idx := range machine.buttonsBinary {
		if idx != buttonIdx {
			res := pressLight(machine, count+1, state, idx, best)
			if res != -1 {
				if best == -1 || res < best {
					best = res
				}
			}
		}
	}
	return best
}

func part2(input Input) {
	var result int

	for i, machine := range input {

		cache = map[string]int{}

		var state = make([]int, len(machine.joltage))
		res := solveMachineJoltage(&machine, state, 0, -1)
		if res == -1 {
			log.Fatalf("Cannot resolve machine %d", i)
		}
		fmt.Printf("Machine %d = %d\n", i, res)
		result += res
	}

	fmt.Printf("Part 2 = %d\n", result)
}

func solveMachineJoltage(machine *Machine, state []int, pressesCount int, previousBest int) int {

	var key = fmt.Sprintf("%s-%d", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(state)), ","), "[]"), pressesCount)

	var res int
	var ok bool

	res, ok = cache[key]
	if !ok {
		res = doSolveMachineJoltage(machine, state, pressesCount, previousBest)
		cache[key] = res
	}

	return res
}

func doSolveMachineJoltage(machine *Machine, state []int, pressesCount int, previousBest int) int {
	if previousBest != -1 && pressesCount > previousBest {
		return -1
	}
	if slices.Equal(state, machine.joltage) {
		return pressesCount
	}

	var best = -1

out:
	for _, button := range machine.buttonsInt {

		for _, b := range button {
			if state[b] >= machine.joltage[b] {
				continue out
			}
		}

		copyState := make([]int, len(state))
		copy(copyState, state)

		for _, b := range button {
			copyState[b]++
		}

		var theBest int
		if best != -1 && best < previousBest {
			theBest = best
		} else {
			theBest = previousBest
		}

		res := solveMachineJoltage(machine, copyState, pressesCount+1, theBest)
		if res != -1 && (best == -1 || res < best) {
			best = res
		}
	}

	return best
}
