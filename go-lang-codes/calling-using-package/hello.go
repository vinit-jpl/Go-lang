package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// get a greeting msg and print it
	message := greetings.Hello("Vinit")

	fmt.Println(message)
}
