//Author : Kamila Michel(g00340498)
//Title: Newton’s method for square roots
//Program which calculate the square root of a number using Newton’s method. Newton’s method is to approximate the square root of a number x by picking a starting point z and then repeating the following operation.
//Source code adapted from:

//Hint: to declare and initialize a floating point value, give it floating-point syntax or use a conversion:
//z := float64(1)
//z := 1.0

//Main package
package main

//Import format
import "fmt"

//Main function
func main() {
	//Declare variables
	//User number
	var x float64

	//Asking user to input number
	fmt.Println("Please enter number")
	//User input
	fmt.Scanln(&x)

	//Output the result
	fmt.Printf("The square root of %.4f is %.4f", x, newtonMethod(x))

} //End of main function

func newtonMethod(x float64) float64 {
	//Declering variables
	z := float64(1)
	var zNext float64

	//For loop  will execute untill condition is true
	for z != zNext {
		//This is newton formula that was given
		zNext = z - ((z*z - x) / (2 * z))
		//Compare z to the value zNext
		z = zNext

		if z == zNext {
			zNext = z - ((z*z - x) / (2 * z))

		}

	} //End of for loop

	//Return z
	return z
} //End of newtonMethod
