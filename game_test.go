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

	assert.Equal(t, 452, g.Score)
	assert.Equal(t, 74, g.Moves)
	assert.True(t, g.GameOver)
}
