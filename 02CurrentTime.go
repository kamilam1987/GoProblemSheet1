//Author : Kamila Michel(g00340498)
//Title: Current time
//Program that prints current time and date.
//Source adapted from: https://tour.golang.org/welcome/4

//Main package
package main

//import format and time
import (
	"fmt"
	"time"
)

//Main function
func main() {
	//Title
	fmt.Println("This is a current time and date!")

	//Prints current date and time to the console
	fmt.Println("The date and time : ", time.Now())

} //End of main function
