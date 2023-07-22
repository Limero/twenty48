package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	g := NewGame()
	for !isGameOver(g.Grid) {
		for i := 0; i < 4; i++ {
			g.performAction(Action(i))
		}
	}

	assert.Greater(t, g.Score, 0)
	assert.Greater(t, g.Moves, 0)
	assert.True(t, g.GameOver)
}
