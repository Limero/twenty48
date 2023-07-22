package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	g := NewGame()
	for !isGameOver(g.Grid) {
		g.performAction(Action(rand.Intn(4)))
	}

	assert.Greater(t, g.Score, 0)
	assert.Greater(t, g.Moves, 0)
	assert.True(t, g.GameOver)
}
