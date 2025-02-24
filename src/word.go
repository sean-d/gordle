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
	if !sort.StringsAreSorted(validWords) {
		sort.Strings(validWords)
	}
	// Sort all theme word lists
	for i := range AllThemes {
		if !sort.StringsAreSorted(AllThemes[i].Words) {
			sort.Strings(AllThemes[i].Words)
		}
	}
}

func isValidWord(letters string) bool {
	// For themed games, only allow words from the current theme
	if CurrentTheme.Name != ClassicTheme.Name {
		i := sort.SearchStrings(CurrentTheme.Words, letters)
		return i < len(CurrentTheme.Words) && CurrentTheme.Words[i] == letters
	}
	// For classic mode, use the full word list
	i := sort.SearchStrings(validWords, letters)
	return i < len(validWords) && validWords[i] == letters
}

var validWordRegex = regexp.MustCompile("[A-Z]{5}")

func NewWord(letters string) (Word, error) {
	if !validWordRegex.MatchString(letters) {
		return Word{}, errors.New("invalid word")
	}
	if !isValidWord(letters) {
		if CurrentTheme.Name != ClassicTheme.Name {
			return Word{}, errors.New("word not in " + CurrentTheme.Name + " theme")
		}
		return Word{}, errors.New("unknown word")
	}
	return Word{Letters: letters}, nil
}

func RandomWord() Word {
	var wordList []string
	if CurrentTheme.Name != ClassicTheme.Name {
		wordList = CurrentTheme.Words
	} else {
		wordList = validWords
	}
	if len(wordList) == 0 {
		// Fallback to classic theme if current theme has no words
		wordList = ClassicTheme.Words
	}
	letters := wordList[rand.Intn(len(wordList))]
	word, _ := NewWord(letters)
	return word
}
