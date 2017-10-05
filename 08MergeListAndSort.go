//Author : Kamila Michel(g00340498)
//Title: Merge list and sort
//Write a function that merges two sorted lists into a new sorted list
//Source code adapted from: https://stackoverflow.com/questions/5958169/how-to-merge-two-sorted-arrays-into-a-sorted-array, 

//Main package
package main

//import format
import 
	"fmt"
	
//Main function
func main () {

	//Declare two sorted arrays
	list1 := []int{1, 4, 6}
	list2 := []int{2, 3, 5}
	
	//Print new array
	fmt.Println(merge(list1,list2))

}//End of main function

//Merge funcion
func merge(list1 [] int, list2 [] int) []int {
	
	//k - counter for merged list, i - counter for list1, j - counter for list2
		size, i, j := len(list1)+len(list2), 0, 0
		//Make a new sorted list which contains list1 and list2
		mergedList := make([]int, size, size)
	
		//Checking if both counters are valid, if only list1's counter are valid or if only list2's counter are valid and resizing the mergedList
		for k := 0; k < size; k++ {
	
			if i > len(list1)-1 && j <= len(list2)-1 {
				mergedList[k] = list2[j]
				j++
			} else if j > len(list2)-1 && i <= len(list1)-1 {
				mergedList[k] = list1[i]
				i++
			
			} else if list1[i] < list2[j] {
				mergedList[k] = list1[i]
				i++
			} else {
				mergedList[k] = list2[j]
				j++
			}
		}
		//Return sorted array
		return mergedList
	}