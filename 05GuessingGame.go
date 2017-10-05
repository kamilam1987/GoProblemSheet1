//Author : Kamila Michel(g00340498)
//Title: Guessing game
//Guessing game where the user must guess a randomly generated number. After every guess the program tells the user whether their number was too high or too low. At the end, the number of tries should be printed. It counts only as one try if they input the same number multiple times consecutively.
//Source code adapted from:
//http://ispycode.com/GO/Flow-Control/Examples/High-Low-Guessing-Game
//http://golangcookbook.blogspot.ie/2012/11/guess-number-game-in-golang.html

//Main package
package main

//Import format, math random and time
import (
	"fmt"
	"math/rand"
	"time"
)

//Main function
func main() {

	var guess int
	number := numRand(1, 50)
	//Set number of guesses to 5 as a constatnt
	const numGuesses = 5
	//Number of guesses taken by user
	guessesTaken := 0

	//For loop is used as while loop. Loop will execute untill condition is true
	for {
		if guessesTaken != numGuesses {

			//fmt.Println("Number random: ", number)
			//Asking user to pick a number from range 1-50
			fmt.Println("Please pick number from 1-50: ")
			fmt.Scanln(&guess)
			//Counter
			guessesTaken++

			//Checking if guess number is smaller then number
			if guess < number {
				fmt.Println("Sorry your guess is too low !\n")
				//Checking if guess number is higher then numer
			} else if guess > number {
				fmt.Println("Sorry your guess is too high\n")
				//Checking if guess is correct
			} else {
				//Prints out if number is correct and says how many tries user used
				fmt.Printf("Congrats! You won! It took you %d number of times\n", guessesTaken)
				break
			}
		} else {
			//Prints out when all guesses where incorrect
			fmt.Printf("The winning number was %d and you reached number of tries\n\n", number)
			break
		}

	} //End of for loop

} //End of main function

//numRand function generates random number between given range
func numRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
} //End of numRand function
