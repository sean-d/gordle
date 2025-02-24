//go:generate fyne bundle -package gordle -o resources.go assets/AppIcon.svg
//go:generate fyne bundle -package gordle -o resources.go -append assets/AppIcon.png
//go:generate fyne bundle -package gordle -o resources.go -append assets/example.png
//go:generate fyne bundle -package gordle -o resources.go -append assets/about_part1.md
//go:generate fyne bundle -package gordle -o resources.go -append assets/about_part2.md

// Package gordle implements a Wordle clone game in Go using the Fyne toolkit.
package gordle

import (
	"fmt"
	c "image/color"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Color constants for the game UI
var (
	green = c.RGBA{R: 106, G: 170, B: 100, A: 255} // Correct letter, correct position
	yellow = c.RGBA{R: 201, G: 180, B: 88, A: 255} // Correct letter, wrong position
	grey = c.RGBA{R: 120, G: 124, B: 126, A: 255} // Letter not in word
	darkGrey = c.RGBA{R: 58, G: 58, B: 60, A: 255} // UI element background (dark mode)
	lightGrey = c.RGBA{R: 211, G: 215, B: 218, A: 255} // UI element background (light mode)
	black = c.Black
	white = c.White
	darkBackground = c.RGBA{R: 18, G: 18, B: 19, A: 255} // Main background (dark mode)
)

// customTheme implements a custom Fyne theme with dark/light mode support
type customTheme struct {
	fyne.Theme
	isDark bool
}

// Color returns the appropriate color for UI elements based on theme variant
func (t *customTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) c.Color {
	if t.isDark {
		if n == theme.ColorNameBackground {
			return darkBackground
		}
		if n == theme.ColorNameForeground || n == theme.ColorNamePrimary {
			return lightGrey
		}
		if n == theme.ColorNameButton {
			return darkGrey
		}
		if n == theme.ColorNameHover {
			return grey
		}
		return t.Theme.Color(n, theme.VariantDark)
	}
	return t.Theme.Color(n, theme.VariantLight)
}

// Size returns the size for UI elements
func (t *customTheme) Size(n fyne.ThemeSizeName) float32 {
	return t.Theme.Size(n)
}

// Font returns the font for text elements
func (t *customTheme) Font(s fyne.TextStyle) fyne.Resource {
	return t.Theme.Font(s)
}

// Icon returns the icon for UI elements
func (t *customTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return t.Theme.Icon(n)
}

// getThemeColors returns the appropriate colors for text, background, and keyboard
// based on whether dark mode is enabled
func getThemeColors(isDark bool) (textColor, bgColor, keyBgColor c.Color) {
	if isDark {
		return lightGrey, darkBackground, darkGrey
	}
	return black, white, lightGrey
}

// StartFyneGame initializes and starts the Gordle game application.
// It sets up the main window, theme, and initial game state.
func StartFyneGame() {
	app := app.New()
	app.SetIcon(resourceAppIconPng)
	app.Settings().SetTheme(&customTheme{Theme: theme.DefaultTheme(), isDark: true})
	window := app.NewWindow("Gordle")
	window.SetFixedSize(true)
	window.SetCloseIntercept(func() {
		app.Quit()
	})

	isDark := true
	_, bgColor, _ := getThemeColors(isDark)
	bg := canvas.NewRectangle(bgColor)

	// Create initial content container
	content := container.NewMax(bg)
	window.SetContent(content)

	showThemeSelection(app, window)
	window.ShowAndRun()
}

// showThemeSelection displays the theme selection screen where players can:
// - Choose between light and dark mode
// - Select a game theme (Classic, Difficult, etc.)
// - View theme descriptions and word counts
func showThemeSelection(app fyne.App, window fyne.Window) {
	// Get current dark mode setting
	isDark := false
	if ct, ok := app.Settings().Theme().(*customTheme); ok {
		isDark = ct.isDark
	}

	// Create a fresh theme instance
	newTheme := &customTheme{Theme: theme.DefaultTheme(), isDark: isDark}
	app.Settings().SetTheme(newTheme)

	// Create background
	_, bgColor, _ := getThemeColors(isDark)
	bg := canvas.NewRectangle(bgColor)

	mainContainer := container.NewVBox()

	// Add header with icon
	header := container.New(layout.NewVBoxLayout())
	iconRect := canvas.NewRectangle(c.Transparent)
	iconRect.SetMinSize(fyne.NewSize(48, 48))
	icon := widget.NewIcon(resourceAppIconSvg)
	iconBox := container.NewStack(iconRect, icon)

	// Update title and description colors based on theme
	textColor, _, _ := getThemeColors(isDark)

	title := canvas.NewText("Welcome To Gordle", textColor)
	title.TextSize = 24
	title.TextStyle.Bold = true
	title.Alignment = fyne.TextAlignCenter

	titleRow := container.New(layout.NewBorderLayout(nil, nil, iconBox, nil))
	titleRow.Add(iconBox)
	titleRow.Add(title)

	header.Add(titleRow)

	borderColor := lightGrey
	if isDark {
		borderColor = darkGrey
	}
	border := canvas.NewRectangle(borderColor)
	border.SetMinSize(fyne.NewSize(0, 5))
	header.Add(border)

	mainContainer.Add(header)

	// Add mode selection buttons
	modeContainer := container.NewHBox()

	lightButton := widget.NewButton("eye bleed mode (light)", func() {
		newTheme := &customTheme{Theme: theme.DefaultTheme(), isDark: false}
		app.Settings().SetTheme(newTheme)
		showThemeSelection(app, window)
	})
	darkButton := widget.NewButton("save your eyes mode (dark)", func() {
		newTheme := &customTheme{Theme: theme.DefaultTheme(), isDark: true}
		app.Settings().SetTheme(newTheme)
		showThemeSelection(app, window)
	})

	// Style the mode buttons
	lightButtonContainer := decorateButton(app, lightButton, nil, nil, fyne.NewSize(120, 40))
	darkButtonContainer := decorateButton(app, darkButton, nil, nil, fyne.NewSize(120, 40))

	modeContainer.Add(lightButtonContainer)
	modeContainer.Add(widget.NewLabel("|"))
	modeContainer.Add(darkButtonContainer)

	// Center the mode buttons
	centeredModeContainer := container.New(layout.NewCenterLayout(), modeContainer)
	mainContainer.Add(centeredModeContainer)

	// Add spacing after mode buttons
	spaceAfterButtons := canvas.NewRectangle(c.Transparent)
	spaceAfterButtons.SetMinSize(fyne.NewSize(0, 10))
	mainContainer.Add(spaceAfterButtons)

	description := canvas.NewText("select your difficulty", textColor)
	description.TextSize = 16
	description.Alignment = fyne.TextAlignCenter
	mainContainer.Add(description)

	space := canvas.NewRectangle(c.Transparent)
	space.SetMinSize(fyne.NewSize(0, 20))
	mainContainer.Add(space)

	for _, theme := range AllThemes {
		themeCard := createThemeCard(theme, app, window)
		mainContainer.Add(themeCard)

		// Add spacing between cards
		space := canvas.NewRectangle(c.Transparent)
		space.SetMinSize(fyne.NewSize(0, 10))
		mainContainer.Add(space)
	}

	window.SetContent(container.NewMax(bg, mainContainer))
	window.Resize(fyne.NewSize(400, 500))
	window.Show()
}

// createThemeCard creates a clickable card for a game theme.
// Each card displays:
// - Theme name in large, bold text
// - Theme description
// - Number of available words
// The card triggers theme selection and game start when clicked.
func createThemeCard(theme Theme, app fyne.App, window fyne.Window) *fyne.Container {
	// Get current theme state
	isDark := false
	if ct, ok := app.Settings().Theme().(*customTheme); ok {
		isDark = ct.isDark
	}
	textColor, _, _ := getThemeColors(isDark)

	card := container.NewVBox()

	// Create a button that looks like a card
	btn := widget.NewButton("", func() {
		// Set theme and create new state
		CurrentTheme = theme
		state := NewAppState()

		// Ensure we have a valid game state
		if state == nil || state.game.Solution.Letters == "" {
			CurrentTheme = ClassicTheme
			state = NewAppState()
		}

		startGame(app, window)
	})

	// Theme name
	name := canvas.NewText(theme.Name, textColor)
	name.TextSize = 20
	name.TextStyle.Bold = true

	// Theme description
	desc := canvas.NewText(theme.Description, textColor)
	desc.TextSize = 14

	// Word count
	count := canvas.NewText(fmt.Sprintf("%d words", len(theme.Words)), textColor)
	count.TextSize = 12

	content := container.NewVBox(name, desc, count)

	// Create a card-like container
	card = container.NewMax(
		btn,
		container.NewPadded(content),
	)

	// Set minimum size for consistent card appearance
	rect := canvas.NewRectangle(c.Transparent)
	rect.SetMinSize(fyne.NewSize(350, 80))

	return container.NewStack(rect, card)
}

// startGame initializes a new game with the selected theme.
// It sets up:
// - Keyboard input handling
// - Game board display
// - Game state management
func startGame(app fyne.App, window fyne.Window) {
	// Get current theme state
	isDark := false
	if ct, ok := app.Settings().Theme().(*customTheme); ok {
		isDark = ct.isDark
	}

	// Ensure theme is properly set
	app.Settings().SetTheme(&customTheme{Theme: theme.DefaultTheme(), isDark: isDark})

	// Create new state and ensure it's valid
	state := NewAppState()
	if state == nil || state.game.Solution.Letters == "" {
		CurrentTheme = ClassicTheme
		state = NewAppState()
	}

	render(app, state, window)

	mappings := map[fyne.KeyName]string{
		fyne.KeyA: "A",
		fyne.KeyB: "B",
		fyne.KeyC: "C",
		fyne.KeyD: "D",
		fyne.KeyE: "E",
		fyne.KeyF: "F",
		fyne.KeyG: "G",
		fyne.KeyH: "H",
		fyne.KeyI: "I",
		fyne.KeyJ: "J",
		fyne.KeyK: "K",
		fyne.KeyL: "L",
		fyne.KeyM: "M",
		fyne.KeyN: "N",
		fyne.KeyO: "O",
		fyne.KeyP: "P",
		fyne.KeyQ: "Q",
		fyne.KeyR: "R",
		fyne.KeyS: "S",
		fyne.KeyT: "T",
		fyne.KeyU: "U",
		fyne.KeyV: "V",
		fyne.KeyW: "W",
		fyne.KeyX: "X",
		fyne.KeyY: "Y",
		fyne.KeyZ: "Z",
	}

	window.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		if state.game.State != Started && key.Name == fyne.KeySpace {
			showThemeSelection(app, window)
			return
		}
		if letter, exists := mappings[key.Name]; exists {
			*state = state.typeLetter(letter)
			render(app, state, window)
		} else if key.Name == fyne.KeyBackspace {
			*state = state.backspace()
			render(app, state, window)
		} else if key.Name == fyne.KeyReturn {
			*state = state.enter()
			if state.errorMessage != "" {
				displayError(app, state, window)
			}
			render(app, state, window)
		}
	})
}

