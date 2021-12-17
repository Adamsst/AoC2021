package main

import (
	"AoC2021/helpers"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {

	input, err := helpers.ReadLinesToString("input.txt")
	if err != nil {
		fmt.Printf("Input read err: %s", err.Error())
		os.Exit(1)
	}

	re := regexp.MustCompile(`[-]?\d[\d,]?[\d{2}]*`)
	targetBounds := re.FindAllString(input[0], -1)
	targetMinX, _ := strconv.Atoi(targetBounds[0])
	targetMaxX, _ := strconv.Atoi(targetBounds[1])
	targetMinY, _ := strconv.Atoi(targetBounds[2])
	targetMaxY, _ := strconv.Atoi(targetBounds[3])

	// For part 1 I'm assuming there has to be an x value that stops between minX and maxX
	// the input is a fairly large range close to 0 so I'm confident this assumption is fine
	p1 := 0
	maxYVelocity := int(math.Abs(float64(targetMinY))) - 1
	for i := 0; i <= maxYVelocity; i++ {
		p1 += i
	}
	fmt.Printf("Part 1: %d\n", p1)

	// For part 2, every single point in the target area is a valid initial
	// velocity The absolute max x velocity is the targetMaxX value, which would
	// put the projectile on the far side of the target after 1 iteration.
	// Similarly, targetMinY is the lowest possible y velocity.
	p2 := 0
	for x := 0; x <= targetMaxX; x++ {
		for y := targetMinY; y <= maxYVelocity; y++ {
			if landsInRange(x, y, targetMinX, targetMaxX, targetMinY, targetMaxY) {
				p2++
			}
		}
	}
	fmt.Printf("Part 2: %d\n", p2)
}

func landsInRange(xVel, yVel, minX, maxX, minY, maxY int) bool {
	var x = 0
	var y = 0
	for {
		x += xVel
		y += yVel
		if x > maxX || y < minY {
			return false
		}
		if x >= minX && x <= maxX && y >= minY && y <= maxY {
			return true
		}
		if xVel > 0 {
			xVel--
		}
		yVel--
	}
}
