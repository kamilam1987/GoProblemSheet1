//Author : Kamila Michel(g00340498)
//Factorial digit sum
//Program which calculates the sum of the digits of the factorial of a number
//Source code adapted from:
// https://stackoverflow.com/questions/17564335/golang-math-big-what-is-the-max-value-of-big-int
//https://golang.org/pkg/math/big/
//https://play.golang.org/
//https://stackoverflow.com/questions/46395819/get-sum-of-bigint-number-golang

//Main package
package main

//Import format and math/big which is used to calculate big integers
import (
	"fmt"
	"math/big"
)

//Main function
func main() {
	//Declate variable
	var number int64

	//Asking user to input number
	fmt.Println("Please enter number from 1-100:")
	//User input
	fmt.Scanln(&number)
	//Prints out factorial
	fmt.Printf("Factorial of %d is: %d\n", number, factorial(number))
	//Prints out sum of digits in factorial
	fmt.Printf("Sum of digits in factorial is: %d ", sumOfDigits(factorial(number)))

} //End of main function

//Function which calculate factorial from given number by user and converts it into big integer
func factorial(number int64) *big.Int {
	//if number smaller then 0 return 1
	if number < 0 {
		return big.NewInt(1)
		//else if number is equal  to 0 return 1
	} else if number == 0 {
		return big.NewInt(1)
	}
	bigN := big.NewInt(number)
	//return formula to calculate factorial
	return bigN.Mul(bigN, factorial(number-1))
}

//Function that calculate sum of the digits from factorial
func sumOfDigits(number *big.Int) *big.Int {
	//Declare variables using big.NewInt for large numbers
	ten := big.NewInt(10)
	sum := big.NewInt(0)
	mod := big.NewInt(0)
	for ten.Cmp(number) < 0 {
		sum.Add(sum, mod.Mod(number, ten))
		number.Div(number, ten)
	}
	sum.Add(sum, number)
	//Return sum of digits from factorial
	return sum
} //End of sumOfDigits function
