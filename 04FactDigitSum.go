//Author : Kamila Michel(g00340498)
//Factorial digit sum
//Program which calculates the sum of the digits of the factorial of a number 
//Source code adapted from:

//Main package
package main

//import format//import format

import (
	"fmt"
	"math/big"
)

 
//Declare variable
// uint64 is the set of all unsigned 64-bit integers.
var fact uint64 = 1 
                      
var i int = 1
var n int
 
//    function declaration        
func factorial(n int) uint64 {   
	//If negative number
    if(n < 0){
        fmt.Print("Factorial of negative number doesn't exist.")    
    }else{        
        for i:=1; i<=n; i++ {
			//types int64 and int
            fact *= uint64(i)  
        }
         
    }    
	//return factorial
    return fact  
}
 
func main(){    
    fmt.Print("Enter a positive integer : ")
    fmt.Scan(&n)   
    fmt.Print("Factorial is: ",factorial(n))
}

