//Author : Kamila Michel(g00340498)
//Title: Palindrome test
//Program that tests whether a string is a palindrome.
//Source code adapted from: http://www.golangpro.com/2016/01/check-string-palindrome-go.html

//Main package
package main 

//import format and string
import (
	"fmt"
	"strings"
)
//Main function
func main() {
	var world string
	fmt.Println("Please enter string:")
	fmt.Scanf("%s",&world)
	world = strings.ToLower(world)
	fmt.Println(palindrome(world))

}//End of main function

//Function which checks if string is a palindrome
func palindrome(s string) string {

	mid := len(s) / 2
	last := len(s) - 1

	//Checking if it's palindrome
	for i :=0; i < mid; i++ {
		if s[i] != s[last-i] {
			return "It's not a Palimdrome!"
			} 
		}//End of for loop
		
		return "It's a Palindrome!"
}//End of palindrome function



	
