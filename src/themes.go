// Package gordle implements a Wordle clone game in Go using the Fyne toolkit.
package gordle

// Theme represents a game mode with its own word list and characteristics.
// Each theme has a unique name, description, and set of valid words.
type Theme struct {
	Name        string   // Display name of the theme
	Description string   // Brief description of the theme's characteristics
	Words       []string // List of valid words for this theme
}

// Predefined themes with different word lengths and characteristics
var (
	// ClassicTheme is the original Wordle experience with 5-letter words
	ClassicTheme = Theme{
		Name:        "Classic",
		Description: "The original list with 5 letter words",
		Words:       fiveLetterWords,
	}

	// DifficultTheme increases the challenge with 6-letter words
	DifficultTheme = Theme{
		Name:        "Difficult",
		Description: "Challenge yourself with 6 letter words",
		Words:       sixLetterWords,
	}

	// HardTheme further increases difficulty with 7-letter words
	HardTheme = Theme{
		Name:        "Hard",
		Description: "Test your skills with 7 letter words",
		Words:       sevenLetterWords,
	}

	// SoulsTheme provides the ultimate challenge with 8-letter words
	SoulsTheme = Theme{
		Name:        "Souls",
		Description: "The ultimate challenge with 8 letter words",
		Words:       eightLetterWords,
	}

	// ChaosTheme randomly selects word length (5-8) for each game
	// The Words slice contains all words from all lengths for display purposes,
	// but during gameplay, only words of the randomly selected length are used.
	ChaosTheme = Theme{
		Name:        "Chaos",
		Description: "Random word length (5-8 letters) each game - pure chaos!",
		Words:       append(append(append(fiveLetterWords, sixLetterWords...), sevenLetterWords...), eightLetterWords...),
	}
)

// AllThemes contains all available game themes in order of increasing difficulty
var AllThemes = []Theme{
	ClassicTheme,
	DifficultTheme,
	HardTheme,
	SoulsTheme,
	ChaosTheme,
}

// CurrentTheme holds the currently active theme, defaulting to ClassicTheme
var CurrentTheme = ClassicTheme