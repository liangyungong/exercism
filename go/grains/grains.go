package grains

import (
	"errors"
)

const size = 64

func chessBoard() *[size]uint64 {
	board := [size]uint64{}

	var grid_grain uint64 = 1
	for i := 0; i < size; i++ {
		board[i] = grid_grain
		grid_grain *= 2
	}
	return &board
}

func Square(number int) (uint64, error) {
	board := chessBoard()
	if number <= 0 || number > 64 {
		return 0, errors.New("invalid grid number")
	}
	return board[number-1], nil
}

func Total() uint64 {
	var total uint64 = 0
	for _, value := range chessBoard() {
		total += value
	}

	return total
}
