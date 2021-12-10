package main

import (
	"AoC2021/helpers"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type line struct {
	x1, x2, y1, y2 int
}

type point struct {
	x, y int
}

const maxInt = 1000

func main() {

	input, err := helpers.ReadLinesToString("input.txt")
	if err != nil {
		fmt.Printf("Input read err: %s", err.Error())
		os.Exit(1)
	}

	var lines = make([]*line, len(input))
	for i, val := range input {
		lines[i], err = getLineFromInput(val)
		if err != nil {
			fmt.Printf("Err parsing input: %s, %s", err, val)
			os.Exit(1)
		}
	}

	points := make(map[point]int, maxInt*maxInt)
	for _, val := range lines {
		visited := getPointsFromLine(*val, false)
		for i := 0; i < len(visited); i++ {
			points[visited[i]]++
		}
	}

	var p1 = 0
	for _, v := range points {
		if v >= 2 {
			p1++
		}
	}
	fmt.Printf("Part 1: %d\n", p1)

	points = make(map[point]int, maxInt*maxInt)
	for _, val := range lines {
		visited := getPointsFromLine(*val, true)
		for i := 0; i < len(visited); i++ {
			points[visited[i]]++
		}
	}

	var p2 = 0
	for _, v := range points {
		if v >= 2 {
			p2++
		}
	}
	fmt.Printf("Part 2: %d", p2)
}

func getLineFromInput(input string) (*line, error) {
	temp := strings.Split(input, " ")
	firstStringCoords := strings.Split(temp[0], ",")
	secondStringCoords := strings.Split(temp[2], ",")

	x1, err := strconv.Atoi(firstStringCoords[0])
	if err != nil {
		return nil, err
	}
	y1, err := strconv.Atoi(firstStringCoords[1])
	if err != nil {
		return nil, err
	}
	x2, err := strconv.Atoi(secondStringCoords[0])
	if err != nil {
		return nil, err
	}
	y2, err := strconv.Atoi(secondStringCoords[1])
	if err != nil {
		return nil, err
	}
	return &line{
		x1: x1,
		x2: x2,
		y1: y1,
		y2: y2,
	}, nil
}

func getPointsFromLine(input line, part2 bool) (result []point) {
	var i int
	var max int
	if input.x1 == input.x2 {
		if input.y1 < input.y2 {
			i = input.y1
			max = input.y2
		} else {
			i = input.y2
			max = input.y1
		}
		for i <= max {
			result = append(result, point{input.x1, i})
			i++
		}
	} else if input.y1 == input.y2 {
		if input.x1 < input.x2 {
			i = input.x1
			max = input.x2
		} else {
			i = input.x2
			max = input.x1
		}
		for i <= max {
			result = append(result, point{i, input.y1})
			i++
		}
	} else if part2 {
		var yAdjust, xAdjust int
		if input.y1 < input.y2 {
			yAdjust = 1
		} else {
			yAdjust = -1
		}
		if input.x1 < input.x2 {
			xAdjust = 1
		} else {
			xAdjust = -1
		}
		for i <= int(math.Abs(float64(input.y1-input.y2))) {
			result = append(result, point{input.x1 + i*xAdjust, input.y1 + i*yAdjust})
			i++
		}
	}
	return result
}
