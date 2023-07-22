package main

import (
	"math/rand"
	"time"
)

type Game struct {
	Score    int
	Moves    int
	GameOver bool
	Grid     [][]int
}

func NewGame() Game {
	return Game{
		Score:    0,
		Moves:    0,
		GameOver: false,
		Grid: [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 2, 2},
		},
	}
}

var directions = [4][2]int{
	{-1, 0}, // Up
	{1, 0},  // Down
	{0, -1}, // Left
	{0, 1},  // Right
}

func performAction(grid [][]int, action Action) (bool, int) {
	moved := false
	points := 0
	rows, cols := len(grid), len(grid[0])

	move := func(row, col, dr, dc int) {
		for {
			r, c := row+dr, col+dc

			// Check if the current cell is valid
			if r < 0 || r >= rows || c < 0 || c >= cols {
				break
			}

			// If the current cell is empty, move the number to it
			if grid[r][c] == 0 {
				grid[r][c] = grid[row][col]
				grid[row][col] = 0
				row, col = r, c
				moved = true
			} else if grid[r][c] == grid[row][col] { // If the numbers are the same, merge them
				grid[r][c] *= 2
				grid[row][col] = 0
				moved = true
				points += grid[r][c]
				break
			} else { // If the numbers are different, stop moving
				break
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] != 0 {
				switch action {
				case ActionUp:
					move(i, j, directions[0][0], directions[0][1])
				case ActionDown:
					move(i, j, directions[1][0], directions[1][1])
				case ActionLeft:
					move(i, j, directions[2][0], directions[2][1])
				case ActionRight:
					move(i, j, directions[3][0], directions[3][1])
				}
			}
		}
	}

	return moved, points
}

func hasEmptyCells(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func getEmptyCells(grid [][]int) [][]int {
	emptyCells := [][]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				emptyCells = append(emptyCells, []int{i, j})
			}
		}
	}
	return emptyCells
}

func addNewNumber(grid [][]int) {
	emptyCells := getEmptyCells(grid)
	if len(emptyCells) == 0 {
		return
	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(emptyCells))
	cell := emptyCells[randomIndex]
	if rand.Intn(10) < 9 {
		grid[cell[0]][cell[1]] = 2 // 90% chance for a new tile with the value of 2
	} else {
		grid[cell[0]][cell[1]] = 4 // 10% chance for a new tile with the value of 4
	}
}

func isGameOver(grid [][]int) bool {
	if hasEmptyCells(grid) {
		return false
	}

	// Check for any adjacent cells with the same value
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			val := grid[i][j]
			if (i > 0 && grid[i-1][j] == val) || (i < len(grid)-1 && grid[i+1][j] == val) || (j > 0 && grid[i][j-1] == val) || (j < len(grid[0])-1 && grid[i][j+1] == val) {
				return false
			}
		}
	}

	return true
}