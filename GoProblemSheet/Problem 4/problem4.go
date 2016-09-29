package main

import "fmt"

func collatz(n uint){

  for ; n != 1; {

    fmt.Print(n, " ")

    if n % 2 == 0{

      n = n / 2

    } else {

      n = (3 * n) + 1
    }
  }

  fmt.Println(n)

}

func main(){

  fmt.Println("Problem 4\n")

  collatz(50)

  fmt.Println()
} // main
