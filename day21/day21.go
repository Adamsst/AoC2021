package main

import (
	"AoC2021/helpers"
	"fmt"
	"os"
	"strconv"
)

type deterministicDice struct {
	number int
}

type player struct {
	position, score int
}

type gs2 struct { // game state 2
	p1Pos   int
	p2Pos   int
	p1Score int
	p2Score int
	p1Turn  bool
}

func main() {
	input, err := helpers.ReadLinesToString("input.txt")
	if err != nil {
		fmt.Printf("Input read err: %s", err.Error())
		os.Exit(1)
	}

	p1Pos, _ := strconv.Atoi(string(input[0][len(input[0])-1]))
	p2Pos, _ := strconv.Atoi(string(input[1][len(input[1])-1]))

	player1 := player{position: p1Pos, score: 0}
	player2 := player{position: p2Pos, score: 0}
	die := deterministicDice{number: 0}
	var rolls = 0
	var p1Turn = true

	for !player1.isWinner(1000) && !player2.isWinner(1000) {
		var points = 0
		if p1Turn {
			points += die.getNumber()
			points += die.getNumber()
			points += die.getNumber()
			player1.setPosition(points)
			p1Turn = false
		} else {
			points += die.getNumber()
			points += die.getNumber()
			points += die.getNumber()
			player2.setPosition(points)
			p1Turn = true
		}
		rolls += 3
	}

	if player1.isWinner(1000) {
		fmt.Printf("Part 1: %d\n", player2.score*rolls)
	} else {
		fmt.Printf("Part 1: %d\n", player1.score*rolls)
	}

	var part2Values = make(map[int]int64, 10) //map[sum of 3 dice values] # of occurrences (universes)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				part2Values[i+j+k]++
			}
		}
	}

	var part2GameStates = make(map[gs2]int64, 88200) // 10 * 10 * 21 * 21 * 2 = 88200
	for a := 1; a <= 10; a++ {  // p1 position
		for b := 1; b <= 10; b++ { // p2 position
			for c := 0; c <= 20; c++ { // p1 score
				for d := 0; d <= 20; d++ { // p2 score
					part2GameStates[gs2{
						p1Pos:   a,
						p2Pos:   b,
						p1Score: c,
						p2Score: d,
						p1Turn:  false}] = 0
					part2GameStates[gs2{
						p1Pos:   a,
						p2Pos:   b,
						p1Score: c,
						p2Score: d,
						p1Turn:  true}] = 0
				}
			}
		}
	}
	p1Pos, _ = strconv.Atoi(string(input[0][len(input[0])-1]))
	p2Pos, _ = strconv.Atoi(string(input[1][len(input[1])-1]))

	part2GameStates[gs2{
		p1Pos:   p1Pos,
		p2Pos:   p2Pos,
		p1Score: 0,
		p2Score: 0,
		p1Turn:  true}]++

	wins := make([]int64, 2) //wins[0] is p1, [1] is p2

	finished := false
	for !finished {
		finished = true
		for k, v := range part2GameStates {
			if v > 0 {
				finished = false
				for k2, v2 := range part2Values {
					if k.p1Turn {
						newPos, score := getNewPosAndScore(k.p1Pos, k2)
						if k.p1Score+score >= 21 {
							wins[0] += v2 * v // inc p1 wins by value
						} else {

							part2GameStates[gs2{
								p1Pos:   newPos,
								p2Pos:   k.p2Pos,
								p1Score: k.p1Score + score,
								p2Score: k.p2Score,
								p1Turn:  false}] += v * v2
						}
					} else {
						newPos, score := getNewPosAndScore(k.p2Pos, k2)
						if k.p2Score+score >= 21 {
							wins[1] += v2 * v // inc p2 wins by value
						} else {
							part2GameStates[gs2{
								p1Pos:   k.p1Pos,
								p2Pos:   newPos,
								p1Score: k.p1Score,
								p2Score: k.p2Score + score,
								p1Turn:  true}] += v * v2
						}
					}
				}
			}
			part2GameStates[k] = 0 // set back to 0
		}
	}

	if wins[0] > wins[1] {
		fmt.Printf("Part 2: %d", wins[0])
	} else {
		fmt.Printf("Part 2: %d", wins[1])
	}
}

func getNewPosAndScore(pos int, roll int) (scoreToAdd int, newPos int) {
	if pos+roll > 10 {
		return pos + roll - 10, pos + roll - 10
	}
	return pos + roll, pos + roll
}

func (d *deterministicDice) getNumber() int {
	if d.number == 100 {
		d.number = 0
	}
	d.number++
	return d.number
}

func (p *player) setPosition(spacesToMove int) {
	p.position += spacesToMove
	for p.position > 10 {
		p.position -= 10
	}
	p.score += p.position
}

func (p *player) isWinner(winningScore int) bool {
	return p.score >= winningScore
}