// Package gordle implements a Wordle clone game in Go using the Fyne toolkit.
// It provides functionality for word validation, feedback generation, and game state management.
package gordle

import (
	"errors"
	"math/rand"
	"regexp"
	"sort"
)

// Word represents a word in the game, containing a string of uppercase letters.
// The length of the word depends on the current theme's word length (5-8 letters).
type Word struct {
	Letters string
}

// init initializes the word lists by sorting them for efficient binary search.
// This ensures that word validation can be performed quickly during gameplay.
func init() {
	if !sort.StringsAreSorted(fiveLetterWords) {
		sort.Strings(fiveLetterWords)
	}
	// Sort all theme word lists
	for i := range AllThemes {
		if !sort.StringsAreSorted(AllThemes[i].Words) {
			sort.Strings(AllThemes[i].Words)
		}
	}
}

// NewWord creates a new Word instance after validating the input letters.
// It checks that:
// - The length matches the current theme's word length
// - The letters are all uppercase
// Returns an error if validation fails.
func NewWord(letters string) (Word, error) {
	expectedLength := len(CurrentTheme.Words[0])
	if len(letters) != expectedLength || !regexp.MustCompile("[A-Z]+").MatchString(letters) {
		return Word{}, errors.New("invalid word")
	}
	return Word{Letters: letters}, nil
}

// RandomWord generates a random word from the appropriate word list based on the current theme.
// For the Chaos theme, it randomly selects a word length (5-8) before choosing a word.
// If the current theme has no words, it falls back to the Classic theme.
func RandomWord() Word {
	var wordList []string

	if CurrentTheme.Name == "Chaos" {
		// For Chaos mode, randomly select which word list to use
		allWordLists := [][]string{
			fiveLetterWords,
			sixLetterWords,
			sevenLetterWords,
			eightLetterWords,
		}
		selectedList := allWordLists[rand.Intn(len(allWordLists))]
		// Update the current theme's word list
		CurrentTheme.Words = selectedList
		wordList = selectedList
	} else if CurrentTheme.Name != ClassicTheme.Name {
		wordList = CurrentTheme.Words
	} else {
		wordList = fiveLetterWords
	}

	if len(wordList) == 0 {
		// Fallback to classic theme if current theme has no words
		wordList = ClassicTheme.Words
	}
	letters := wordList[rand.Intn(len(wordList))]
	word, _ := NewWord(letters)
	return word
}
