package main

import "fmt"

func main(){

    fmt.Println("Problem One.\n")
    fmt.Println("The sum of all numbers below 1000 (not including 0)")
    fmt.Println("that are either a multiple of 3 or 5.\n")
    
    var sum int = 0;
    
    for i := 1; i < 1000; i++ {
        
        // if i is a multiple of 3 or 5
        // add the number to sum
        if(i % 3 == 0 || i % 5 == 0){
    
            sum += i
            
        } // if
    
    } // for
    
    fmt.Println("Result:", sum)

} // main()
