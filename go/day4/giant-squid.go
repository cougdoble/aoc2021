package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

func main() {
	input, _ := readFile("./day4/input.txt")

	part1Res := playBingo(input, 1)
	fmt.Printf("Part 1: %v\n", part1Res)

	part2Res := playBingo(input, 2)
	fmt.Printf("Part 2: %v\n", part2Res)
}

func readFile(path string) (string, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(contents)), nil
}

func playBingo(inputStr string, part int) int {
	input := strings.Split(inputStr, "\n\n")

	bingoNums := funk.Map(strings.Split(input[0], ","), func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}).([]int)

	boards := funk.Map(input[1:], func(s string) Board {
		return createBoard(s)
	}).([]Board)

	var bingos []Board

	for _, num := range bingoNums {
		for _, board := range boards {
			board.MarkSpace(num)
			bingo := board.CheckBingo()
			if bingo && part == 1 {
				fmt.Println("bingo")
				return board.sumBoard(num)
			} else if bingo && part == 2 {
				if !funk.Contains(bingos, board) {
					bingos = append(bingos, board)
				}
				if len(bingos) == len(boards) {
					return board.sumBoard(num)
				}
			}
		}
	}

	return 0
}

type Space struct {
	value  int
	marked bool
}

type Board struct {
	boardMatrix [][]Space
	bingo       bool
}

func createBoard(input string) Board {
	var board Board
	for _, row := range strings.Split(input, "\n") {
		var spaces []Space
		for _, field := range strings.Fields(row) {
			num, _ := strconv.Atoi(field)
			space := Space{num, false}
			spaces = append(spaces, space)
		}
		board.boardMatrix = append(board.boardMatrix, spaces)
	}
	return board
}

func (b *Board) sumBoard(num int) int {
	sum := 0
	for _, row := range b.boardMatrix {
		for _, space := range row {
			if !space.marked {
				sum += space.value
			}
		}
	}
	return sum * num
}

func (b *Board) MarkSpace(num int) {
	for i, row := range b.boardMatrix {
		for j, space := range row {
			if space.value == num && !space.marked {
				b.boardMatrix[i][j].marked = true
			}
		}
	}
}

func (b *Board) CheckBingo() bool {
	for _, row := range b.boardMatrix {
		if isBingo(row) {
			return true
		}

		transposedBoard := transpose(b.boardMatrix)
		for _, row := range transposedBoard {
			if isBingo(row) {
				return true
			}
		}
	}

	return false
}

func isBingo(spaces []Space) bool {
	for _, space := range spaces {
		if !space.marked {
			return false
		}
	}
	return true
}

func transpose(boardVals [][]Space) [][]Space {
	transposed := make([][]Space, len(boardVals[0]))

	for _, row := range boardVals {
		for i, val := range row {
			transposed[i] = append(transposed[i], val)
		}
	}
	return transposed
}
