//Author : Kamila Michel(g00340498)
//Source adapted from: http://wiki.c2.com/?FizzBuzzTest
//Go Program which prints the numbers from 1 to 100, each on a new line, to the console, except for the following conditions. For multiples of three print Fizz instead of the number, and for multiples of five print Buzz. For numbers that are multiples of both three and five print FizzBuzz.

package main

import "fmt"

func main() {

	//Loopin number from 1 to 100
	for i := 1; i <= 100; i++ {
		//if numbers that are multiples of both 3 and 5 prints FizzBuzz
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			// else if number multiples of five print Buzz
		} else if i%5 == 0 {
			fmt.Println("Buzz")
			//else if number multiples of five print fizz
		} else if i%3 == 0 {
			fmt.Println("Fizz")
			//else prints number
		} else {
			fmt.Println(i)
		}
	} //end for
} //end main
