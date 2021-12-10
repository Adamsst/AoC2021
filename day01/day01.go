package main

import (
	"AoC2021/helpers"
	"fmt"
)

func main() {
	input, err := helpers.ReadLinesToInt("input.txt")
	if err != nil {
		fmt.Printf("input read err: %s", err.Error())
		return
	}

	var p1Score = 0
	var previous = input[0]
	for _, val := range input {
		if val > previous {
			p1Score++
		}
		previous = val
	}
	fmt.Printf("Part 1: %d\n", p1Score)

	var p2Score = 0
	previous = input[0] + input[1] + input[2]
	for i := 3; i < len(input); i++ {
		if input[i]+input[i-1]+input[i-2] > previous {
			p2Score++
		}
		previous = previous - input[i-3] + input[i]
	}
	fmt.Printf("Part 2: %d", p2Score)
}
