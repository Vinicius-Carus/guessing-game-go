package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	wordLetters := getWordLetters()

	for {
		fmt.Println("Welcome to the guessing game!!")
		fmt.Print("Let's start? [Y] - Yes / [N] - No: ")
	
		var start string
	
		fmt.Scanln(&start)
	
		start = strings.TrimSpace(strings.ToUpper(start))

		if start == "Y" {
			clearScreen()
			playGame(wordLetters)
			return
		}

		if start == "N" {
			fmt.Println("Ok!! Bye :(")
			return
		}

		clearScreen()
		fmt.Println("Please only 'Y' and 'N' answers")
	}
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type Words struct {
	Words []string `json:"words"`
}

func readJsonFile() Words {
	jsonFile, err := os.Open("words.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var words Words

	json.Unmarshal(byteValue, &words)

	return words
}

func getWordLetters() []string {
	words := readJsonFile()

	wordIndex := rand.Intn(len(words.Words))

	chosenWord := words.Words[wordIndex] 

	splittedWord := strings.Split(chosenWord, "")

	return splittedWord
}

func feedbackMessage(isRight bool) {
	if isRight {
		fmt.Println("Nice!!")
	} else {
		fmt.Println("Oops -1 life")
	}
}

func displayLetters(chosenLetters []string, wordLetters []string) {
	var isChosenLetter bool

	for _, letter := range wordLetters {

		for _, chosenLetter := range chosenLetters {
			isChosenLetter = letter == chosenLetter

			if isChosenLetter { break }
		}

		if isChosenLetter {
			fmt.Print(letter)
		} else {
			fmt.Print("_")
		}

		fmt.Print(" ")
	}

	fmt.Println()
}

func checkIfLetterWasChosenBefore(letter string, chosenLetters []string) bool {
	for _, chosenLetter := range chosenLetters {
		if chosenLetter == letter {
			return true
		}
	}

	return false
}

func checkIfChosenRight(letter string, wordLetters []string) bool {
	for _, wordLetter := range wordLetters {
		if wordLetter == letter {
			return true
		}
	}

	return false
}

func checkIfWordComplete(chosenLetters []string, wordLetters []string) bool {
	for _, letter := range wordLetters {
		var hasTheLetter = false

		for _, chosenLetter := range chosenLetters {
			hasTheLetter = chosenLetter == letter

			if hasTheLetter {
				break
			}
		}

		if !hasTheLetter {
			return	false
		}
	}

	return true
}

func playGame(wordLetters []string) {
	chosenLettersSlice := make([]string, 0, 2)

	var chosenLetter string

	continueAsking := true

	var lives = 3

	fmt.Println()

	fmt.Println("The Word is")
	displayLetters(chosenLettersSlice, wordLetters)

	for continueAsking && lives != 0 {
		fmt.Println()
		fmt.Printf("You have %d/3 lives\n", lives)
		fmt.Print("Choose a letter: ")
		fmt.Scan(&chosenLetter)
		chosenLetter = strings.TrimSpace(strings.ToLower(chosenLetter))

		if len(chosenLetter) == 1 {
			if !checkIfLetterWasChosenBefore(chosenLetter, chosenLettersSlice) {
				chosenLettersSlice = append(chosenLettersSlice, chosenLetter)

				isRight := checkIfChosenRight(chosenLetter, wordLetters)

				displayLetters(chosenLettersSlice, wordLetters)

				feedbackMessage(isRight)

				if !isRight {
					lives--
				}

				isWordComplete := checkIfWordComplete(chosenLettersSlice, wordLetters)

				if isWordComplete {
					fmt.Println("Congratulations you got it right!!")
					continueAsking = false
				}

			} else {
				fmt.Println("You already chose this letter!")
			}

		} else {
			fmt.Println()
			fmt.Println("One letter only!")
			fmt.Println()
		}
	}

	if lives == 0 {
		fmt.Println()
		fmt.Println("Oh no you are dead! :(")
		fmt.Printf("The word was %s!", strings.Join(wordLetters, ""))
	}
}

