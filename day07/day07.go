package main

import (
	"AoC2021/helpers"
	"fmt"
	"math"
	"os"
)

func main() {

	input, err := helpers.ReadLinesToString("input.txt")
	if err != nil {
		fmt.Printf("Input read err: %s", err.Error())
		os.Exit(1)
	}
	crabPositions, err := helpers.SplitStringToIntSlice(input[0], ",")
	if err != nil {
		fmt.Printf("Input parse err: %s", err.Error())
		os.Exit(1)
	}

	var p1Sum = math.MaxFloat64
	var p2Sum = math.MaxFloat64
	for key := range crabPositions {
		var tempSum = 0.0
		var tempSum2 = 0.0
		for i := 0; i < len(crabPositions); i++ {
			tempSum += math.Abs(float64(crabPositions[key] - crabPositions[i]))
			tempSum2 += getCrabSum(math.Abs(float64(crabPositions[key] - crabPositions[i])))
		}
		if tempSum < p1Sum {
			p1Sum = tempSum
		}
		if tempSum2 < p2Sum {
			p2Sum = tempSum2
		}
	}
	fmt.Printf("Part 1: %g\nPart 2: %.0f", p1Sum, p2Sum)
}

func getCrabSum(input float64) float64 {
	return (input) * (input + 1) / 2
}
