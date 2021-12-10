package main

import (
	"AoC2021/helpers"
	"fmt"
	"os"
	"strings"
)

// a bingoSquare has two fields: the numerical value of the square and a
// boolean of whether that value has been called yet
type bingoSquare struct {
	value  int
	called bool
}

const (
	boardHeight = 5
	boardWidth  = 5
)

func main() {
	// bingoBoards is a slice that holds each bingoBoard read from the input.
	bingoBoards := make([]map[int]bingoSquare, 0)

	input, err := helpers.ReadLinesToString("input.txt")
	if err != nil {
		fmt.Printf("Input read err: %s", err.Error())
		os.Exit(1)
	}

	// bingoOrder in an integer slice that holds the bingo numbers called in order
	bingoOrder, err := helpers.SplitStringToIntSlice(input[0], ",")
	if err != nil {
		fmt.Printf("Input err: %s", err.Error())
		os.Exit(1)
	}

	// bingoBoard is a map used to represent a board of bingoSquares. To make it easy to iterate over
	// the squares, the key of each row is a multiple of 10 and each column increases by 1. Thus the
	// "top left" square of a 5x5 board has a key of 10 and the "bottom right" square has a key of 54.
	bingoBoard := make(map[int]bingoSquare)
	var x = 0
	for i := 2; i < len(input); i++ {
		if input[i] == "" {
			continue
		}
		x += 10

		bingoValues, err := helpers.SplitStringToIntSlice(strings.Join(strings.Fields(input[i]), " "), " ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for y := 0; y < len(bingoValues); y++ {
			bingoBoard[x+y] = bingoSquare{bingoValues[y], false}
		}

		if len(bingoBoard) == boardHeight*boardWidth {
			bingoBoards = append(bingoBoards, bingoBoard)
			bingoBoard = make(map[int]bingoSquare)
			x = 0
		}

	}

	// bingoBoardWinners will holds a boolean for each board index and it gets set to true when the board wins.
	bingoBoardWinners := make(map[int]bool, len(bingoBoard))
	var winnersCount = 0
	var latestWinningScore int
	for i := 0; i < len(bingoOrder); i++ { // Iterate through each called bingo value.
		for j := 0; j < len(bingoBoards); j++ { // Iterate each called value through each board.
			if !bingoBoardWinners[j] { // Skip the boards that have already won.
				latestWinningScore = updateBoardCheckWinner(bingoBoards[j], bingoOrder[i])
				if latestWinningScore >= 0 {
					bingoBoardWinners[j] = true // Once a board has won, we don't need to update it anymore.
					winnersCount++
					if winnersCount == 1 {
						fmt.Printf("Part 1: %d\n", latestWinningScore)
					}
					if winnersCount == len(bingoBoards) {
						fmt.Printf("Part 2: %d", latestWinningScore)
					}
				}
			}
		}
	}
}

// updateBoardCheckWinner updates the board square's boolean to true if it matches
// the given value and checks to see if the board has won whenever a match is found.
// Returns a value of -1 when the board has not won.
func updateBoardCheckWinner(board map[int]bingoSquare, value int) int {
	var updated = false
	for x := 10; x <= boardHeight*10; x += 10 {
		for y := 0; y < boardWidth; y++ {
			if board[x+y].value == value {
				board[x+y] = bingoSquare{value, true}
				updated = true
			}
		}
	}

	if updated {
		for x := 10; x <= boardHeight*10; x += 10 { // Check "rows"
			winner := true
			for y := 0; y < boardWidth; y++ {
				if !board[x+y].called {
					winner = false
				}
			}
			if winner {
				return getWinningValue(board, value)
			}
		}
		for x := 10; x < 10+boardWidth; x++ { // Check "columns"
			winner := true
			for y := 0; y < boardHeight*10; y += 10 {
				if !board[x+y].called {
					winner = false
				}
			}
			if winner {
				return getWinningValue(board, value)
			}
		}
	}

	return -1
}

func getWinningValue(board map[int]bingoSquare, lastVal int) (sum int) {
	for _, number := range board {
		if !number.called {
			sum += number.value
		}
	}
	return sum * lastVal
}