var errorTicker *time.Ticker

// displayError shows an error message for 1 second then auto-clears it
func displayError(app fyne.App, state *AppState, window fyne.Window) {
	if errorTicker != nil {
		errorTicker.Stop()
	}
	errorTicker = time.NewTicker(1 * time.Second)

	theState := state
	go func() {
		<-errorTicker.C
		errorTicker.Stop()
		*theState = theState.resetError()
		render(app, theState, window)
	}()
}

// render updates the game UI to reflect the current state.
// It displays:
// - Game header with icon
// - Status messages (errors, win/loss)
// - Word grid with guesses and feedback
// - Virtual keyboard with color-coded keys
func render(app fyne.App, state *AppState, window fyne.Window) {
	isDark := false
	if ct, ok := app.Settings().Theme().(*customTheme); ok {
		isDark = ct.isDark
	}
	_, bgColor, _ := getThemeColors(isDark)

	// Create background
	bg := canvas.NewRectangle(bgColor)

	space := canvas.NewRectangle(bgColor)
	space.SetMinSize(fyne.NewSize(0, 40))

	content := container.NewVBox(
		header(app, state),
		statusMessage(app, state),
		wordRows(app, state),
		space,
		keyboard(app, state, window))

	window.SetContent(container.NewMax(bg, content))
}

