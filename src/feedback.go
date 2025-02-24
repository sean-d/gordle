// Package gordle implements a Wordle clone game in Go using the Fyne toolkit.
package gordle

import (
	"errors"
	"fmt"
)

// Color represents the feedback color for a letter in a guess.
// There are three possible colors:
// - Grey: Letter is not in the word
// - Yellow: Letter is in the word but in the wrong position
// - Green: Letter is in the word and in the correct position
type Color string

const (
	Grey   Color = "grey"   // Letter not in word
	Yellow Color = "yellow" // Letter in word, wrong position
	Green  Color = "green"  // Letter in word, correct position
)

// Feedback represents the colored feedback for a guess.
// Each letter in the guess gets a color indicating how close it is to the solution.
type Feedback struct {
	Colors []Color // Colors for each letter position
}

// IsWin checks if the feedback represents a winning guess.
// Returns true if all colors are Green, indicating all letters are in correct positions.
func (feedback *Feedback) IsWin() bool {
	for _, color := range feedback.Colors {
		if color != Green {
			return false
		}
	}
	return true
}

// NewFeedback creates a new Feedback instance after validating the input colors.
// It checks that:
// - The number of colors matches the current theme's word length
// - All colors are valid (Grey, Yellow, or Green)
// Returns an error if validation fails.
func NewFeedback(colors []Color) (Feedback, error) {
	expectedLength := len(CurrentTheme.Words[0])
	if len(colors) != expectedLength {
		return Feedback{}, errors.New("need " + fmt.Sprint(expectedLength) + " colors")
	}
	for _, color := range colors {
		if color != Grey && color != Yellow && color != Green {
			return Feedback{}, errors.New("invalid color")
		}
	}
	return Feedback{colors}, nil
}

// FromGuess generates feedback by comparing a guess word against the solution.
// The algorithm:
// 1. First marks all exact matches (same letter, same position) as Green
// 2. Then marks letters that exist in solution but in wrong position as Yellow
// 3. Remaining letters are marked as Grey
// This ensures proper handling of duplicate letters in the guess or solution.
func FromGuess(guess Word, solution Word) Feedback {
	wordLength := len(solution.Letters)
	colors := make([]Color, wordLength)
	excluded := make([]bool, wordLength)

	for i := range colors {
		colors[i] = Grey
		excluded[i] = false
	}

	for i := range colors {
		if guess.Letters[i] == solution.Letters[i] {
			colors[i] = Green
			excluded[i] = true
		}
	}

	for i := range colors {
		if colors[i] != Grey {
			continue
		}
		for j := range colors {
			if excluded[j] {
				continue
			}
			if guess.Letters[i] == solution.Letters[j] {
				colors[i] = Yellow
				excluded[j] = true
				break
			}
		}
	}

	feedback, _ := NewFeedback(colors)
	return feedback
}
