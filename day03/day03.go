package main

import (
	"AoC2021/helpers"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Length of the number in the input
const length = 12

func main() {
	input, err := helpers.ReadLinesToString("input.txt")
	if err != nil {
		fmt.Printf("Input read err: %s", err.Error())
		return
	}

	var gammaRate = make([]int, length)
	var epsilonRate = make([]int, length)
	for _, val := range input {
		temp := []rune(val)
		for i := 0; i < len(temp); i++ {
			if temp[i] == '1' {
				gammaRate[i]++
				epsilonRate[i]--
			} else {
				gammaRate[i]--
				epsilonRate[i]++
			}
		}
	}

	// Now lets get each binary string by the final values in the integer slices
	// There's no need for both checks since these are treating in an opposite fashion
	// but I'm leaving them separate in case one is treating uniquely in part 2.
	// Directions don't say how ties should be handled so I'll assume it should never
	// happen and is a cause for exiting.
	var gammaBinary, epsilonBinary string
	for i := 0; i < length; i++ {
		if gammaRate[i] > 0 {
			gammaBinary = gammaBinary + "1"
		} else if gammaRate[i] < 0 {
			gammaBinary = gammaBinary + "0"
		} else {
			fmt.Println("There's a tie for our input values count! Unsure how to handle...")
			os.Exit(1)
		}
		if epsilonRate[i] > 0 {
			epsilonBinary = epsilonBinary + "1"
		} else if epsilonRate[i] < 0 {
			epsilonBinary = epsilonBinary + "0"
		} else {
			fmt.Println("There's a tie for our input values count! Unsure how to handle...")
			os.Exit(1)
		}
	}

	// Convert the binary strings to int and print the product
	gamma, err := strconv.ParseInt(gammaBinary, 2, 64)
	if err != nil {
		fmt.Println("Could not convert gamma binary string to decimal")
	}
	epsilon, err := strconv.ParseInt(epsilonBinary, 2, 64)
	if err != nil {
		fmt.Println("Could not convert epsilon binary string to decimal")
	}
	fmt.Printf("Part 1: %d\n", gamma*epsilon)

	// Part 2, get the new binary strings using the previous answers and print the product
	oxygenBinary, err := getMatchingString(input, '1', true)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	carbonBinary, err := getMatchingString(input, '0', false)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	oxygen, err := strconv.ParseInt(oxygenBinary, 2, 64)
	if err != nil {
		fmt.Println("Could not convert oxygen binary string to decimal")
	}
	carbon, err := strconv.ParseInt(carbonBinary, 2, 64)
	if err != nil {
		fmt.Println("Could not convert carbon binary string to decimal")
	}
	fmt.Printf("Part 2: %d\n", oxygen*carbon)
}

// getMatchingString will return the sole binary string of input which matches each character belonging to the
// highest or lowest count of remaining strings with that character. For example, if the majority of remaining
// strings have a '1' as the first character and keepHigherCount is true, then all strings without a '1' as
// the first character will be pruned from the checked list. See the adventofCode Day 3 for more details.
func getMatchingString(input []string, tieBreaker byte, keepHigherCount bool) (string, error) {
	var remainingValues = make([]string, len(input))
	_ = copy(remainingValues, input)
	for i := 0; i < length; i++ {
		if len(remainingValues) == 2 {
			if remainingValues[0][i] == tieBreaker {
				return remainingValues[0], nil
			}
			return remainingValues[1], nil
		}
		var zeroCount = 0
		var oneCount = 0
		// We have to recount each time to find the greatest count of the remaining values because our input keeps changing.
		for j := 0; j < len(remainingValues); j++ {
			if remainingValues[j][i] == '0' {
				zeroCount++
			} else {
				oneCount++
			}
		}
		// When there's a tie in counts, assume the one count is bigger when we want to keep the higher counts
		// (for oxygen value) and also assume the one count is bigger when we want to keep the lower counts (for carbon).
		// We leave the keepHigherCount boolean in charge of dictating when we want to keep higher or lower values
		if zeroCount == oneCount {
			oneCount++
		}
		tempValues := make([]string, 0)
		for j := 0; j < len(remainingValues); j++ {
			if zeroCount > oneCount {
				if remainingValues[j][i] == '0' && keepHigherCount {
					tempValues = append(tempValues, remainingValues[j])
				}
				if remainingValues[j][i] == '1' && !keepHigherCount {
					tempValues = append(tempValues, remainingValues[j])
				}
			} else {
				if remainingValues[j][i] == '1' && keepHigherCount {
					tempValues = append(tempValues, remainingValues[j])
				}
				if remainingValues[j][i] == '0' && !keepHigherCount {
					tempValues = append(tempValues, remainingValues[j])
				}
			}
		}
		remainingValues = tempValues
		if len(remainingValues) == 1 {
			return remainingValues[0], nil
		}
	}
	return "", errors.New("parsed all values and found no winner, check input")
}
