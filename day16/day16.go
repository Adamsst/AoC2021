package main

import (
	"AoC2021/helpers"
	"fmt"
	"os"
)

type packet struct {
	typeID                int
	version               int
	value                 int
	lengthTypeID          string
	totalLengthSubPackets int
	numberOfSubPackets    int
	endIndex              int
	isSubPacket           bool
}

func main() {
	hexToBin := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}

	input, err := helpers.ReadLinesToString("input.txt")
	if err != nil {
		fmt.Printf("Input read err: %s", err.Error())
		os.Exit(1)
	}

	var binInput = ""
	for i := 0; i < len(input[0]); i++ {
		binInput += hexToBin[string(input[0][i])]
	}

	packets := make([]packet, 0)
	var p1Version = 0
	var endIndex = 0
	var i = 0
	var lv = 0  // int to hold literal values
	var sl = 0  // int to hold sub lengths
	var lsp = 0 // int to length of sub packets
	var version = 0
	var typeID = 0
	var subLengthID = ""
	for i < len(binInput) {
		if !helpers.StringContainsChars(binInput[i:], "1") {
			break
		}
		version, i = getResultAndIndex(binInput, i, 3)
		typeID, i = getResultAndIndex(binInput, i, 3)
		switch typeID {
		case 4:
			lv, i = getLiteralValue(binInput, i)
		default:
			subLengthID = string(binInput[i])
			sl, i, _ = getSubLengthByBit(binInput, i)
			lsp, i = getResultAndIndex(binInput, i, sl)
		}
		endIndex = i - 1

		p := packet{
			typeID:                typeID,
			version:               version,
			value:                 0,
			lengthTypeID:          subLengthID,
			totalLengthSubPackets: 0,
			numberOfSubPackets:    0,
			endIndex:              endIndex,
			isSubPacket:           false,
		}
		if typeID == 4 {
			p.value = lv
		}
		if subLengthID == "0" {
			p.totalLengthSubPackets = lsp
		} else {
			p.numberOfSubPackets = lsp
		}
		packets = append(packets, p)
		p1Version += p.version
	}
	fmt.Printf("Part 1: %d\n", p1Version)

	for i = len(packets) - 1; i >= 0; i-- {
		if packets[i].typeID == 4 {
			continue
		}
		valuesToProcess := make([]int, 0)
		if packets[i].lengthTypeID == "0" { // Process sub packets by length
			j := i + 1
			for packets[j].endIndex <= packets[i].endIndex+packets[i].totalLengthSubPackets {
				if !packets[j].isSubPacket {
					packets[j].isSubPacket = true
					valuesToProcess = append(valuesToProcess, packets[j].value)
				}
				if packets[j].endIndex == packets[i].endIndex+packets[i].totalLengthSubPackets {
					break
				}
				j++
			}
		} else { // Process sub packets by count
			j := packets[i].numberOfSubPackets
			k := 0
			index := 1
			for k < j {
				if !packets[i+index].isSubPacket {
					packets[i+index].isSubPacket = true
					valuesToProcess = append(valuesToProcess, packets[i+index].value)
					k++
				}
				index++
			}
		}

		val := 0
		switch packets[i].typeID {
		case 0:
			val = helpers.GetSumInts(valuesToProcess)
		case 1:
			val, err = helpers.GetProductInts(valuesToProcess)
			if err != nil {
				fmt.Println("Case 1 err: " + err.Error())
			}
		case 2:
			val, err = helpers.GetMinInt(valuesToProcess)
			if err != nil {
				fmt.Println("Case 2 err: " + err.Error())
			}
		case 3:
			val, err = helpers.GetMaxInt(valuesToProcess)
			if err != nil {
				fmt.Println("Case 3 err: " + err.Error())
			}
		case 5:
			if valuesToProcess[0] > valuesToProcess[1] {
				val = 1
			}
		case 6:
			if valuesToProcess[0] < valuesToProcess[1] {
				val = 1
			}
		case 7:
			if valuesToProcess[0] == valuesToProcess[1] {
				val = 1
			}
		default:
			fmt.Println("Default case! Yikes!")
		}
		packets[i].value = val
	}
	fmt.Printf("Part 2: %d", packets[0].value)
}

func getResultAndIndex(str string, in, length int) (int, int) {
	res64, err := helpers.GetDecFromBinStr(str[in : in+length])
	if err != nil {
		fmt.Printf("getResultAndIndex: %s: err:%s\n", str, err.Error())
		return 0, 0
	}
	return int(res64), in + length
}

func getLiteralValue(str string, in int) (int, int) {
	var final = false
	var result = ""
	for !final {
		next := str[in : in+5]
		if string(next[0]) == "0" {
			final = true
		}
		result += next[1:]
		in += 5
	}
	res64, err := helpers.GetDecFromBinStr(result)
	if err != nil {
		fmt.Printf("getLiteralValue: %s: err:%s\n", str, err.Error())
		return 0, 0
	}
	return int(res64), in
}

func getSubLengthByBit(str string, in int) (int, int, bool) {
	if string(str[in]) == "0" {
		return 15, in + 1, true
	}
	return 11, in + 1, false
}
