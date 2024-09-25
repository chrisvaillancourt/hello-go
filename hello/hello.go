package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	  // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
		log.SetPrefix("greetings: ")
		log.SetFlags(0)

		// a slice of names
		names := []string{"Tophs", "Sam", "Olivia"}

		// request a greeting for each name
		messages, err := greetings.Hellos(names)
		// if an error exists, print it and exit the program
		if err != nil {
			log.Fatal(err)
		}
		// if no error, print the message
    fmt.Println(messages)
}
