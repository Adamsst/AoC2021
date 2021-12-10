package main

import (
	"AoC2021/helpers"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	chunks := map[string]string{
		"{": "}",
		"[": "]",
		"<": ">",
		"(": ")",
	}
	illegalScores := map[string]int{
		"}": 1197,
		"]": 57,
		">": 25137,
		")": 3,
	}
	p2Chunks := map[string]string{
		"}": "{",
		"]": "[",
		">": "<",
		")": "(",
	}

	input, err := helpers.ReadLinesToString("input.txt")
	if err != nil {
		fmt.Printf("Input read err: %s", err.Error())
		os.Exit(1)
	}

	var p1Score = 0
	p2Input := make([]string, 0)
	for i := 0; i < len(input); i++ {
		stack := make([]string, 0)
		for j := 0; j < len(input[i]); j++ {
			curElement := fmt.Sprintf("%c", input[i][j])
			if len(stack) == 0 { // Corruption doesnt occur by closing without opens, per instruction
				stack = append(stack, curElement) // So we can safely push next element whenever stack is empty
				if j == len(input[i])-1 {
					p2Input = append(p2Input, helpers.ReverseString(input[i]))
				}
				continue
			}
			if _, exists := chunks[curElement]; exists {
				stack = append(stack, curElement)
				if j == len(input[i])-1 {
					p2Input = append(p2Input, helpers.ReverseString(input[i]))
				}
				continue
			}
			topElement := stack[len(stack)-1]
			if curElement == chunks[topElement] {
				stack = stack[:len(stack)-1]
				if j == len(input[i])-1 {
					p2Input = append(p2Input, helpers.ReverseString(input[i]))
				}
				continue
			}
			p1Score += illegalScores[curElement]
			break
		}

	}
	fmt.Printf("Part 1: %d\n", p1Score)

	p2 := make([]int, 0)
	for i := 0; i < len(p2Input); i++ {
		stack := make([]string, 0)
		for j := 0; j < len(p2Input[i]); j++ {
			curElement := fmt.Sprintf("%c", p2Input[i][j])
			if len(stack) == 0 { // Corruption doesnt occur by closing without opens, per instruction
				stack = append(stack, curElement) // So we can safely push next element whenever stack is empty
				continue
			}
			if _, exists := p2Chunks[curElement]; exists {
				stack = append(stack, curElement)
				continue
			}
			topElement := stack[len(stack)-1]
			if curElement == p2Chunks[topElement] {
				stack = stack[:len(stack)-1]
				continue
			}
			stack = append(stack, curElement)
		}
		p2 = append(p2, p2Score(stack))
	}
	sort.Ints(p2)
	fmt.Printf("Part 2: %d", p2[int(math.Floor(float64(len(p2))/2))]) // p2[len(p2)/2] also works due to constant expression rules
}

func p2Score(input []string) int {
	p2Scores := map[string]int{
		"{": 3,
		"[": 2,
		"<": 4,
		"(": 1,
	}
	score := 0
	for i := 0; i < len(input); i++ {
		score *= 5
		score += p2Scores[input[i]]
	}
	return score
}
