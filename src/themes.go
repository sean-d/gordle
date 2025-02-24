package gordle

type Theme struct {
	Name        string
	Description string
	Words       []string
}

var (
	AnimalTheme = Theme{
		Name:        "Animals",
		Description: "Test your knowledge of the animal kingdom",
		Words:       animalWords,
	}

	PlantTheme = Theme{
		Name:        "Plants",
		Description: "A botanical challenge",
		Words:       plantWords,
	}

	SportsTheme = Theme{
		Name:        "Sports",
		Description: "For the sports enthusiasts",
		Words:       sportsWords,
	}

	FoodTheme = Theme{
		Name:        "Food & Drink",
		Description: "Delicious five-letter treats",
		Words:       foodWords,
	}

	ClassicTheme = Theme{
		Name:        "Classic",
		Description: "The original word list with all words",
		Words:       fiveLetterWords,
	}
)

var AllThemes = []Theme{
	ClassicTheme,
	AnimalTheme,
	PlantTheme,
	SportsTheme,
	FoodTheme,
}

var CurrentTheme = ClassicTheme