// header creates the game header with:
// - App icon
// - Title
// - Help button
// - Separator line
func header(app fyne.App, state *AppState) *fyne.Container {
	header := container.New(layout.NewVBoxLayout())

	iconRect := canvas.NewRectangle(c.Transparent)
	iconRect.SetMinSize(fyne.NewSize(48, 48))
	icon := widget.NewIcon(resourceAppIconSvg)
	iconBox := container.NewStack(iconRect, icon)

	isDark := false
	if ct, ok := app.Settings().Theme().(*customTheme); ok {
		isDark = ct.isDark
	}
	textColor, _, _ := getThemeColors(isDark)

	title := canvas.NewText("Gordle", textColor)
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24
	title.TextStyle.Bold = true

	// Style the help button for dark mode
	var helpBgColor, helpFgColor c.Color
	if isDark {
		helpBgColor = darkGrey
		helpFgColor = lightGrey
	}
	helpButtonRect := canvas.NewRectangle(c.Transparent)
	helpButtonRect.SetMinSize(fyne.NewSize(48, 48))
	helpButton := widget.NewButton("?", func() { openAboutDialog(app, state) })
	helpButtonBox := container.NewStack(
		helpButtonRect,
		decorateButton(app, helpButton, helpBgColor, helpFgColor, fyne.NewSize(48, 48)),
	)

	// Add both theme and help buttons to the right
	rightButtons := container.NewHBox(helpButtonBox)
	titleRow := container.New(layout.NewBorderLayout(nil, nil, iconBox, rightButtons))
	titleRow.Add(iconBox)
	titleRow.Add(rightButtons)
	titleRow.Add(title)

	header.Add(titleRow)

	borderColor := lightGrey
	if isDark {
		borderColor = darkGrey
	}
	border := canvas.NewRectangle(borderColor)
	border.SetMinSize(fyne.NewSize(0, 5))
	header.Add(border)

	return header
}

