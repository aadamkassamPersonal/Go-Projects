package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var guess int
	target := rand.Intn(100) + 1
	attempts := 0
	right := false

	for attempts <= 10 {
		if right == true {
			break
		}

		if attempts == 9 {
			fmt.Println("Last Guess!")
		}

		fmt.Print("Guess a number between 1 and 100: ")
		fmt.Scanf("%d", &guess)
		attempts++
		right = guesstimate(guess, target)

	}

	if right {
		fmt.Printf("\nTarget was %v ", target)
		fmt.Printf("\nYou Guessed it in %v ", attempts)
	} else {
		fmt.Printf("\nTarget was %v, Better luck next time!", target)
	}
}

func guesstimate(guess int, target int) bool {

	if guess < target {
		fmt.Println("Too low")
	} else if guess > target {
		fmt.Println("Too high")
	} else {
		return true // Correct guess!
	}

	return false

}
