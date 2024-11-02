package main



import (

	"math/rand"

)



var wordCategories = map[string][]string{

	"animals":     {"dog", "cat", "elephant", "giraffe"},

	"countries":   {"canada", "brazil", "france", "japan"},

	"programming": {"golang", "python", "javascript", "java"},

}



func selectRandomWord(category string) string {

	words, exists := wordCategories[category]

	if !exists {

		return ""

	}

	return words[rand.Intn(len(words))]

}



func displayWord(word string, guessedLetters []rune) {

	for _, letter := range word {

		if contains(guessedLetters, letter) {

			print(string(letter), " ")

		} else {

			print("_ ")

		}

	}

	println()

}



func displayGuessedLetters(guessedLetters []rune) {

	print("Guessed letters: ")

	for _, letter := range guessedLetters {

		print(string(letter), " ")

	}

	println()

}



func isWordGuessed(word string, guessedLetters []rune) bool {

	for _, letter := range word {

		if !contains(guessedLetters, letter) {

			return false

		}

	}

	return true

}



