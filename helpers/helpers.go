package helpers

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type sortRuneString []rune

// readLinesToString reads lines of a file into a slice of strings
func ReadLinesToString(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// readLinesToInt reads lines of a file into a slice of integers
func ReadLinesToInt(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("err converting to int: " + scanner.Text())
			continue
		}
		lines = append(lines, val)
	}
	return lines, scanner.Err()
}

// SplitStringToIntSlice splits the input string on the given separator and attempts to convert
// each element of the resulting slice to an integer and return the result as a slice of integers.
// If any conversion fails an err will be returned alongside a nil result.
func SplitStringToIntSlice(input string, separator string) ([]int, error) {
	var strInput = strings.Split(input, separator)
	var intResult = make([]int, len(strInput))
	for i, val := range strInput {
		var intVal, err = strconv.Atoi(strings.Trim(val, " "))
		if err != nil {
			return nil, fmt.Errorf("SplitStringToIntSlice err converting to int: %s, err: %s", val, err.Error())
		}
		intResult[i] = intVal
	}
	return intResult, nil
}

// SortString is used to sort the character in a string and return the result
// The following was taken from https://golangbyexample.com/sort-string-golang/
// This example implements the same interface as the Go library in order to
// leverage the provided sort function on a rune slice of the input string
func SortString(input string) string {
	runeArray := []rune(input)
	sort.Sort(sortRuneString(runeArray))
	return string(runeArray)
}

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}

// StringContainsChars will return true if the target string contains all chars
// in the provided char string. Number of occurrences does not matter.
func StringContainsChars(target string, chars string) bool {
	for i := 0; i < len(chars); i++ {
		if !strings.Contains(target, string(chars[i])) {
			return false
		}
	}
	return true
}

// ReverseString will return the input string backwards. Abc is returned as cbA.
// The following was taken from https://golangbyexample.com/reverse-a-string-in-golang/
func ReverseString(str string) string {
	if len(str) == 0 {
		return str
	}
	strRunes := []rune(str)
	var result []rune
	for i := len(strRunes) - 1; i >= 0; i-- {
		result = append(result, strRunes[i])
	}
	return string(result)
}
