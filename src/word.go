package gordle

import (
	"errors"
	"math/rand"
	"regexp"
	"sort"
)

type Word struct {
	Letters string
}

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

func isValidWord(letters string) bool {
	// For all themes (including classic), use the current theme's word list
	i := sort.SearchStrings(CurrentTheme.Words, letters)
	return i < len(CurrentTheme.Words) && CurrentTheme.Words[i] == letters
}

func NewWord(letters string) (Word, error) {
	expectedLength := len(CurrentTheme.Words[0])
	if len(letters) != expectedLength || !regexp.MustCompile("[A-Z]+").MatchString(letters) {
		return Word{}, errors.New("invalid word")
	}
	return Word{Letters: letters}, nil
}

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
