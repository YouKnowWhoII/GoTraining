package main

import (
	"fmt"
	"strings"
)

func analyzeText() {
	var text string
	fmt.Print("Enter a text string: ")
	fmt.Scanln(&text)

	words := strings.Fields(text)
	numWords := len(words)
	numChars := len(text)

	fmt.Printf("Number of words: %d\n", numWords)
	fmt.Printf("Number of characters: %d\n", numChars)

	var letter string
	fmt.Print("Enter a letter to count: ")
	fmt.Scanln(&letter)
	count := strings.Count(text, letter)

	fmt.Printf("Number of '%s': %d\n", letter, count)
}
