package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var secretNumber int

func init() {
	GenerateSecretNumber()
}

func main() {
	// run the game
	guessGame()
}

func printRules() {
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
}

func GenerateSecretNumber() {
	rand.Seed(time.Now().UnixNano())

	// generate a random number between 1000 and 9999 (inclusive) without repeating digits
	secretNumber = rand.Intn(9000) + 1000

	for {
		// check if secretNum is valid
		if !IsValid(secretNumber) {
			// if not, generate a new number
			secretNumber = rand.Intn(9000) + 1000
		}

		if HasRepeatingNum(secretNumber) {
			secretNumber = rand.Intn(9000) + 1000
		} else {
			break
		}

	}

}

func guessGame() {
	// Ask the player to guess the secret number
	var guess int

	// Variable to keep track of all guesses
	var guesses []int

	// Get the player's input for what to do
	choice := ""

	for {

		fmt.Println("Secret Number: ", secretNumber)
		fmt.Println("Welcome to the number guessing game!")
		fmt.Println("If you want to start, type 'start'.")
		fmt.Println("If you want to see the rules, type 'rules'.")
		fmt.Println("If you want to quit, type 'quit'.")
		fmt.Println("If you want to see the secret number, type 'secret'.")

		fmt.Scan(&choice)

		if choice == "rules" {
			printRules()
		} else if choice == "quit" {
			fmt.Println("Goodbye!")
			return
		} else if choice == "secret" {
			fmt.Println("Secret number: ", secretNumber)
		} else if choice == "start" {
			fmt.Println("Let's start!")
			// Get the player's guess until the player guesses the secret number
			for {
				fmt.Print("Enter your guess: ")
				_, err := fmt.Scan(&guess)
				if err != nil {
					panic(err)
				}

				// check if guess is valid
				if !IsValid(guess) {
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
					fmt.Println("You won!")
					break
				}

			}
		} else {
			fmt.Println("Invalid input!")
			return
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

func HasRepeatingNum(num int) bool {

	// get digits of num
	firstDigit, secondDigit, thirdDigit, fourthDigit := GetDigits(num)

	// check each digit of secretNum against the other digits
	if firstDigit == secondDigit || firstDigit == thirdDigit || firstDigit == fourthDigit {
		// calculate a new number without repeating digits
		return true
	} else if secondDigit == thirdDigit || secondDigit == fourthDigit {
		return true
	} else if thirdDigit == fourthDigit {
		return true
	} else {
		// if secretNum has no repeating digits, break the loop
		return false
	}

}

func IsValid(num int) bool {
	// check if num is valid
	if num < 1000 || num > 10000 {
		return false
	}
	return true

}

func GetDigits(num int) (int, int, int, int) {

	// get first digit of num
	firstDigit := num / int(math.Pow10(0)) % 10

	// get second digit of num
	secondDigit := num / int(math.Pow10(1)) % 10

	// get third digit of num
	thirdDigit := num / int(math.Pow10(2)) % 10

	// get fourth digit of num
	fourthDigit := num / int(math.Pow10(3)) % 10

	return firstDigit, secondDigit, thirdDigit, fourthDigit
}
