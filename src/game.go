// Package gordle implements a Wordle clone game in Go using the Fyne toolkit.
package gordle

import (
	"errors"
)

// GameState represents the current state of a game
type GameState string

const (
	Started GameState = "started" // Game is in progress
	Won     GameState = "won"     // Player has won
	Lost    GameState = "lost"    // Player has lost (used all attempts)
)

const maxTries int = 6

// Game represents a single game session
type Game struct {
	State    GameState
	Tries    int
	Solution Word
}

// NewGame creates a new game with:
// - A random solution word from the current theme
// - Started state
// - 0 attempts
func NewGame() Game {
	return Game{State: "started", Tries: 0, Solution: RandomWord()}
}

// Guess processes a guess attempt and returns:
// - Updated game state
// - Feedback for the guess
// - Error if guess is invalid
//
// The game ends (state changes) if:
// - Player guesses correctly (Won)
// - Player uses all 6 attempts (Lost)
func (game *Game) Guess(guess Word) (Game, Feedback, error) {
	if game.State != Started {
		return *game, Feedback{}, errors.New("Game already finished")
	}
	feedback := FromGuess(guess, game.Solution)
	if feedback.IsWin() {
		return Game{State: Won, Tries: game.Tries + 1, Solution: game.Solution}, feedback, nil
	}
	if game.Tries == maxTries-1 {
		return Game{State: Lost, Tries: game.Tries + 1, Solution: game.Solution}, feedback, nil
	}
	return Game{State: game.State, Tries: game.Tries + 1, Solution: game.Solution}, feedback, nil
}
