package main

import (
	"AoC2021/helpers"
	"fmt"
	"os"
	"strconv"
)

type point struct {
	x, y int
}

const steps = 100

func main() {

	input, err := helpers.ReadLinesToString("input.txt")
	if err != nil {
		fmt.Printf("Input read err: %s", err.Error())
		os.Exit(1)
	}

	energyLevel := make(map[point]int, len(input)*len(input[0]))
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			energyLevel[point{j, i}], _ = strconv.Atoi(fmt.Sprintf("%c", input[i][j]))
		}
	}

	var step = 0
	var p1Sum = 0
	var p2AllFlash = false // Part 2 we have to print the step in which each enery level flashes
	for step < steps || !p2AllFlash {
		flashMap := make(map[point]int, len(input)*len(input[0]))
		for key := range energyLevel {
			energyLevel[key]++ // First increase the value of each energy level by 1
		}
		var flashers = true
		for flashers {
			flashers = false
			numFlashes := len(flashMap)
			for key, value := range energyLevel {
				if _, exists := flashMap[key]; !exists { // if flash map already contains the key, dont reset its value to 10
					if value == 10 {
						flashMap[key] = value
					}
				}
			}
			for key, value := range flashMap {
				if value < 11 {
					incSurrounding(key, energyLevel) // for each flashMap value increase surrounding
				}
				flashMap[key] = 11 // after processing a flash map key once, set a high value to prevent processing again
			}
			flashers = !(numFlashes == len(flashMap)) // repeat until theres no new flashes
		}
		p1Sum += len(flashMap)
		for key, value := range energyLevel {
			if value == 10 {
				energyLevel[key] = 0 // reset all the flashed energy levels
			}
		}
		step++
		if step == 100 {
			fmt.Printf("Part 1: %d\n", p1Sum)
		}
		if len(flashMap) == len(energyLevel) {
			p2AllFlash = true
			fmt.Printf("Part 2: %d", step)
		}
	}
}

func incSurrounding(p point, energy map[point]int) {
	for k := -1; k <= 1; k++ {
		if val, exists := energy[point{p.x + k, p.y - 1}]; exists {
			if val <= 9 {
				energy[point{p.x + k, p.y - 1}]++
			}
		}
		if val, exists := energy[point{p.x + k, p.y + 1}]; exists {
			if val <= 9 {
				energy[point{p.x + k, p.y + 1}]++
			}
		}
	}
	if val, exists := energy[point{p.x - 1, p.y}]; exists {
		if val <= 9 {
			energy[point{p.x - 1, p.y}]++
		}
	}
	if val, exists := energy[point{p.x + 1, p.y}]; exists {
		if val <= 9 {
			energy[point{p.x + 1, p.y}]++
		}
	}
}
