package main

import (
	"AoC2021/helpers"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func main() {

	var maxX = 0
	var maxY = 0
	folds := make([]string, 0)

	input, err := helpers.ReadLinesToString("input.txt")
	if err != nil {
		fmt.Printf("Input read err: %s", err.Error())
		os.Exit(1)
	}

	var paper = make(map[point]int)
	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			continue
		}
		if input[i][0:2] == "fo" {
			folds = append(folds, strings.Split(input[i], " ")[2])
		} else {
			var temp = strings.Split(input[i], ",")
			x, err := strconv.Atoi(temp[0])
			if err != nil {
				fmt.Printf("Input parse err: %s", err.Error())
				os.Exit(1)
			}
			if x > maxX {
				maxX = x
			}
			y, err := strconv.Atoi(temp[1])
			if err != nil {
				fmt.Printf("Input parse err: %s", err.Error())
				os.Exit(1)
			}
			if y > maxY {
				maxY = y
			}
			paper[point{x, y}] = 8 // Occupied points initialized to 8
		}
	}
	fmt.Printf("x:%d y:%d\n", maxX, maxY)

	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			if _, exists := paper[point{x, y}]; !exists {
				paper[point{x, y}] = 0 // Empty points initialized to 8
			}
		}
	}

	for i := 0; i < len(folds); i++ {
		temp := strings.Split(folds[i], "=")
		foldNum, _ := strconv.Atoi(temp[1])
		var num = foldNum + 1
		if strings.Contains(folds[i], "y=") {
			for y := num; y <= maxY; y++ {
				for x := 0; x <= maxX; x++ {
					if paper[point{x: x, y: y}] == 8 {
						paper[point{x: x, y: (y - ((y - foldNum) * 2))}] = 8
					}
				}
			}
			maxY = foldNum - 1
		} else {
			for x := num; x <= maxX; x++ {
				for y := 0; y <= maxY; y++ {
					if paper[point{x: x, y: y}] == 8 {
						paper[point{x: (x - ((x - foldNum) * 2)), y: y}] = 8
					}
				}
			}
			maxX = foldNum - 1
		}
		if i == 0 {
			var p1Score = 0
			for y := 0; y <= maxY; y++ {
				for x := 0; x <= maxX; x++ {
					if paper[point{x: x, y: y}] == 8 {
						p1Score++
					}
				}
			}
			fmt.Printf("Part 1: %d\n", p1Score)
		}
	}

	fmt.Println("Part 2:")
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			var s = " "
			if paper[point{x: x, y: y}] == 8 {
				s = "█"
			}
			fmt.Print(s)
		}
		fmt.Println()
	}
}
