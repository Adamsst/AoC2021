package main

import (
	"AoC2021/helpers"
	"fmt"
	"os"
	"strings"
)

const caveSize = 2

func main() {

	input, err := helpers.ReadLinesToString("input.txt")
	if err != nil {
		fmt.Printf("Input read err: %s", err.Error())
		os.Exit(1)
	}

	newInput := make([]string, 0)
	for i := 0; i < len(input); i++ {
		newInput = append(newInput, input[i])
		splitInput := strings.Split(input[i], "-")
		newInput = append(newInput, splitInput[1]+"-"+splitInput[0])
	}

	paths := make([]string, 0)
	for i := 0; i < len(newInput); i++ {
		if strings.Split(newInput[i], "-")[0] == "start" {
			paths = append(paths, newInput[i])
		}
	}
	var continued = true
	for continued {
		continued = false
		tempPaths := make([]string, 0)
		for i := 0; i < len(paths); i++ {
			pathCaves := strings.Split(paths[i], "-")
			curCave := pathCaves[len(pathCaves)-1]
			if paths[i][len(paths[i])-3:] == "end" {
				tempPaths = append(tempPaths, paths[i])
				continue
			}

			for j := 0; j < len(newInput); j++ {
				adjoinCave := strings.Split(newInput[j], "-")[0]
				exitCave := strings.Split(newInput[j], "-")[1]

				if helpers.StringIsLower(adjoinCave) {
					if strings.Contains(paths[i][:len(paths[i])-caveSize], adjoinCave) { //adjoinCave already present in path. can't revisit.
						continue
					}
				}
				if helpers.StringIsLower(exitCave) {
					if strings.Contains(paths[i], exitCave) { //exitCave already present in path. can't revisit.
						continue
					}
				}

				if adjoinCave == curCave {
					continued = true
					tempPaths = append(tempPaths, paths[i]+newInput[j][caveSize:])
				}
			}
		}
		paths = make([]string, len(tempPaths))
		copy(paths, tempPaths)
	}
	fmt.Printf("Part 1: %d\n", len(paths))

	// Part 2 is quite similar to part 1, but I want both scores available so I copy/pasted the analyzing logic
	paths = make([]string, 0)
	for i := 0; i < len(newInput); i++ {
		if strings.Split(newInput[i], "-")[0] == "start" {
			paths = append(paths, newInput[i])
		}
	}
	continued = true
	for continued {
		continued = false
		tempPaths := make([]string, 0)
		for i := 0; i < len(paths); i++ {
			pathCaves := strings.Split(paths[i], "-")
			curCave := pathCaves[len(pathCaves)-1]
			if paths[i][len(paths[i])-3:] == "end" {
				tempPaths = append(tempPaths, paths[i])
				continue
			}

			for j := 0; j < len(newInput); j++ {
				adjoinCave := strings.Split(newInput[j], "-")[0]
				exitCave := strings.Split(newInput[j], "-")[1]
				if adjoinCave == "start" || adjoinCave == "end" || exitCave == "start" {
					continue
				}

				if helpers.StringIsLower(adjoinCave) {
					if strings.Contains(paths[i][:len(paths[i])-caveSize], adjoinCave) {
						if PathHasRevisitedSmallCave(paths[i][:len(paths[i])-caveSize]) {
							continue // adjoiningCave is small and path has already visited a small cave more than once. Can't revisit.
						}
					}
				}
				if helpers.StringIsLower(exitCave) {
					if strings.Contains(paths[i], exitCave) {
						if PathHasRevisitedSmallCave(paths[i]) {
							continue // exitCave is small and path has already visited a small cave more than once. Can't revisit.
						}
					}
				}

				if adjoinCave == curCave {
					continued = true
					tempPaths = append(tempPaths, paths[i]+"-"+newInput[j])
				}
			}
		}
		paths = make([]string, len(tempPaths))
		copy(paths, tempPaths)
	}
	fmt.Printf("Part 2: %d", len(paths))
}

func PathHasRevisitedSmallCave(path string) bool {
	caves := strings.Split(path, "-")
	for i := 0; i < len(caves); i++ {
		if !helpers.StringIsLower(caves[i]) {
			continue
		}
		if strings.Count(path, caves[i]) > 3 {
			return true
		}
	}
	return false
}
