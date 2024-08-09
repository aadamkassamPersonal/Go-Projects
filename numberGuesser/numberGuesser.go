package main

import (
	"fmt"
	"math/rand"
)

func main() {
	target := rand.Intn(100) + 1

	var guess int
	attempts := 0
	right := false

	for {

		if attempts == 9 {
			fmt.Println("Last Guess!")
		}

		fmt.Print("Guess a number between 1 and 100: ")
		fmt.Scanf("%d", &guess)

		attempts++

		if guess < target {
			guess = 0
			fmt.Println("Too low")
		} else if guess > target {
			guess = 0
			fmt.Println("Too high")
		} else {
			right = true
			break // Correct guess!
		}

		if attempts == 10 {
			break
		}
	}

	if right {
		fmt.Printf("\nTarget was %v ", target)
		fmt.Printf("\nYou Guessed it in %v ", attempts)
	} else {
		fmt.Printf("\nTarget was %v, Better luck next time!", target)
	}
}
