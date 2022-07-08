package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Implement a number-guessing game in which the computer computes a four digit number as a secret number
// and a player tries to guess that number correctly. Player would enter her guess and the computer would
// produce a feedback on the positions of the digits. Four-digit number can't start with 0 and have
// repeating digits. Let's say the computer computes 2658 as a secret number to be guessed by the player.
// When player enters her guess such as 1234, the computer would display -1 meaning that only one digit of 1234
// exist in the secret number and its position is wrong. When the player enters 5678 the similarly the computer
// displays +2-1. And the game goes on until the player correctly guess the secret number and the computer
// displays +4. The game also   keeps track of all guesses entered by the players so far and lists them
// when it displays its feedback to the player so that the player can compute her next guess correctly.

var secretNumber int

func init() {
	// Initialize the secret number to be guessed and it should create a different secret number each time the program is run
	rand.Seed(time.Now().UnixNano())
	secretNumber = rand.Intn(9999)

	// Print the secret number to the console
	fmt.Println("Secret number: ", secretNumber)

}

func main() {

	fmt.Println("Welcome to the number guessing game!")
	fmt.Println("Rules:")
	fmt.Println("1. The secret number is a four digit number that cannot start with 0 and can have repeating digits.")
	fmt.Println("2. The player enters a four digit number and the computer displays the feedback on the positions of the digits.")
	fmt.Println("3. The game keeps track of all guesses entered by the players so far and lists them when it displays its feedback to the player.")
	fmt.Println("4. The game ends when the player correctly guesses the secret number and the computer displays +4.")
	fmt.Println("5. Computer displays -1 if only one digit of the player's guess exists in the secret number and its position is wrong.")
	fmt.Println("6. Computer displays +1 if only one digit of the player's guess exists in the secret number and its position is correct.")
	fmt.Println("7. Computer displays -2 if two digits of the player's guess exist in the secret number and their positions are wrong.")
	fmt.Println("8. Computer displays +2 if two digits of the player's guess exist in the secret number and their positions are correct.")
	fmt.Println("9. Computer displays 0 if no digit of the player's guess exists in the secret number.")
	fmt.Println("Note that if you enter repeated digits, the computer will delete them before computing the feedback.")
	fmt.Println("For example; if you enter 1111 computer will compute this guess as 1000")

	// Ask the player to guess the secret number
	var guess int

	// Variable to keep track of all guesses
	var guesses []int

	// Get the player's guess until the player guesses the secret number
	for {
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scan(&guess)
		if err != nil {
			panic(err)
		}
		// check if guess is valid
		if guess < 1000 || guess > 9999 {
			fmt.Println("Invalid guess! Guess must be four digits long and can't start with 0!")
			continue
		}

		// Keep track of all guesses entered by the player
		guesses = append(guesses, guess)

		// Display the feedback to the player
		fmt.Println(feedback(guess))

		// Display all guesses entered by the player
		fmt.Println("All guesses so far: ", guesses)

		if guess == secretNumber {
			break
		}

	}
}

func feedback(guess int) string {

	feedback := ""
	correctPos := 0
	wrongPosButConsist := 0
	wrong := 0
	deleted := 0

	// Check if guess has repeating digits
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			if guess%10 == guess/10%10 {
				// Delete the repeating digit, so that it doesn't affect the feedback
				guess = guess / 10
				deleted++

			}
		}
	}
	// add guess 0 as much as deleted digits
	for i := 0; i < deleted; i++ {
		guess = guess * 10
	}

	// Check if player's guess has one of the digits in the secret number but in the wrong position
	for i := 0; i < 4; i++ {
		// Get the digit at the i-th position of the player's guess
		digit := guess / int(math.Pow10(i)) % 10

		for j := 0; j < 4; j++ {
			// Get the digit at the j-th position of the secret number
			secretDigit := secretNumber / int(math.Pow10(j)) % 10

			// Check if the digit at the i-th position of the player's guess is the same as the digit at the j-th position of the secret number
			if digit == secretDigit {
				// Check if the digit at the i-th position of the player's guess is in the right position
				if i == j {
					correctPos += +1
				} else {
					wrongPosButConsist += -1
				}
			}
			// check if any digit in the player's guess is not in the secret number
			if digit != secretDigit {
				wrong += 1
			}
		}

	}

	// Add the feedback to the feedback string if not wrong
	if correctPos != 0 {
		feedback += fmt.Sprintf("+%d", correctPos)
	}
	if wrongPosButConsist != 0 {
		feedback += fmt.Sprintf("%d", wrongPosButConsist)
	}
	if wrong != 0 && correctPos == 0 && wrongPosButConsist == 0 {
		feedback = fmt.Sprintf("%d", 0)
		return feedback
	}

	return feedback

}
