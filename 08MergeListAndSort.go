//Author : Kamila Michel(g00340498)
//Title: Merge list and sort
//Write a function that merges two sorted lists into a new sorted list
//Source code adapted from: 

package main

import 
	"fmt"
	
//Main function
func main () {
	list1 := []int{1, 4, 6}
	list2 := []int{2, 3, 5}
	
	//Print new array
	fmt.Println(merge(list1,list2))

}//End of main function

//Merge funcion
func merge(list1 [] int, list2 [] int) []int {
	
		size, i, j := len(list1)+len(list2), 0, 0
		//Result which takes size of two arrays and make a new array
		result := make([]int, size, size)
	
		for k := 0; k < size; k++ {
			if i > len(list1)-1 && j <= len(list2)-1 {
				result[k] = list2[j]
				j++
			} else if j > len(list2)-1 && i <= len(list1)-1 {
				result[k] = list1[i]
				i++
			} else if list1[i] < list2[j] {
				result[k] = list1[i]
				i++
			} else {
				result[k] = list2[j]
				j++
			}
		}
		return result
	}