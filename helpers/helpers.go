package helpers

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
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
// leverage the provided sort function on a rune slice of the input string.
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

// StringIsLower will return true when all characters in the provided string are lowercase.
func StringIsLower(str string) bool {
	if len(str) == 0 {
		return false
	}
	for i := 0; i < len(str); i++ {
		if !unicode.IsLower(rune(str[i])) {
			return false
		}
	}
	return true
}

// GetDecFromBinStr expects the given input string to be binary and will convert it to its
// decimal respresentation and return the resulting int64. An err is returned on failure.
func GetDecFromBinStr(str string) (int64, error) {
	result, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// GetMaxInt returns the maximum integer value from a slice of ints
func GetMaxInt(in []int) (int, error) {
	var max = -1 * math.MaxInt64
	if len(in) == 0 {
		return 0, errors.New("GetMaxInt empty slice provided")
	}
	for i := 0; i < len(in); i++ {
		if in[i] > max {
			max = in[i]
		}
	}
	return max, nil
}

// GetMinInt returns the minimum integer value from a slice of ints
func GetMinInt(in []int) (int, error) {
	var min = math.MaxInt64
	if len(in) == 0 {
		return 0, errors.New("GetMinInt empty slice provided")
	}
	for i := 0; i < len(in); i++ {
		if in[i] < min {
			min = in[i]
		}
	}
	return min, nil
}

// GetSumInts returns the sum of a slice of ints
func GetSumInts(in []int) int {
	var sum = 0
	for i := 0; i < len(in); i++ {
		sum += in[i]
	}
	return sum
}

// GetProductInts returns the product of a slice of ints
// An err is return if the slice is empty
// The lone value is returned is the slice has a single element
func GetProductInts(in []int) (int, error) {
	var product = 1
	if len(in) < 1 {
		return 0, errors.New("GetProductInts no elements in provided input")
	}
	for i := 0; i < len(in); i++ {
		product *= in[i]
	}
	return product, nil
}
