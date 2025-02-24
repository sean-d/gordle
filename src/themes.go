package gordle

type Theme struct {
	Name        string
	Description string
	Words       []string
}

var (
	ClassicTheme = Theme{
		Name:        "Classic",
		Description: "The original list with 5 letter words",
		Words:       fiveLetterWords,
	}

	DifficultTheme = Theme{
		Name:        "Difficult",
		Description: "Challenge yourself with 6 letter words",
		Words:       sixLetterWords,
	}

	HardTheme = Theme{
		Name:        "Hard",
		Description: "Test your skills with 7 letter words",
		Words:       sevenLetterWords,
	}

	SoulsTheme = Theme{
		Name:        "Souls",
		Description: "The ultimate challenge with 8 letter words",
		Words:       eightLetterWords,
	}

	ChaosTheme = Theme{
		Name:        "Chaos",
		Description: "Random word length (5-8 letters) each game - pure chaos!",
		Words:       append(append(append(fiveLetterWords, sixLetterWords...), sevenLetterWords...), eightLetterWords...),
	}
)

var AllThemes = []Theme{
	ClassicTheme,
	DifficultTheme,
	HardTheme,
	SoulsTheme,
	ChaosTheme,
}

var CurrentTheme = ClassicTheme