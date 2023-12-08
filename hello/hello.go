package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// setup log
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it
	message, err := greetings.Hello("Joao")

	//if has errors
	if err != nil {
		log.Fatal(err)
	}

	// if no error as returned, print the message

	fmt.Println(message)

}
