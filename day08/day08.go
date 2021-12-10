package main

import (
	"AoC2021/helpers"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := helpers.ReadLinesToString("input.txt")
	if err != nil {
		fmt.Printf("Input read err: %s", err.Error())
		os.Exit(1)
	}

	p1Count := 0
	output := make(map[int][]string, len(input))
	allInput := make(map[int][]string, len(input))
	for i := 0; i < len(input); i++ {
		temp := strings.Split(input[i], "|")
		output[i] = strings.Split(strings.Trim(temp[1], " "), " ")
		allInput[i] = append(strings.Split(strings.Trim(temp[0], " "), " "), strings.Split(strings.Trim(temp[1], " "), " ")...)
	}
	for _, val := range output {
		for i := 0; i < len(val); i++ {
			if len(val[i]) == 2 || len(val[i]) == 3 || len(val[i]) == 4 || len(val[i]) == 7 { // 1, 7, 4, 8
				p1Count++
			}
		}
	}
	fmt.Printf("Part 1: %d\n", p1Count)

	var p2Sum = 0
	// sort all of the strings, then get their four digit int
	for i := 0; i < len(allInput); i++ {
		for key, val := range allInput[i] {
			allInput[i][key] = helpers.SortString(val)
		}
		temp, err := getFourDigit(allInput[i])
		if err != nil {
			fmt.Println("err determining four digit output")
			os.Exit(1)
		}
		p2Sum += temp
	}
	fmt.Printf("Part 2: %d", p2Sum)
}

// there's a lot of optimizations to be done here, but I left it as is to show my thought process
func getFourDigit(inputStr []string) (int, error) {
	tenUniqueVals := inputStr[:10]
	fourDigits := inputStr[10:]
	stringToNum := make(map[int]string, 10)

	for i := 0; i < len(tenUniqueVals); i++ { // Assign the known easy ones: 1, 7, 4, 8
		if len(tenUniqueVals[i]) == 2 {
			stringToNum[1] = tenUniqueVals[i]
		} else if len(tenUniqueVals[i]) == 3 {
			stringToNum[7] = tenUniqueVals[i]
		} else if len(tenUniqueVals[i]) == 4 {
			stringToNum[4] = tenUniqueVals[i]
		} else if len(tenUniqueVals[i]) == 7 {
			stringToNum[8] = tenUniqueVals[i]
		}
	}
	for i := 0; i < len(tenUniqueVals); i++ { // Assign 3
		if len(tenUniqueVals[i]) == 5 && helpers.StringContainsChars(tenUniqueVals[i], stringToNum[1]) {
			stringToNum[3] = tenUniqueVals[i]
		}
	}
	for i := 0; i < len(tenUniqueVals); i++ { // Assign 6
		if len(tenUniqueVals[i]) == 6 && !helpers.StringContainsChars(tenUniqueVals[i], stringToNum[1]) {
			stringToNum[6] = tenUniqueVals[i]
		}
	}
	for i := 0; i < len(tenUniqueVals); i++ { // Assign 9
		if len(tenUniqueVals[i]) == 6 && helpers.StringContainsChars(tenUniqueVals[i], stringToNum[4]) {
			stringToNum[9] = tenUniqueVals[i]
		}
	}
	for i := 0; i < len(tenUniqueVals); i++ { // Assign 0
		if len(tenUniqueVals[i]) == 6 && tenUniqueVals[i] != stringToNum[9] && tenUniqueVals[i] != stringToNum[6] {
			stringToNum[0] = tenUniqueVals[i]
		}
	}
	for i := 0; i < len(tenUniqueVals); i++ { // Assign 2, 5
		if len(tenUniqueVals[i]) == 5 && tenUniqueVals[i] != stringToNum[3] {
			count := 0
			for j := 0; j < len(stringToNum[4]); j++ {
				if strings.Contains(tenUniqueVals[i], string(stringToNum[4][j])) {
					count++
				}
			}
			if count == 3 {
				stringToNum[5] = tenUniqueVals[i]
			} else {
				stringToNum[2] = tenUniqueVals[i]
			}
		}
	}

	var resultStr = ""
	for i := 0; i < len(fourDigits); i++ {
		for key, value := range stringToNum {
			if value == fourDigits[i] {
				resultStr += fmt.Sprintf("%d", key)
			}
		}
	}

	return strconv.Atoi(resultStr)
}