// openAboutDialog displays the help/about window with:
// - Game rules
// - Example gameplay
// - Color meaning explanations
func openAboutDialog(app fyne.App, state *AppState) {
	if state.aboutWindow == nil {
		window := app.NewWindow("Gordle")
		window.SetFixedSize(true)

		aboutPart1 := widget.NewRichTextFromMarkdown(string(resourceAboutpart1Md.Content()))
		exampleRect := canvas.NewRectangle(c.Transparent)
		exampleRect.SetMinSize(fyne.NewSize(331, 69))
		example := container.New(layout.NewStackLayout(), exampleRect, widget.NewIcon(resourceExamplePng))
		aboutPart2 := widget.NewRichTextFromMarkdown(string(resourceAboutpart2Md.Content()))

		about := container.NewVBox(aboutPart1, container.New(layout.NewHBoxLayout(), example), aboutPart2)
		window.SetContent(about)
		window.Show()
		window.SetCloseIntercept(func() {
			window.Close()
			state.aboutWindow = nil
		})
		state.setAboutWindow(&window)
	} else {
		(*state.aboutWindow).RequestFocus()
	}
}

// statusMessage creates a message box for:
// - Win/loss messages
// - Error messages
// - Instructions for next action
func statusMessage(app fyne.App, state *AppState) *fyne.Container {
	isDark := false
	if ct, ok := app.Settings().Theme().(*customTheme); ok {
		isDark = ct.isDark
	}
	_, bgColor, _ := getThemeColors(isDark)

	message := ""
	if state.game.State == Won {
		message = "You won! Press ENTER for a new game, or SPACE to change theme."
	} else if state.game.State == Lost {
		message = "The solution was " + state.game.Solution.Letters + ". Press ENTER for a new game, or SPACE to change theme."
	} else if state.errorMessage != "" {
		message = state.errorMessage
	}

	// Create container with theme background
	messageBox := canvas.NewRectangle(bgColor)
	messageBox.SetMinSize(fyne.NewSize(0, 50))

	// Only show black background when there's a message
	if message != "" {
		messageBox.FillColor = black
		messageBox.StrokeWidth = 15
		messageBox.StrokeColor = black
	}

	statusText := canvas.NewText(message, white)
	statusText.Alignment = fyne.TextAlignCenter
	statusText.TextSize = 14
	statusText.TextStyle.Bold = true

	return container.NewStack(messageBox, statusText)
}

