package main

import (
	"fmt"
)

const N = 9

// Print the board
func printBoard(board [N][N]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}
}

// Check if placing num at board[row][col] is valid
func isSafe(board [N][N]int, row, col, num int) bool {
	// Check the row
	for x := 0; x < N; x++ {
		if board[row][x] == num {
			return false
		}
	}

	// Check the column
	for x := 0; x < N; x++ {
		if board[x][col] == num {
			return false
		}
	}

	// Check the 3x3 subgrid
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}

// Find an empty location on the board
func findEmptyLocation(board [N][N]int) (int, int, bool) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == 0 {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

// Solve the Sudoku puzzle using backtracking
func solveSudoku(board *[N][N]int) bool {
	row, col, found := findEmptyLocation(*board)
	if !found {
		return true // Puzzle solved
	}

	for num := 1; num <= 9; num++ {
		if isSafe(*board, row, col, num) {
			board[row][col] = num

			if solveSudoku(board) {
				return true
			}

			board[row][col] = 0 // Backtrack
		}
	}

	return false
}

func main() {
	board := [N][N]int{
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

	if solveSudoku(&board) {
		fmt.Println("Sudoku solved successfully:")
		printBoard(board)
	} else {
		fmt.Println("No solution exists for the given Sudoku.")
	}
}
