// This sample program demonstrates how to create goroutines and
// how the scheduler behaves.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// main is the entry point for all Go programs
func main() {
	// Allocate 1 logical processor for the scheduler to use
	runtime.GOMAXPROCS(1)

	// WaitGroups are used to wait for the program to finish
	// Add a count of two, one for each goroutine
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Start off goroutines")
	// Declare an anonymous function and create a goroutine
	go func() {
		defer wg.Done()
		// Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
			fmt.Println()
		}
	}()
	// Declare another anonymous function and create a goroutine
	go func() {
		defer wg.Done()
		// Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
			fmt.Println()
		}
	}()
	// Wait goroutine to finish
	wg.Wait()
	fmt.Println("\n Terminate program")
}
