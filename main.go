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

	indexWord := rand.Intn(len(words.Words))

	choosedWord := words.Words[indexWord] 

	splitedWord := strings.Split(choosedWord, "")

	return splitedWord
}

func feedbackMessage(isRight bool) {
	if isRight {
		fmt.Println("Nice!!")
	} else {
		fmt.Println("Oops -1 life")
	}
}

func displayLetters(choosedLetters []string, wordLetters []string) {
	var isChoosedLetter bool


	for _, letter := range wordLetters {

		for _, choosedLetter := range choosedLetters {
			isChoosedLetter = letter == choosedLetter

			if isChoosedLetter { break }
		}

		if isChoosedLetter {
			fmt.Print(letter)
		} else {
			fmt.Print("_")
		}

		fmt.Print(" ")
	}

	fmt.Println()

}

func checkIfLetterWasChoosedBefore(letter string, choosedLetters []string) bool {

	for _, letterChoosed := range choosedLetters {
		if letterChoosed == letter {
			return true
		}
	}

	return false
}

func checkIfChoosedRight(letter string, wordLetters []string) bool {
	for _, wordLetter := range wordLetters {
		if wordLetter == letter {
			return true
		}
	}

	return false
}

func checkIfCompletedWord(choosedLetters []string, wordLetters []string) bool {
	for _, letter := range wordLetters {
		var hasTheLetter = false

		for _, choosedLetter := range choosedLetters {
			hasTheLetter = choosedLetter == letter

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

	choosedLettersSlice := make([]string, 0, 2)

	var choosedLetter string

	constinueAsking := true

	var lifes = 3

	fmt.Println()

	fmt.Println("The Word is")
	displayLetters(choosedLettersSlice, wordLetters)


	for constinueAsking && lifes != 0 {
		fmt.Println()
		fmt.Printf("You have %d/3 lifes\n", lifes)
		fmt.Print("Choose a letter: ")
		fmt.Scan(&choosedLetter)
		choosedLetter = strings.TrimSpace(strings.ToLower(choosedLetter))

		if len(choosedLetter) == 1 {
			if !checkIfLetterWasChoosedBefore(choosedLetter, choosedLettersSlice) {
				choosedLettersSlice = append(choosedLettersSlice, choosedLetter)

				isRight := checkIfChoosedRight(choosedLetter, wordLetters)

				displayLetters(choosedLettersSlice, wordLetters)

				feedbackMessage(isRight)

				if !isRight {
					lifes--
				}

				isCompleteWord := checkIfCompletedWord(choosedLettersSlice, wordLetters)

				if isCompleteWord {
					fmt.Println("Congratulations you got it right!!")
					constinueAsking = false
				}

			} else {
				fmt.Println("You already choosed this letter!")
			}

		} else {
			fmt.Println()
			fmt.Println("One letter only!")
			fmt.Println()
		}

	}

	if lifes == 0 {
		fmt.Println()
		fmt.Println("Oh no you are dead! :(")
		fmt.Printf("The word was %s!", strings.Join(wordLetters, ""))
	}

}

