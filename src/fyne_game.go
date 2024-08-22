//go:generate fyne bundle -package gordle -o resources.go assets/AppIcon.svg
//go:generate fyne bundle -package gordle -o resources.go -append assets/AppIcon.png
//go:generate fyne bundle -package gordle -o resources.go -append assets/example.png
//go:generate fyne bundle -package gordle -o resources.go -append assets/about_part1.md
//go:generate fyne bundle -package gordle -o resources.go -append assets/about_part2.md

package gordle

import (
	c "image/color"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var green = c.RGBA{R: 106, G: 170, B: 100, A: 255}
var yellow = c.RGBA{R: 201, G: 180, B: 88, A: 255}
var grey = c.RGBA{R: 120, G: 124, B: 126, A: 255}
var darkGrey = c.RGBA{R: 135, G: 138, B: 140, A: 255}
var lightGrey = c.RGBA{R: 211, G: 215, B: 218, A: 255}
var black = c.Black
var white = c.White

func StartFyneGame() {
	state := NewAppState()

	app := app.New()
	app.SetIcon(resourceAppIconPng)
	window := app.NewWindow("Gordle")
	window.SetFixedSize(true)
	window.SetCloseIntercept(func() {
		app.Quit()
	})
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
	window.ShowAndRun()
}

var errorTicker *time.Ticker

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

func render(app fyne.App, state *AppState, window fyne.Window) {
	space := canvas.NewRectangle(c.Transparent)
	space.SetMinSize(fyne.NewSize(0, 40))

	window.SetContent(container.NewVBox(
		header(app, state),
		statusMessage(state),
		wordRows(state),
		space,
		keyboard(app, state, window)))
}

func header(app fyne.App, state *AppState) *fyne.Container {
	header := container.New(layout.NewVBoxLayout())

	iconRect := canvas.NewRectangle(c.Transparent)
	iconRect.SetMinSize(fyne.NewSize(48, 48))
	icon := widget.NewIcon(resourceAppIconSvg)
	iconBox := container.NewStack(iconRect, icon)

	title := canvas.NewText("Gordle", black)
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24
	title.TextStyle.Bold = true

	helpButtonRect := canvas.NewRectangle(c.Transparent)
	helpButtonRect.SetMinSize(fyne.NewSize(48, 48))
	helpButton := widget.NewButton("?", func() { openAboutDialog(app, state) })
	helpButtonBox := container.NewStack(helpButtonRect, helpButton)

	titleRow := container.New(layout.NewBorderLayout(nil, nil, iconBox, helpButtonBox))
	titleRow.Add(iconBox)
	titleRow.Add(helpButtonBox)
	titleRow.Add(title)

	header.Add(titleRow)

	border := canvas.NewRectangle(lightGrey)
	border.SetMinSize(fyne.NewSize(0, 5))
	header.Add(border)

	return header
}

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

func statusMessage(state *AppState) *fyne.Container {
	message := ""
	if state.game.State == Won {
		message = "You won! Press ENTER to start again."
	} else if state.game.State == Lost {
		message = "The solution was " + state.game.Solution.Letters + ".  Press ENTER to try again."
	} else if state.errorMessage != "" {
		message = state.errorMessage
	}

	blackBox := canvas.NewRectangle(c.Transparent)
	if message != "" {
		blackBox.FillColor = black
	}
	blackBox.SetMinSize(fyne.NewSize(0, 50))
	blackBox.StrokeWidth = 15
	blackBox.StrokeColor = white

	statusText := canvas.NewText(message, white)
	statusText.Alignment = fyne.TextAlignCenter
	statusText.TextSize = 14
	statusText.TextStyle.Bold = true

	return container.NewStack(blackBox, statusText)
}

func wordRows(state *AppState) *fyne.Container {
	rows := container.New(layout.NewVBoxLayout())
	for i, guess := range state.guesses {
		feedback := state.feedbacks[i]
		rows.Add(wordRow(guess, feedback))
	}
	remaining := 6 - len(state.guesses)
	if state.game.State == Started {
		remaining = remaining - 1
		rows.Add(currentWordRow(state.currentWord))
	}
	for i := 0; i < remaining; i++ {
		rows.Add(emptyWordRow())
	}
	return container.New(layout.NewCenterLayout(), rows)
}

func wordRow(word Word, feedback Feedback) *fyne.Container {
	grid := container.New(layout.NewGridLayout(5))
	for i, letter := range strings.Split(word.Letters, "") {
		color := feedback.Colors[i]
		grid.Add(letterBox(letter, color))
	}
	return grid
}

func currentWordRow(word string) *fyne.Container {
	grid := container.New(layout.NewGridLayout(5))
	for _, letter := range strings.Split(word, "") {
		grid.Add(currentWordLetterBox(letter))
	}
	for i := len(word); i < 5; i++ {
		grid.Add(emptyLetterBox())
	}
	return grid
}

func emptyWordRow() *fyne.Container {
	grid := container.New(layout.NewGridLayout(5))
	for i := 0; i < 5; i++ {
		grid.Add(emptyLetterBox())
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

func currentWordLetterBox(letter string) *fyne.Container {
	box := canvas.NewRectangle(white)
	box.SetMinSize(fyne.NewSize(62, 62))
	box.StrokeColor = darkGrey
	box.StrokeWidth = 2.0

	text := canvas.NewText(letter, black)
	text.Alignment = fyne.TextAlignCenter
	text.TextSize = 32
	text.TextStyle.Bold = true
	return container.NewStack(box, text)
}

func emptyLetterBox() *canvas.Rectangle {
	box := canvas.NewRectangle(white)
	box.SetMinSize(fyne.NewSize(62, 62))
	box.StrokeColor = lightGrey
	box.StrokeWidth = 2.0
	return box
}

func keyboard(app fyne.App, state *AppState, window fyne.Window) *fyne.Container {
	letterRows := [][]string{
		{"Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P"},
		{"A", "S", "D", "F", "G", "H", "J", "K", "L"},
		{"Z", "X", "C", "V", "B", "N", "M"},
	}
	keyboard := container.New(layout.NewVBoxLayout())
	for i, letters := range letterRows {
		row := container.New(layout.NewHBoxLayout())
		if i == 2 {
			theState := state
			enterButton := widget.NewButton("ENTER", func() {
				*theState = theState.enter()
				if theState.errorMessage != "" {
					displayError(app, theState, window)
				}
				render(app, theState, window)
			})
			decoratedEnterButton := decorateButton(enterButton, nil, nil, fyne.NewSize(65.4, 58))
			row.Add(decoratedEnterButton)
		}
		for _, letter := range letters {
			theLetter := letter
			theState := state
			button := widget.NewButton(theLetter, func() {
				*theState = theState.typeLetter(theLetter)
				render(app, state, window)
			})
			decoratedButton := decorateLetterButton(button, state)
			row.Add(decoratedButton)
		}
		if i == 2 {
			theState := state
			backButton := widget.NewButton("BACK", func() {
				*theState = theState.backspace()
				render(app, theState, window)
			})
			decoratedBackButton := decorateButton(backButton, nil, nil, fyne.NewSize(65.4, 58))
			row.Add(decoratedBackButton)
		}
		keyboard.Add(container.New(layout.NewCenterLayout(), row))
	}
	return keyboard
}

func decorateLetterButton(button *widget.Button, state *AppState) *fyne.Container {
	letter := button.Text

	var bgColor c.Color
	var fgColor c.Color
	color, exists := state.letterColors[letter]
	if exists {
		fgColor = white
		switch color {
		case Green:
			bgColor = green
		case Yellow:
			bgColor = yellow
		case Grey:
			bgColor = grey
		}
	}

	return decorateButton(button, bgColor, fgColor, fyne.NewSize(43, 58))
}

func decorateButton(button *widget.Button, bgColor c.Color, fgColor c.Color, size fyne.Size) *fyne.Container {
	if bgColor == nil {
		bgColor = lightGrey
	}
	if fgColor == nil {
		fgColor = black
	}
	button.Importance = widget.LowImportance

	bg := canvas.NewRectangle(bgColor)
	bg.SetMinSize(size)
	border := canvas.NewRectangle(c.Transparent)
	border.StrokeWidth = 2.0
	border.StrokeColor = grey
	return container.NewStack(
		bg,
		button,
		border,
	)
}
