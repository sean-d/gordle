# Gordle – Yet another Wordle clone

- The original Gordle was created by [scastiel](https://github.com/scastiel/gordle).
- andydotxyz forked the above [here](https://github.com/andydotxyz/gordle) and updated the version of Fyne, added some nice API updates, and addressed the vendor folder.
- I forked andydotxyz's version and created word themes (e.g. "sports", "food", "colors", "animals", etc.), and added a dark/light mode selection as per scastiel's original wishlist. I also updated the readme and help docs to reflect the changes made.

![Gordle Demo](assets/GordleDemo.gif)

## Features

- [x] Works on arm64/x64 macOS; should work on Linux and Windows as well (not tested)
- [x] Classic wordle gameplay
- [x] Select different word lists by theme
- [x] Dark and light theme selection

## Prerequisites

To compile and run Gordle, you will need to have Go installed on your machine. Refer to the [installation instructions](https://go.dev/doc/install).

## Install

```shell
go install github.com/sean-d/gordle@latest
```

## Run

```shell
gordle
```

## Package

Gordle can be packaged for your operating system using [Fyne](https://fyne.io/). You’ll need to install it first:

```shell
go install fyne.io/fyne/v2/cmd/fyne@latest
```

Then clone this repository, and run the following command:

```shell
fyne package
```

A binary for your OS will be generated in the same directory.

## License

[MIT, see LICENSE](LICENSE)
