package main

import "fmt"

func main() {
	i := 1
	//Looping from 1 to 100
	//for loop
	for i <= 100 {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")

		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizzbuzz")
		} else {
			fmt.Println(i)
		}

		i = i + 1
	}
}
