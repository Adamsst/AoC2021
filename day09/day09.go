package main

import (
	"AoC2021/helpers"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type point struct {
	x, y int
}

func main() {

	input, err := helpers.ReadLinesToString("input.txt")
	if err != nil {
		fmt.Printf("Input read err: %s", err.Error())
		os.Exit(1)
	}

	smokeHeight := make(map[point]int, len(input)*len(input[0]))
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			smokeHeight[point{j, i}], _ = strconv.Atoi(fmt.Sprintf("%c", input[i][j]))
		}
	}

	var p1Risk = 0
	lowPoints := make(map[point]int)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			var isLowPoint = true
			for k := -1; k <= 1; k++ {
				if val, exists := smokeHeight[point{j + k, i - 1}]; exists {
					if val <= smokeHeight[point{j, i}] {
						isLowPoint = false
					}
				}
				if val, exists := smokeHeight[point{j + k, i + 1}]; exists {
					if val <= smokeHeight[point{j, i}] {
						isLowPoint = false
					}
				}
			}
			if val, exists := smokeHeight[point{j - 1, i}]; exists {
				if val <= smokeHeight[point{j, i}] {
					isLowPoint = false
				}
			}
			if val, exists := smokeHeight[point{j + 1, i}]; exists {
				if val <= smokeHeight[point{j, i}] {
					isLowPoint = false
				}
			}
			if isLowPoint {
				lowPoints[point{j, i}] = smokeHeight[point{j, i}]
				p1Risk += 1 + smokeHeight[point{j, i}]
			}
		}
	}
	fmt.Printf("Part 1: %d\n", p1Risk)

	biggestScores := []int{0, 0, 0}
	for k := range lowPoints {
		var pointsToCheck = make([]point, 0)
		var checkedPoints = make(map[point]int)
		pointsToCheck = append(pointsToCheck, k)
		for i := 0; i < len(pointsToCheck); i++ {
			if _, exists := checkedPoints[pointsToCheck[i]]; exists {
				continue
			}
			if val, exists := smokeHeight[point{pointsToCheck[i].x + 1, pointsToCheck[i].y}]; exists {
				if val > smokeHeight[point{pointsToCheck[i].x, pointsToCheck[i].y}] && val != 9 {
					pointsToCheck = append(pointsToCheck, point{pointsToCheck[i].x + 1, pointsToCheck[i].y})
				}
			}
			if val, exists := smokeHeight[point{pointsToCheck[i].x - 1, pointsToCheck[i].y}]; exists {
				if val > smokeHeight[point{pointsToCheck[i].x, pointsToCheck[i].y}] && val != 9 {
					pointsToCheck = append(pointsToCheck, point{pointsToCheck[i].x - 1, pointsToCheck[i].y})
				}
			}
			if val, exists := smokeHeight[point{pointsToCheck[i].x, pointsToCheck[i].y - 1}]; exists {
				if val > smokeHeight[point{pointsToCheck[i].x, pointsToCheck[i].y}] && val != 9 {
					pointsToCheck = append(pointsToCheck, point{pointsToCheck[i].x, pointsToCheck[i].y - 1})
				}
			}
			if val, exists := smokeHeight[point{pointsToCheck[i].x, pointsToCheck[i].y + 1}]; exists {
				if val > smokeHeight[point{pointsToCheck[i].x, pointsToCheck[i].y}] && val != 9 {
					pointsToCheck = append(pointsToCheck, point{pointsToCheck[i].x, pointsToCheck[i].y + 1})
				}
			}
			checkedPoints[pointsToCheck[i]] = 1
		}
		sort.Ints(biggestScores)
		for j := 0; j <= 2; j++ {
			if len(checkedPoints) >= biggestScores[j] {
				biggestScores[j] = len(checkedPoints)
				break
			}
		}
	}
	fmt.Printf("Part 2: %d", biggestScores[0]*biggestScores[1]*biggestScores[2])
}
