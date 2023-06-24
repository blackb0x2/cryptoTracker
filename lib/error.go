package lib

import (
	"fmt"
	"log"
)

// Function to print error and exit the program
func CriticalError(err error) {
	log.Panicf("Critical Error: %s\n", err)
}

// Function to print an error message
func PrintError(msg string) {
	fmt.Printf("Error: %s\n", msg)
}
