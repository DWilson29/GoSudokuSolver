package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

const size = 9

func is_valid(board [9][9]uint8, row uint8, col uint8, digit uint8) bool {
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

func find_empty_cell(board [9][9]uint8) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func solve_board(board [9][9]uint8) bool {
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
	fmt.Println("Welcome to Dawn's Sudoku Solver!")
	fmt.Print("Enter 1 to use an example sudoku board\n" +
		"Enter 2 to use a CSV file\n" +
		"Enter 3 to manually enter board state\n")

	var input int

	fmt.Scanln(&input)

	//First we need to pass in a 2D board
	//0 will represent an empty cell
	//1-9 will represent a known cell
	var a [9][9]uint8

	switch input {
	case 1:
		//Example is provided, will modify later
		a = [9][9]uint8{
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
	case 2:
		fmt.Println("Enter the filepath of your CSV")
		var filepath string
		fmt.Scanln(&filepath)
		f, err := os.Open(filepath)
		if err != nil {
			log.Fatal(err)
		}
		csvReader := csv.NewReader(f)
		data, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				val, err := strconv.Atoi(data[i][j])
				if err != nil {
					log.Fatal(err)
				}
				a[i][j] = uint8(val)
			}
		}
	case 3:
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				fmt.Printf("Enter digit for Row : %d, Col : %d\n", i+1, j+1)
				var digit int
				_, err := fmt.Scanln(&digit)
				if err != nil {
					log.Fatal(err)
				}
				a[i][j] = uint8(digit)
			}
		}
	default:
		fmt.Println("Invalid input!")
		return
	}

	fmt.Println("Using board")
	for i := 0; i < 9; i++ {
		fmt.Println(a[i])
	}

	if solve_board(a) {
		println("Board is solvable!")
		for i := 0; i < 9; i++ {
			fmt.Println(a[i])
		}
	} else {
		println("Board is not solvable!")
	}
}
