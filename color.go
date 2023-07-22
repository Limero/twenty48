package main

import (
	"github.com/gdamore/tcell/v2"
)

func getColor(number int) tcell.Style {
	style := tcell.StyleDefault

	switch number {
	case 2:
		return style.Background(tcell.ColorWhite).Foreground(tcell.ColorGray)
	case 4:
		return style.Background(tcell.ColorBeige).Foreground(tcell.ColorGray)
	case 8:
		return style.Background(tcell.ColorOrange).Foreground(tcell.ColorGray)
	case 16:
		return style.Background(tcell.ColorOrangeRed).Foreground(tcell.ColorGray)
	case 32:
		return style.Background(tcell.ColorRed)
	case 64:
		return style.Background(tcell.ColorDarkRed)
	case 128:
		return style.Background(tcell.ColorYellow).Foreground(tcell.ColorGray)
	case 256:
		return style.Background(tcell.ColorLightGreen).Foreground(tcell.ColorGray)
	case 512:
		return style.Background(tcell.ColorGreen).Foreground(tcell.ColorGray)
	case 1024:
		return style.Background(tcell.ColorDarkGreen)
	case 2048:
		return style.Background(tcell.ColorBlue)
	}

	return style
}
