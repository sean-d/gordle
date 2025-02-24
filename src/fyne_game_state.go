// Package gordle implements a Wordle clone game in Go using the Fyne toolkit.
package gordle

import (
	"strings"

	"fyne.io/fyne/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// AppState represents the complete state of the game application.
// It tracks the game progress, user interface state, and any error conditions.
type AppState struct {
	game         Game              // Current game instance with solution and state
	guesses      []Word           // List of words guessed so far
	feedbacks    []Feedback       // Feedback for each guess
	currentWord  string           // Word being typed but not yet submitted
	errorMessage string           // Current error message to display, if any
	letterColors map[string]Color // Keyboard letter colors based on all guesses
	aboutWindow  *fyne.Window    // Reference to about dialog window if open
}

// NewAppState creates a new application state with a fresh game.
// Initializes empty guesses, feedbacks, and letter colors.
func NewAppState() *AppState {
	return &AppState{
		game:         NewGame(),
		guesses:      make([]Word, 0),
		feedbacks:    make([]Feedback, 0),
		currentWord:  "",
		errorMessage: "",
		letterColors: make(map[string]Color),
		aboutWindow:  nil,
	}
}

// typeLetter adds a letter to the current word if:
// - The game is in progress (Started state)
// - The current word length is less than the theme's word length
func (state AppState) typeLetter(letter string) AppState {
	if state.game.State == Started && len(state.currentWord) < len(CurrentTheme.Words[0]) {
		state.currentWord += letter
	}
	return state
}

// backspace removes the last letter from the current word if:
// - The game is in progress (Started state)
// - The current word is not empty
func (state AppState) backspace() AppState {
	if state.game.State == Started && len(state.currentWord) > 0 {
		state.currentWord = state.currentWord[0 : len(state.currentWord)-1]
	}
	return state
}

// enter handles the enter key press which either:
// - Submits the current guess if game is in progress
// - Starts a new game if game is finished (Won or Lost)
// For a guess submission:
// 1. Validates the word
// 2. Gets feedback by comparing against solution
// 3. Updates letter colors on keyboard
// 4. Clears current word for next guess
func (state AppState) enter() AppState {
	if state.game.State == Started {
		guess, err := NewWord(state.currentWord)
		if err != nil {
			return state.setError(err)
		}
		var feedback Feedback
		state.game, feedback, err = state.game.Guess(guess)
		if err != nil {
			return state.setError(err)
		}
		state.guesses = append(state.guesses, guess)
		state.feedbacks = append(state.feedbacks, feedback)

		for i, letter := range strings.Split(guess.Letters, "") {
			color := feedback.Colors[i]
			switch color {
			case Green:
				state.letterColors[letter] = Green
			case Yellow:
				if state.letterColors[letter] != Green {
					state.letterColors[letter] = Yellow
				}
			case Grey:
				if state.letterColors[letter] != Green && state.letterColors[letter] != Yellow {
					state.letterColors[letter] = Grey
				}
			}
		}

		state.currentWord = ""
	} else {
		state = *NewAppState()
	}
	return state
}

// resetError clears any error message in the state
func (state AppState) resetError() AppState {
	state.errorMessage = ""
	return state
}

// setError sets an error message in the state
// The message is capitalized for display purposes
func (state AppState) setError(err error) AppState {
	state.errorMessage = cases.Title(language.English).String(err.Error())
	return state
}

// setAboutWindow stores a reference to the about dialog window
func (state AppState) setAboutWindow(window *fyne.Window) AppState {
	state.aboutWindow = window
	return state
}
