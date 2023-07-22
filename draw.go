package main

import (
	"fmt"
	"strings"

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
		"│      │      │      │      │",
		"├──────┼──────┼──────┼──────┤",
		"│      │      │      │      │",
		"├──────┼──────┼──────┼──────┤",
		"│      │      │      │      │",
		"├──────┼──────┼──────┼──────┤",
		"│      │      │      │      │",
		"└──────┴──────┴──────┴──────┘",
	}

	for y, row := range g {
		drawLine(screen, y+1, row)
	}

	for y, row := range grid {
		for x, col := range row {
			if col == 0 {
				continue
			}

			for i, c := range []rune(centerString(fmt.Sprintf("%d", col), 6)) {
				screen.SetContent((1+7*x)+i, (y*2)+2, c, nil, tcell.StyleDefault)
			}
		}
	}
}

func centerString(s string, width int) string {
	spaces := width - len(s)
	leftSpaces := spaces / 2
	rightSpaces := spaces - leftSpaces

	if spaces%2 != 0 {
		leftSpaces++
	}

	s = strings.Repeat(" ", leftSpaces) + s + strings.Repeat(" ", rightSpaces)
	return s[:width]
}
