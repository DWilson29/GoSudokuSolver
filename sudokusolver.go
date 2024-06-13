package main

import "fmt"

const size = 9

func is_valid(board [][]uint8, row uint8, col uint8, digit uint8) bool {
	for i := 0; i < size; i++ {
		if board[row][i] == digit {
			return false
		}
	}
	for i := 0; i < size; i++ {
		if board[i][col] == digit {
			return false
		}
	}
	start_row := (row / 3) * 3
	start_col := (col / 3) * 3
	for j := 0; j < 3; j++ {
		for k := 0; k < 3; k++ {
			if board[int(start_row)+j][int(start_col)+k] == digit {
				return false
			}
		}
	}
	return true
}

func find_empty_cell(board [][]uint8) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func solve_board(board [][]uint8) bool {
	i, j := find_empty_cell(board)
	if i == -1 && j == -1 {
		return true
	}
	for k := 1; k < 10; k++ {
		if is_valid(board, uint8(i), uint8(j), uint8(k)) {
			board[i][j] = uint8(k)
			if solve_board(board) {
				return true
			}
			board[i][j] = 0
		}
	}
	return false
}

func main() {
	fmt.Println("Hello, World!")

	//First we need to pass in a 2D board
	//0 will represent an empty cell
	//1-9 will represent a known cell
	//Example is provided, will modify later
	a := [][]uint8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	fmt.Println(a)

	if solve_board(a) {
		println("Board is solvable!")
		for i := 0; i < 9; i++ {
			fmt.Println(a[i])
		}
	} else {
		println("Board is not solvable!")
	}
}
