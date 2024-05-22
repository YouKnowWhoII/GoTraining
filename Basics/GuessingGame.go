package main

import (
	"fmt"
	"math/rand"
	"time"
)

func playGuessingGame() {
	// Create a new random source with the current time as the seed
	// Ensures that the random number generated is different each time the program is run
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	// +1 shifts the range from [0, 99] to [1, 100]
	randomNumber := r.Intn(100) + 1
	var guess int

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess < randomNumber {
			fmt.Println("Too low!")
		} else if guess > randomNumber {
			fmt.Println("Too high!")
		} else {
			fmt.Println("You guessed right!")
			break
		}
	}
}
