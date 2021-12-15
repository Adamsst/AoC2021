package main

import (
	"AoC2021/helpers"
	"fmt"
	"math"
	"os"
	"strings"
)

const p1Steps = 10
const p2Steps = 40

func main() {

	input, err := helpers.ReadLinesToString("input.txt")
	if err != nil {
		fmt.Printf("Input read err: %s", err.Error())
		os.Exit(1)
	}

	var template = input[0]
	pairs := make(map[string]string, len(input)-2)
	p2Pairs := make(map[string]string)
	for i := 2; i < len(input); i++ {
		temp := strings.Split(input[i], " ")
		pairs[temp[0]] = temp[2]
		p2Pairs[temp[0]] = string(temp[0][0]) + temp[2] + temp[2] + string(temp[0][1])
	}

	for i := 0; i < p1Steps; i++ {
		var result = ""
		for j := 0; j < len(template)-1; j++ {
			result = result + string(template[j]) + pairs[template[j:j+2]]
		}
		result = result + string(template[len(template)-1])
		template = result
	}
	fmt.Printf("Part 1: %d\n", getP1Score(template))

	p2Scores := make(map[string]int)
	template = input[0]
	for i := 0; i < len(template)-1; i++ {
		p2Scores[template[i:i+2]]++
	}
	addOne := template[len(template)-1]

	for i := 0; i < p2Steps; i++ {
		p2New := make(map[string]int)
		for key, val := range p2Scores {
			if val > 0 {
				p2New[p2Pairs[key][0:2]] += val
				p2New[p2Pairs[key][2:4]] += val
			}
		}
		for k := range p2Scores {
			p2Scores[k] = 0
		}
		for k, v := range p2New {
			p2Scores[k] = v
		}
	}

	fmt.Printf("Part 2: %d", getP2Score(p2Scores, addOne))
}

func getP1Score(input string) int {
	scores := make(map[byte]int)
	for i := 0; i < len(input); i++ {
		if _, exists := scores[input[i]]; !exists {
			scores[input[i]] = 1
		} else {
			scores[input[i]]++
		}
	}
	var max = 0
	var min = len(input)
	for _, val := range scores {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	return max - min
}

func getP2Score(input map[string]int, addOne byte) int {
	scores := make(map[byte]int)
	for key, val := range input {
		if _, exists := scores[key[0]]; !exists {
			scores[key[0]] = val
		} else {
			scores[key[0]] += val
		}
	}
	scores[addOne]++
	var max = 0
	var min = math.MaxInt64
	for _, val := range scores {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	return max - min
}
