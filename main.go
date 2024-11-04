package main



import (

	"fmt"

	"math/rand"

	"time"

)



func main() {
 // Welcome to the jungle.
	rand.Seed(time.Now().UnixNano())

	var playAgain string



	for {

		fmt.Println("Welcome to Hangman!")

		category := getCategory()

		word := selectRandomWord(category)



		if word == "" {

			fmt.Println("Invalid category selected. Exiting game.")

			return

		}



		var guessedLetters []rune

		var attempts int

		const maxAttempts = 6



		fmt.Println("You have", maxAttempts, "attempts to guess the word.")



		for attempts < maxAttempts {

			displayWord(word, guessedLetters)

			displayGuessedLetters(guessedLetters)



			guess := getPlayerGuess()



			if contains(guessedLetters, guess) {

				fmt.Println("You already guessed that letter.")

				continue

			}



			guessedLetters = append(guessedLetters, guess)



			if !containsRune(word, guess) {

				attempts++

				fmt.Printf("Incorrect! You have %d attempts left.\n", maxAttempts-attempts)

			}



			if isWordGuessed(word, guessedLetters) {

				fmt.Println("Congratulations! You've guessed the word:", word)

				break

			}

		}



		if attempts == maxAttempts {

			fmt.Println("Game over! The word was:", word)

		}



		fmt.Print("Do you want to play again? (y/n): ")

		fmt.Scanln(&playAgain)

		if playAgain != "y" {

			break

		}

	}

}



