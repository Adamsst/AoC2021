package main

import (
	"AoC2021/helpers"
	"fmt"
	"os"
)

const days int = 256

func main() {

	input, err := helpers.ReadLinesToString("input.txt")
	if err != nil {
		fmt.Printf("Input read err: %s", err.Error())
		os.Exit(1)
	}
	lanternFish, err := helpers.SplitStringToIntSlice(input[0], ",")
	if err != nil {
		fmt.Printf("Input parse err: %s", err.Error())
		os.Exit(1)
	}

	var day = 0
	lanternFishOld := make(map[int]int, 9) // Used to reference the previous iteration of lanternFish
	lanternFishNew := make(map[int]int, 9) // Holds the latest lanternFish values after spawns
	for _, val := range lanternFish {
		lanternFishOld[val]++
	}

	for day < days {
		for i := 8; i >= 1; i-- {
			lanternFishNew[i-1] = lanternFishOld[i]
		}
		lanternFishNew[6] += lanternFishOld[0]
		lanternFishNew[8] = lanternFishOld[0]
		day++
		for key, value := range lanternFishNew {
			lanternFishOld[key] = value
		}
		if day == 80 {
			fmt.Printf("Part 1: %d\n", getFishCount(lanternFishOld))
		}
	}
	fmt.Printf("Part 2: %d\n", getFishCount(lanternFishOld))
}

func getFishCount(fish map[int]int) (count int) {
	for _, v := range fish {
		count += v
	}
	return count
}
