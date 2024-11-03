package main

import (
	"fmt"
	"strings"
)

 // func getCategory() string {
	fmt.Println("Choose a category: animals, countries, programming")
	var category string
	fmt.Scanln(&category)
 // fmt.Scanln(&category)	return strings.ToLower(category)
}

func getPlayerGuess() rune {
	var guess string
	for {
		fmt.Print("Guess a letter: ")
		fmt.Scanln(&guess)

		if len(guess) != 1 || !isLetter(guess) {
			fmt.Println("Please enter a valid single letter.")
			continue
		}
		return rune(strings.ToLower(guess)[0])
	}
}

func isLetter(s string) bool {
	return len(s) == 1 && s[0] >= 'a' && s[0] <= 'z'
}

