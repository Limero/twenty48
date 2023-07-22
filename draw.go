package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func drawLine(screen tcell.Screen, y int, line string) {
	for x, c := range []rune(line) {
		screen.SetContent(x, y, c, nil, tcell.StyleDefault)
	}
}

func drawInfo(screen tcell.Screen, score, moves int) {
	drawLine(screen, 0, fmt.Sprintf("Score: %7d Moves: %7d", score, moves))
}

func drawGameOver(screen tcell.Screen) {
	drawLine(screen, 10, "Game over!")
	drawLine(screen, 11, "Press \"r\" to restart")
}

func drawGrid(screen tcell.Screen, grid [][]int) {
	g := []string{
		"┌──────┬──────┬──────┬──────┐",
		gridRow(grid[0]),
		"├──────┼──────┼──────┼──────┤",
		gridRow(grid[1]),
		"├──────┼──────┼──────┼──────┤",
		gridRow(grid[2]),
		"├──────┼──────┼──────┼──────┤",
		gridRow(grid[3]),
		"└──────┴──────┴──────┴──────┘",
	}

	for y, row := range g {
		drawLine(screen, y+1, row)
	}
}

func gridRow(col []int) string {
	l := "│"

	for _, c := range col {
		if c == 0 {
			l += fmt.Sprintf("%6s│", " ")
			continue
		}
		l += fmt.Sprintf("%6d│", c)
	}

	return l
}