package main

import (
	"fmt"
)

// Function that will take in the command line arguments and process them.
// Will most likely be a void function for the time being. Haven't decided yet.
func ProcessArgs(args []string) {
	
	// If the only argument is the build/run command then there is nothing for us to do.
	if len(args) == 1 {
		fmt.Println("Please provide the necessary flags to use this ding dang tool, ya frank")
		return
	}

	// Grab the url from the arguments. 
	// Probably should do some sort of error checking.
	url := args[1]

	fmt.Println("Ah, You would like to benchmark " , url, " What a lovely choice.")

	// Get rid of that pesky build/run command.
	newArgs := args[2:]

	// Iterate over the arguments.
	// Printing purely for testing purposes. 
	for i := 0 ; i < len(newArgs); i++ {
		if newArgs[i] == "-h" {
			fmt.Print("\tYou passed in a header flag\n")
			i++
			fmt.Println("\tYou passed in ", newArgs[i], " as the parameter for header flag")
		} else if newArgs[i] == "-n" {
			fmt.Print("\tYou passed in a number of requests flag\n")
			i++
			fmt.Println("\tYou passed in ", newArgs[i], " as the parameter for requests flag")
		} else if newArgs[i] == "-c" {
			fmt.Print("\tYou passed in a number of concurrent request flag\n")
			i++
			fmt.Println("\tYou passed in ", newArgs[i], " as the parameter for concurrency flag")
		} else if newArgs[i] == "-m" {
			fmt.Print("\tYou passed in a request type flag\n")
			i++
			fmt.Println("\tYou passed in ", newArgs[i], " as the parameter for request type flag")
		}
	}

}
