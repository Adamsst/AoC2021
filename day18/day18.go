package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// struct to represent snailfish values
type sfVal struct {
	isSet bool
	value int
}

// struct to represent snailfish
type sf struct {
	// left right depth
	l, r   sfVal
	d, mag int
}

func main() {

	input, err := readLinesToString("input.txt")
	if err != nil {
		fmt.Println("err reading input file")
		os.Exit(1)
	}

	magStr := partOne(input)
	p1Slice, err := scanToSlice(magStr)
	if err != nil {
		fmt.Println("couldn't get p1 slice")
		os.Exit(1)
	}

	fmt.Printf("Part 1: %d\n", getMagnitude(p1Slice))

	var p2Mag = 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if i == j {
				continue
			}
			magStr = partOne([]string{input[i], input[j]})
			p2Slice, err := scanToSlice(magStr)
			if err != nil {
				fmt.Println("couldn't get p12 slice")
				os.Exit(1)
			}
			if tempMag := getMagnitude(p2Slice); tempMag > p2Mag {
				p2Mag = tempMag
			}
		}
	}
	fmt.Printf("Part 2: %d", p2Mag)
}

// readLinesToString reads lines of a file into a slice of strings.
func readLinesToString(path string) ([]string, error) {
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

// partOne returns the final string after adding, exploding and splitting all input values.
func partOne(in []string) (out string) {
	curStr := in[0]

	for i := 1; i <= len(in); i++ {
		inputSlice, err := scanToSlice(curStr)
		if err != nil {
			fmt.Println("could not scan to slice")
			return
		}

		var newSlice []sf
		var exploded = true
		var isSplit = true

		for exploded || isSplit {
			newSlice, exploded = explode(inputSlice)
			if exploded {
				inputSlice = newSlice
				continue
			}

			newSlice, isSplit = split(inputSlice)
			if isSplit {
				inputSlice = newSlice
				continue
			}
		}

		curStr, err = sliceToString(newSlice)

		if i != len(in) {
			curStr = addString(curStr, in[i])
		} else {
			return curStr
		}
	}
	return ""
}

// getFirstNumber returns the first number (not digit) found in a string and returns an error otherwise.
// This prunes only the hardcoded characters due to task constraints.
func getFirstNumber(in string) (out int, err error) {
	if in == "" {
		return 0, errors.New("empty string provided")
	}

	in = strings.ReplaceAll(in, "[", "")
	in = strings.ReplaceAll(in, "]", "")

	var res = strings.Split(in, ",")

	if len(res) < 1 {
		return 0, errors.New("no number in string")
	}

	for _, r := range res {
		out, err = strconv.Atoi(r)
		if err == nil {
			return out, err
		}
	}

	return 0, errors.New("no number in string")
}

// scanToSlice returns a snailfish slice representation of an input string.
func scanToSlice(in string) (out []sf, err error) {
	var depth = 0
	var onLeft = true // If not on the left number, then waiting on the right number.

	for i := 0; i < len(in); i++ {
		switch in[i] {
		case '[':
			onLeft = true
			depth++
		case ']':
			onLeft = true
			depth--
		case ',':
			onLeft = false
			_, err = strconv.Atoi(string(in[i+1]))
			if err == nil {
				continue
			}
			out = append(out, sf{
				l: sfVal{
					isSet: false,
					value: 0,
				},
				r: sfVal{
					isSet: false,
					value: 0,
				},
				d: depth,
			})
		default:
			num, err := getFirstNumber(in[i:])
			if err != nil {
				return nil, errors.New("scanToSlice tried to access a number where there was none")
			}
			if !onLeft {
				out = append(out, sf{
					l: sfVal{
						isSet: false,
						value: 0,
					},
					r: sfVal{
						isSet: true,
						value: num,
					},
					d: depth,
				})
				i += getDigitsInInt(num) - 1
			} else {
				var leftlength = getDigitsInInt(num)
				i += leftlength + 1
				if in[i] != '[' && in[i] != ']' {
					num2, err := getFirstNumber(in[i:])
					if err != nil {
						return nil, errors.New("scanToSlice2 tried to access a number where there was none")
					}
					out = append(out, sf{
						l: sfVal{
							isSet: true,
							value: num,
						},
						r: sfVal{
							isSet: true,
							value: num2,
						},
						d: depth,
					})
					i += getDigitsInInt(num2) - 1
				} else {
					out = append(out, sf{
						l: sfVal{
							isSet: true,
							value: num,
						},
						r: sfVal{
							isSet: false,
							value: 0,
						},
						d: depth,
					})
					i--
				}
			}
		}
	}
	return out, nil
}

// getDigitsInInt returns the number of digits in the provided integer.
func getDigitsInInt(in int) (out int) {
	out = 1
	in = in / 10
	for in > 0 {
		in = in / 10
		out++
	}
	return out
}

// sliceToString returns the string representation of the provided snailfish slice.
func sliceToString(in []sf) (out string, err error) {
	prevClosed := 0

	for i, s := range in {
		if i == 0 {
			for i2 := 0; i2 < s.d; i2++ {
				out += "["
			}
			if s.l.isSet {
				out += fmt.Sprintf("%d", s.l.value)
			}
			out += ","
			if s.r.isSet {
				out += fmt.Sprintf("%d]", s.r.value)
				prevClosed = 1
			} else {
				prevClosed = 0
			}
		} else {
			if s.l.isSet && s.r.isSet {
				// Both set.
				for i2 := 0 + prevClosed; i2 < s.d-in[i-1].d; i2++ {
					out += "["
				}
				out += fmt.Sprintf("%d,%d]", s.l.value, s.r.value)
				prevClosed = 1
			} else if s.l.isSet && !s.r.isSet {
				// Left set.
				for i2 := 0; i2 < s.d-in[i-1].d; i2++ {
					out += "["
				}
				out += fmt.Sprintf("%d,", s.l.value)
				prevClosed = 0
			} else if !s.l.isSet && s.r.isSet {
				// Right set.
				for i2 := 0 + prevClosed; i2 < in[i-1].d-s.d; i2++ {
					out += "]"
				}
				out += fmt.Sprintf(",%d]", s.r.value)
				prevClosed = 1
			} else {
				// None set.
				for i2 := 0 + prevClosed; i2 < in[i-1].d-s.d; i2++ {
					out += "]"
				}
				out += ","
				prevClosed = 0
			}
		}

		if i == len(in)-1 {
			for i2 := 0 + prevClosed; i2 < s.d; i2++ {
				out += "]"
			}
		}

	}

	return out, nil
}

// addString adds two snailfish string together.
func addString(str1, str2 string) (out string) {
	return fmt.Sprintf("[%s,%s]", str1, str2)
}

// explode tries to explode an eligible snailfish in the input slice.
// a boolean is returned to indicate whether an explosion took place.
func explode(in []sf) (out []sf, exploded bool) {
	var index = -1

	for i, s := range in {
		if s.d >= 5 && s.l.isSet && s.r.isSet {
			index = i
			break
		}
	}

	if index == -1 {
		return in, false
	}

	var left = in[0:index]
	var leftIndex = len(left) - 1
	if len(left) >= 1 {
		for leftIndex >= 0 {
			if left[leftIndex].r.isSet {
				left[leftIndex].r.value += in[index].l.value
				break
			} else if left[leftIndex].l.isSet {
				left[leftIndex].l.value += in[index].l.value
				break
			}
			leftIndex--
		}
	}
	if leftIndex == -1 {
		leftIndex = 0
	}

	var right = in[index+1:]
	var rightIndex = 0
	if len(right) >= 1 {
		for rightIndex < len(right) {
			if right[rightIndex].l.isSet {
				right[rightIndex].l.value += in[index].r.value
				break
			} else if right[rightIndex].r.isSet {
				right[rightIndex].r.value += in[index].r.value
				break
			}
			rightIndex++
		}
	}

	if leftIndex < len(left) {
		if !left[leftIndex].r.isSet && math.Abs(float64(left[leftIndex].d-in[index].d)) == 1.0 {
			left[leftIndex].r.isSet = true
			left[leftIndex].r.value = 0
			return append(left, right...), true
		}
	}

	if rightIndex < len(right) {
		if !right[rightIndex].l.isSet && math.Abs(float64(right[rightIndex].d-in[index].d)) == 1.0 {
			right[rightIndex].l.isSet = true
			right[rightIndex].l.value = 0
			return append(left, right...), true
		} else {
			right[0].l.isSet = true
			right[0].l.value = 0
			return append(left, right...), true
		}
	}

	fmt.Println("Uhh shouldn't be here in explode")

	return []sf{}, false
}

// split tries to split an eligible snailfish value in the input slice.
// a boolean is return to indicate whether a split took place.
func split(in []sf) (out []sf, split bool) {
	var index = -1
	for i, s := range in {
		if s.l.value >= 10 || s.r.value >= 10 {
			index = i
			break
		}
	}

	if index == -1 {
		return in, false
	}

	/*                  [[[[0,7],4],[15,[0,13]]],[1,1]]
	after split:    [[[[0,7],4],[[7,8],[0,13]]],[1,1]]
	after split:    [[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]  */

	var left = make([]sf, len(in[0:index]))
	_ = copy(left, in[0:index])

	var right []sf
	if index < len(in)-1 {
		right = make([]sf, len(in[index+1:]))
		_ = copy(right, in[index+1:])
	}

	var old = in[index]
	var new sf

	if in[index].l.value >= 10 {
		new.l.isSet = true
		new.r.isSet = true
		new.d = old.d + 1
		new.l.value = int(math.Floor(float64(in[index].l.value) / 2))
		new.r.value = int(math.Ceil(float64(in[index].l.value) / 2))

		old.l.isSet = false
		old.l.value = 0

		if right != nil {
			var temp = append(left, new, old)
			return append(temp, right...), true
		}
		return append(left, new, old), true
	} else {
		new.l.isSet = true
		new.r.isSet = true
		new.d = old.d + 1
		new.l.value = int(math.Floor(float64(in[index].r.value) / 2))
		new.r.value = int(math.Ceil(float64(in[index].r.value) / 2))

		old.r.isSet = false
		old.r.value = 0

		if right != nil {
			var temp = append(left, old, new)
			return append(temp, right...), true
		}
		return append(left, old, new), true
	}

	fmt.Println("Uhh shouldn't be here in split")

	return []sf{}, false
}

// getMagnitude returns the magnitude of the provided snailfish slice.
func getMagnitude(in []sf) (out int) {
	var depth = 0

	for _, s := range in {
		if s.d > depth {
			depth = s.d
		}
	}

	for depth >= 0 {
		for i := 0; i < len(in); i++ {
			if in[i].d == depth {
				if !in[i].l.isSet {
					for j := 1; i-j >= 0; j++ {
						if in[i-j].d == in[i].d+1 {
							in[i].l.value = in[i-j].mag
							in[i].l.isSet = true
							break
						}
					}
				}
				if !in[i].r.isSet {
					for j := 1; j+i < len(in); j++ {
						if in[i+j].d == in[i].d+1 {
							in[i].r.value = in[i+j].mag
							in[i].r.isSet = true
							break
						}
					}
				}

				if in[i].l.isSet && in[i].r.isSet {
					in[i].setLocalMagnitude()
				}
			}
		}
		depth--
	}

	var mag = 0
	for i := 0; i < len(in); i++ {
		if in[i].mag > mag {
			mag = in[i].mag
		}
	}

	return mag
}

// setLocalMagnitude sets the magnitude of the calling snailfish.
func (s *sf) setLocalMagnitude() {
	s.mag = (s.l.value * 3) + (s.r.value * 2)
}
