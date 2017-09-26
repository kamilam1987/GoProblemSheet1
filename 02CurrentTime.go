//Author : Kamila Michel(g00340498)
//Source adapted from: https://tour.golang.org/welcome/4
//Program that prints current time and date.

package main

import (
	"fmt"
	"time"
)

func main() {
	//Title
	fmt.Println("This is a current time and date!")

	//Prints current date and time to the console
	fmt.Println("The date and time : ", time.Now())

} //End of main function
