package main

import (
	"fmt"
	"log"

	"github.com/hisamcode/learn-go-from-doc/create-a-go-module/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"hisam", "meja", "pc"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

	message, err := greetings.Hello("Gladys")
	// if an error was returned, print it to the console and exit the program
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

}
