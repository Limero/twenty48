package main

import "github.com/gdamore/tcell/v2"

type Action int

const (
	ActionUp Action = iota
	ActionDown
	ActionLeft
	ActionRight
	ActionReset
	ActionQuit
	ActionNothing
)

func handleInput(screen tcell.Screen) Action {
	ev := screen.PollEvent()
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyRune:
			switch ev.Rune() {
			case 'q', 'Q':
				return ActionQuit
			case 'r':
				return ActionReset
			case 'h', 'a':
				return ActionLeft
			case 'j', 's':
				return ActionDown
			case 'k', 'w':
				return ActionUp
			case 'l', 'd':
				return ActionRight
			}
		case tcell.KeyCtrlL:
			screen.Sync()
		case tcell.KeyCtrlC:
			return ActionQuit
		case tcell.KeyLeft:
			return ActionLeft
		case tcell.KeyUp:
			return ActionUp
		case tcell.KeyDown:
			return ActionDown
		case tcell.KeyRight:
			return ActionRight
		}
	case *tcell.EventResize:
		screen.Sync()
	}
	return ActionNothing
}
