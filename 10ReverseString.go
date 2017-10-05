//Author : Kamila Michel(g00340498)
//Title: Reverse string
//Program that reverse a string in Go.
//Sourcecode adapted from: http://golangcookbook.com/chapters/strings/reverse/

//Main package
package main 

//import format
import 
	"fmt"

//Main function
func main() {
	//Declare vriable
	var s string
	//Ask user to input string
	fmt.Println("Please input your word to reverse: ")
	fmt.Scanf("%s",&s)
	//Prints out reversed string
	fmt.Println("Your reverse world is:",reverse(s))
	//Prints reversed world
	//fmt.Println(reverse("kamila"))

}//End of main function
	
//Reverse function - reverse string
func reverse(s string) string {
	characters := []rune(s)

	//Looping array and checking lenght of string
	for i, j := 0, len(characters)-1; i < j; i, j = i + 1, j - 1 {
		characters[i], characters[j] = characters[j], characters[i]
		}//End of foor loop

		//Return string
		return string(characters)
}//End of reverse function




