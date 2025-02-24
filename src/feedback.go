package gordle

import (
	"errors"
	"fmt"
)

type Color string

const (
	Grey   Color = "grey"
	Yellow Color = "yellow"
	Green  Color = "green"
)

type Feedback struct {
	Colors []Color
}

func (feedback *Feedback) IsWin() bool {
	for _, color := range feedback.Colors {
		if color != Green {
			return false
		}
	}
	return true
}

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
