package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	g := NewGame(1)

	for !isGameOver(g.Grid) {
		for i := 0; i < 4; i++ {
			g.performAction(Action(i))
		}
	}

	assert.Equal(t, g.Score, 800)
	assert.Equal(t, g.Moves, 99)
	assert.True(t, g.GameOver)
}
