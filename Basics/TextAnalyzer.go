package main

import (
	"fmt"
	"strings"
)

func analyzeText() {
	var text string
	fmt.Print("Enter a text string: ")
	fmt.Scanln(&text)

	//strings.Fields() splits the text into words based on spaces
	//returns a slice of strings
	words := strings.Fields(text)
	numWords := len(words)
	numChars := len(text)

	fmt.Printf("Number of words: %d\n", numWords)
	fmt.Printf("Number of characters: %d\n", numChars)

	print("Do you want to count a specific character? (y/n): ")
	var choice string
	fmt.Scanln(&choice)
	if choice == "y" {
		countSpecificCharacter(text)
	}
}

func countSpecificCharacter(text string) {
	var letter string
	fmt.Println("Enter a character to count: ")
	fmt.Scanln(&letter)
	count := strings.Count(text, letter)

	fmt.Printf("Number of '%s': %d\n", letter, count)
}