// wordRows creates the game board grid showing:
// - Previous guesses with color feedback
// - Current guess being typed
// - Empty rows for remaining guesses
func wordRows(app fyne.App, state *AppState) *fyne.Container {
	rows := container.New(layout.NewVBoxLayout())
	for i, guess := range state.guesses {
		feedback := state.feedbacks[i]
		rows.Add(wordRow(guess, feedback))
	}
	remaining := 6 - len(state.guesses)
	if state.game.State == Started {
		remaining = remaining - 1
		rows.Add(currentWordRow(app, state.currentWord))
	}
	for i := 0; i < remaining; i++ {
		rows.Add(emptyWordRow(app))
	}
	return container.New(layout.NewCenterLayout(), rows)
}

func wordRow(word Word, feedback Feedback) *fyne.Container {
	wordLength := len(word.Letters)
	grid := container.New(layout.NewGridLayout(wordLength))
	for i, letter := range strings.Split(word.Letters, "") {
		color := feedback.Colors[i]
		grid.Add(letterBox(letter, color))
	}
	return grid
}

func currentWordRow(app fyne.App, word string) *fyne.Container {
	wordLength := len(CurrentTheme.Words[0]) // Get length from current theme's words
	grid := container.New(layout.NewGridLayout(wordLength))
	for _, letter := range strings.Split(word, "") {
		grid.Add(currentWordLetterBox(app, letter))
	}
	for i := len(word); i < wordLength; i++ {
		grid.Add(emptyLetterBox(app))
	}
	return grid
}

func emptyWordRow(app fyne.App) *fyne.Container {
	wordLength := len(CurrentTheme.Words[0]) // Get length from current theme's words
	grid := container.New(layout.NewGridLayout(wordLength))
	for i := 0; i < wordLength; i++ {
		grid.Add(emptyLetterBox(app))
	}
	return grid
}

func letterBox(letter string, color Color) *fyne.Container {
	var fill c.Color
	switch color {
	case Green:
		fill = green
	case Yellow:
		fill = yellow
	case Grey:
		fill = grey
	}

	box := canvas.NewRectangle(fill)
	box.SetMinSize(fyne.NewSize(62, 62))

	text := canvas.NewText(letter, white)
	text.Alignment = fyne.TextAlignCenter
	text.TextSize = 32
	text.TextStyle.Bold = true
	return container.NewStack(box, text)
}

func currentWordLetterBox(app fyne.App, letter string) *fyne.Container {
	isDark := false
	if ct, ok := app.Settings().Theme().(*customTheme); ok {
		isDark = ct.isDark
	}
	textColor, bgColor, _ := getThemeColors(isDark)

	box := canvas.NewRectangle(bgColor)
	box.SetMinSize(fyne.NewSize(62, 62))
	box.StrokeColor = darkGrey
	box.StrokeWidth = 2.0

	text := canvas.NewText(letter, textColor)
	text.Alignment = fyne.TextAlignCenter
	text.TextSize = 32
	text.TextStyle.Bold = true
	return container.NewStack(box, text)
}

func emptyLetterBox(app fyne.App) *canvas.Rectangle {
	isDark := false
	if ct, ok := app.Settings().Theme().(*customTheme); ok {
		isDark = ct.isDark
	}
	_, bgColor, _ := getThemeColors(isDark)

	box := canvas.NewRectangle(bgColor)
	box.SetMinSize(fyne.NewSize(62, 62))
	box.StrokeColor = darkGrey
	box.StrokeWidth = 2.0
	return box
}

