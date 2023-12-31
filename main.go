package main

import (
	"fmt"
	"time"

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

	game := NewGame(time.Now().UnixNano())

	for {
		screen.Clear()

		drawInfo(screen, game.Score, game.Moves)
		drawGrid(screen, game.Grid)
		if game.GameOver {
			drawGameOver(screen)
		}

		screen.Show()
		action := handleInput(screen)
		switch action {
		case ActionReset:
			game = NewGame(time.Now().UnixNano())
		case ActionUp, ActionDown, ActionLeft, ActionRight:
			if game.GameOver {
				continue
			}
			game.performAction(action)
		case ActionQuit:
			return
		}
	}
}
