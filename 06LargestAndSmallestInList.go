//Author : Kamila Michel(g00340498)
//Title: Largest and smallest in list
//Go program that returns the largest and smallest elements in a list.
//Source code adapted from: https://www.socketloop.com/tutorials/golang-find-biggest-largest-number-in-array

//Main package
package main 

//import format
import 
	"fmt"

//Main function
func main () {
	//Declare array of integers
	array := []int64{6, 8, 100, 15, 45, 69, 68}

	maxNumber := array[0]
	minNumber := array[0]

	//Looping array
	for _, number := range array {
		if number > maxNumber {
			maxNumber = number
		}else if number < minNumber {
			minNumber = number

		}

	}//End of for Loop

	//Outpu the largest and smalest number from array
	fmt.Println("The max number is:", maxNumber, "and min number is:", minNumber)

}//End of main function