// keyboard creates the virtual keyboard with:
// - Letter keys in QWERTY layout
// - Enter and Backspace keys
// - Color-coded feedback based on guesses
func keyboard(app fyne.App, state *AppState, window fyne.Window) *fyne.Container {
	isDark := false
	if ct, ok := app.Settings().Theme().(*customTheme); ok {
		isDark = ct.isDark
	}
	_, bgColor, _ := getThemeColors(isDark)

	letterRows := [][]string{
		{"Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P"},
		{"A", "S", "D", "F", "G", "H", "J", "K", "L"},
		{"Z", "X", "C", "V", "B", "N", "M"},
	}

	// Create keyboard container with background
	keyboardBg := canvas.NewRectangle(bgColor)
	keyboardContent := container.New(layout.NewVBoxLayout())

	for i, letters := range letterRows {
		row := container.New(layout.NewHBoxLayout())
		rowBg := canvas.NewRectangle(bgColor)
		rowContainer := container.NewMax(rowBg, row)

		if i == 2 {
			theState := state
			enterButton := widget.NewButton("ENTER", func() {
				*theState = theState.enter()
				if theState.errorMessage != "" {
					displayError(app, theState, window)
				}
				render(app, theState, window)
			})
			decoratedEnterButton := decorateButton(app, enterButton, nil, nil, fyne.NewSize(65.4, 58))
			row.Add(decoratedEnterButton)
		}
		for _, letter := range letters {
			theLetter := letter
			theState := state
			button := widget.NewButton(theLetter, func() {
				*theState = theState.typeLetter(theLetter)
				render(app, state, window)
			})
			decoratedButton := decorateLetterButton(app, button, state)
			row.Add(decoratedButton)
		}
		if i == 2 {
			theState := state
			backButton := widget.NewButton("BACK", func() {
				*theState = theState.backspace()
				render(app, theState, window)
			})
			decoratedBackButton := decorateButton(app, backButton, nil, nil, fyne.NewSize(65.4, 58))
			row.Add(decoratedBackButton)
		}
		keyboardContent.Add(container.New(layout.NewCenterLayout(), rowContainer))
	}
	return container.NewMax(keyboardBg, keyboardContent)
}

func decorateLetterButton(app fyne.App, button *widget.Button, state *AppState) *fyne.Container {
	letter := button.Text
	color, exists := state.letterColors[letter]

	isDark := false
	if ct, ok := app.Settings().Theme().(*customTheme); ok {
		isDark = ct.isDark
	}

	var bgColor c.Color
	var fgColor c.Color

	if exists {
		fgColor = white // Used letters always have white text
		switch color {
		case Green:
			bgColor = green
		case Yellow:
			bgColor = yellow
		case Grey:
			bgColor = grey
		}
	} else {
		if isDark {
			bgColor = darkGrey
			fgColor = lightGrey // Directly use lightGrey to match entered letters
		} else {
			bgColor = lightGrey
			fgColor = black
		}
	}

	return decorateButton(app, button, bgColor, fgColor, fyne.NewSize(43, 58))
}

func decorateButton(app fyne.App, button *widget.Button, bgColor c.Color, fgColor c.Color, size fyne.Size) *fyne.Container {
	isDark := false
	if ct, ok := app.Settings().Theme().(*customTheme); ok {
		isDark = ct.isDark
	}

	if bgColor == nil {
		if isDark {
			bgColor = darkGrey
		} else {
			bgColor = lightGrey
		}
	}
	if fgColor == nil {
		if isDark {
			fgColor = lightGrey
		} else {
			fgColor = black
		}
	}

	bg := canvas.NewRectangle(bgColor)
	bg.SetMinSize(size)
	border := canvas.NewRectangle(c.Transparent)
	border.StrokeWidth = 2.0
	border.StrokeColor = darkGrey

	// Use LowImportance for both light and dark modes to ensure consistent coloring
	button.Importance = widget.LowImportance

	return container.NewStack(
		bg,
		button,
		border,
	)
}
