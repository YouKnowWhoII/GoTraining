package main

import (
	"fmt"
	"math/rand"
	"time"
)

func playGuessingGame() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
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
