package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func main() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)

	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("Error creating screen: %s\n", err)
		return
	}
	defer screen.Fini()

	if err = screen.Init(); err != nil {
		fmt.Printf("Error initializing screen: %s\n", err)
		return
	}

	game := NewGame()

	for {
		screen.Clear()

		drawInfo(screen, game.Score, game.Moves)
		drawGrid(screen, game.Grid)
		if isGameOver(game.Grid) {
			game.GameOver = true
			drawGameOver(screen)
		}

		screen.Show()
		action := handleInput(screen)
		switch action {
		case ActionReset:
			game = NewGame()
		case ActionUp, ActionDown, ActionLeft, ActionRight:
			if game.GameOver {
				continue
			}
			moved, points := performAction(game.Grid, action)
			if moved {
				game.Moves++
				game.Score += points
				addNewNumber(game.Grid)
			}
		case ActionQuit:
			return
		}
	}
}
