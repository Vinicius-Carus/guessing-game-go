# Guessing Game

A simple word-guessing game written in Go.

## Overview

The game randomly picks a word from a JSON file (`words.json`), and you must guess it one letter at a time. You have 3 lives and lose one each time you guess an incorrect letter. The game ends when you either guess the entire word or run out of lives.

## Prerequisites

- Go 1.16 or later
- Windows, macOS, or Linux

## File Structure

```plaintext
├── main.go     # Game source code
└── words.json  # JSON file containing the word list
```

### Sample `words.json`

```json
{
  "words": [
    "banana",
    "apple",
    "computer",
    "gopher",
    "programming"
  ]
}
```

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/guessing-game.git
   cd guessing-game
   ```

## Running the Game

### Without building

```bash
go run main.go
```

### Building and running

```bash
go build -o guessing-game main.go
./guessing-game    # Windows: guessing-game.exe
```

## How to Play

1. At the prompt, type `Y` to start or `N` to exit.
2. You'll see `_` for each undiscovered letter with spaces between them.
3. Enter one letter at a time and press Enter.
4. You have 3 lives; each wrong guess costs one life.
5. The game ends when you guess the word or lose all your lives.

## Clear Screen Compatibility

The screen-clearing command works on Windows (`cls`) and Unix-like systems (`clear`). If it doesn’t work, you can ignore it or run the game directly in your terminal

