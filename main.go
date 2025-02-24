package main

import (
	"math/rand"
	"os"
	"time"

	gordle "github.com/sean-d/gordle/src"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	os.Setenv("FYNE_THEME", "light")
	gordle.StartFyneGame()
}
