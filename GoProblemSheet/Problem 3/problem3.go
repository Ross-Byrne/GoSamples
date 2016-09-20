package main

import "fmt"
import "bufio"
import "os"

func main() {

	var inputText string = ""

  reader := bufio.NewReader(os.Stdin)

	fmt.Println("Problem 3\n")
	fmt.Println("Accept user string input")
	fmt.Println("and output it reversed.\n")

	// save user input
	fmt.Print("Enter text: ")
  inputText, _ = reader.ReadString('\n')

	// reverse the text
	inputText = reverseString(inputText)

	// print out reversed text
  fmt.Println("\nReversed string: " + inputText)

} // main()

// reverses the string that is passed in
func reverseString(str string) string {

	var result string = ""

  // loop through string backwards
	for i := len(str) - 2; i >= 0; i-- {

    // build the string backwards
		result += string(str[i])

	} // for

	return result

} // reverseString()
