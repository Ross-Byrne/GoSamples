package main

import "fmt"

func main(){

    var result int = 0;
    var primeCount int = 0;
    var count int = 0;
   
    fmt.Println("Problem 2.\n")
    
    fmt.Println("Find the 10,001st Prime Number.\n")
    fmt.Println("Please Wait, Crunching Numbers...")
    
    // loop for 10000 prime numbers, loop breaks after 10001st prime is obtained
    for primeCount < 10001 {
        // check if the count variable is a prime
        if(isPrime(count) == true){
            
            // count the number of primes
            primeCount++
            
            // save the prime
            result = count
            
            // print out the prime number
            //fmt.Println(result)
            
            // increament count
            count++
            
        } else { // if not a prime
        
            // increament count
            count++
            
        } // if
    
    } // while
    
    // print out 10001st prime number
    fmt.Println("\nThe 10001st Prime Number is:", result)

} // main()

// function to find if int is a prime number
// number is prime if only 1 and itself divides into it evenly
func isPrime(num int) bool {
    
    // 0 and 1 are not prime numbers
    if num == 0 || num == 1 {
    
        return false
        
    } // if
    
    // loop from 2 up to num -1
    // num -1 because every number divides into themselves
    for i := 2; i < num; i++ {
        
        // if num divides evenly
        if num % i == 0 {
            
            // number not prime
            // only 1 and num itself should divide into num evenly
            return false
            
        } // if
        
    } // for 
   
    // if code made it this far, number is prime
    return true

} // isPrime()


