package main

import (
	"AoC2021/helpers"
	"fmt"
	"strconv"
	"strings"
)

type diveInstruction struct {
	direction string
	magnitude int
}

func main() {
	input, err := helpers.ReadLinesToString("input.txt")
	if err != nil {
		fmt.Printf("input read err: %s", err.Error())
		return
	}

	diveInstructions := make([]diveInstruction, len(input))
	for i, val := range input {
		temp := strings.Split(val, " ")
		mag, err := strconv.Atoi(temp[1])
		if err != nil {
			fmt.Println("found a non-int magnitude: ", err.Error())
		}
		diveInstructions[i] = diveInstruction{
			direction: temp[0],
			magnitude: mag,
		}
	}

	var x, y int
	for _, val := range diveInstructions {
		switch val.direction {
		case "forward":
			x += val.magnitude
		case "up":
			y -= val.magnitude
		case "down":
			y += val.magnitude
		default:
			fmt.Println("unknown direction found: " + val.direction)
		}
	}
	fmt.Printf("Part 1: %d\n", x*y)

	x = 0
	y = 0
	var a = 0
	for _, val := range diveInstructions {
		switch val.direction {
		case "forward":
			y += a * val.magnitude
			x += val.magnitude
		case "up":
			a -= val.magnitude
		case "down":
			a += val.magnitude
		default:
			fmt.Println("unknown direction found: " + val.direction)
		}
	}
	fmt.Printf("Part 2: %d\n", x*y)
}
