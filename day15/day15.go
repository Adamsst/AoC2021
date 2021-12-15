package main

import (
	"AoC2021/helpers"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

type vertex struct {
	x, y int
}

const part2Multiplier = 5

func main() {

	partPtr := flag.String("part", "1", "which part to run")
	flag.Parse()
	part, err := strconv.Atoi(*partPtr)
	if err != nil {
		fmt.Println("Pass valid Part number, 1 or 2, as -part flag")
		os.Exit(1)
	}

	input, err := helpers.ReadLinesToString("input.txt")
	if err != nil {
		fmt.Printf("Input read err: %s", err.Error())
		os.Exit(1)
	}

	// For this I used the Algorithm section of
	// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
	riskLevels := make(map[vertex]int, len(input)*len(input[0]))
	distances := make(map[vertex]int, len(input)*len(input[0]))
	nextLoc := make(map[vertex]bool) // Use a map because its easy to call delete. Value is unimportant
	destination := vertex{x: len(input[0]) - 1, y: len(input) - 1}
	if part == 2 {
		destination = vertex{x: len(input[0])*part2Multiplier - 1, y: len(input)*part2Multiplier - 1}
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			r, _ := strconv.Atoi(string(input[i][j]))
			riskLevels[vertex{x: j, y: i}] = r
			distances[vertex{x: j, y: i}] = math.MaxInt32
			if part == 2 { // copy input in x-direction >>>
				for k := 1; k < part2Multiplier; k++ {
					riskLevels[vertex{x: j + k*len(input[0]), y: i}] = getP2RiskLevel(r + k)
					distances[vertex{x: j + k*len(input[0]), y: i}] = math.MaxInt32
				}
			}
		}
		if part == 2 { // copy input in y-direction vvv
			for j := 0; j < len(input[i])*part2Multiplier; j++ {
				r := riskLevels[vertex{x: j, y: i}]
				for k := 1; k < part2Multiplier; k++ {
					riskLevels[vertex{x: j, y: i + k*len(input)}] = getP2RiskLevel(r + k)
					distances[vertex{x: j, y: i + k*len(input)}] = math.MaxInt32
				}
			}
		}
	}

	distances[vertex{x: 0, y: 0}] = 0
	next := vertex{x: 0, y: 0}
	done := false
	for !done {
		if _, exists := distances[vertex{x: next.x, y: next.y + 1}]; exists {
			if distances[next]+riskLevels[vertex{x: next.x, y: next.y + 1}] < distances[vertex{x: next.x, y: next.y + 1}] {
				distances[vertex{x: next.x, y: next.y + 1}] = distances[next] + riskLevels[vertex{x: next.x, y: next.y + 1}]
				nextLoc[vertex{x: next.x, y: next.y + 1}] = false
			}
		}
		if _, exists := distances[vertex{x: next.x, y: next.y - 1}]; exists {
			if distances[next]+riskLevels[vertex{x: next.x, y: next.y - 1}] < distances[vertex{x: next.x, y: next.y - 1}] {
				distances[vertex{x: next.x, y: next.y - 1}] = distances[next] + riskLevels[vertex{x: next.x, y: next.y - 1}]
				nextLoc[vertex{x: next.x, y: next.y - 1}] = false
			}
		}
		if _, exists := distances[vertex{x: next.x + 1, y: next.y}]; exists {
			if distances[next]+riskLevels[vertex{x: next.x + 1, y: next.y}] < distances[vertex{x: next.x + 1, y: next.y}] {
				distances[vertex{x: next.x + 1, y: next.y}] = distances[next] + riskLevels[vertex{x: next.x + 1, y: next.y}]
				nextLoc[vertex{x: next.x + 1, y: next.y}] = false
			}
		}
		if _, exists := distances[vertex{x: next.x - 1, y: next.y}]; exists {
			if distances[next]+riskLevels[vertex{x: next.x - 1, y: next.y}] < distances[vertex{x: next.x - 1, y: next.y}] {
				distances[vertex{x: next.x - 1, y: next.y}] = distances[next] + riskLevels[vertex{x: next.x - 1, y: next.y}]
				nextLoc[vertex{x: next.x - 1, y: next.y}] = false
			}
		}

		var min = math.MaxInt32
		for key := range nextLoc {
			if distances[key] < min {
				min = distances[key]
				next = key
			}
		}
		if next == destination {
			done = true
		}
		delete(nextLoc, next)
	}

	fmt.Printf("Part %d: %d", part, distances[destination])
}

func getP2RiskLevel(in int) int {
	for in > 9 {
		in -= 9
	}
	return in
}
