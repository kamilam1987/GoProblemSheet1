//Author : Kamila Michel(g00340498)
//Title: Reverse string
//Program that reverse a string in Go.
//Sourcecode adapted from: 

//Main package
package main 

//import format
import 
	"fmt"

//Reverse function - reverse string
func reverse(s string) string {
	characters := []rune(s)

	//Looping array and checking lenght of string
	for i, j := 0, len(characters)-1; i < j; i, j = i + 1, j - 1 {
		characters[i], characters[j] = characters[j], characters[i]
		}//End of foor loop
		return string(characters)
}//End of reverse function

//Main function
func main(){
	//Prints reverse world
	fmt.Println(reverse("kamila"))
}//End of main function




