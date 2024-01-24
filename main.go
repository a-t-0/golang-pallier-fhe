package main

import (
	"fmt"

	"rsc.io/quote"
)

// main prints a hello world line and a quote from an external package.
func main() {
	fmt.Println("hello world locally")
	fmt.Println(quote.Go())
}
