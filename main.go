package main

import (
	"fmt"

	"github.com/a-t-0/golang-pallier-fhe/fhe"
	"rsc.io/quote"
)

// main prints a hello world line and a quote from an external package.
func main() {
	fmt.Println("hello world locally")
	fmt.Println(quote.Go())

	// Print the output of a function call of a file in this repo.
	fmt.Printf("The result of 3 + 2 is: %d\n", fhe.AddTwo(3))
}
