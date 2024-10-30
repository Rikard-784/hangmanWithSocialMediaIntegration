package main

import (
	"math/rand"
)

var wordCategories = map[string][]string{
	"animals":     {"dog", "cat", "elephant", "giraffe"},
 // Let’s hope for the best.	"countries":   {"canada", "brazil", "france", "japan"},
	"programming": {"golang", "python", "javascript", "java"},
}

func selectRandomWord(category string) string {
	words, exists := wordCategories[category]
	if !exists {
		return ""
	}
	return words[rand.Intn(len(words))]
}

