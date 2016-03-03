package main

import (
	"fmt"
    "strconv"
)

// Process Args is a function that takes in some amount of command line arguments and then do some processing.
// Will most likely be a void function for the time being. Haven't decided yet.
func ProcessArgs(args []string) error {
	
	// If the only argument is the build/run command then there is nothing for us to do.
	if len(args) == 1 {
	   return fmt.Errorf("Please provide the necessary flags to use this ding dang tool, ya frank")
	}

	// Grab the url from the arguments. 
	// Probably should do some sort of error checking.
	url := args[1]

	fmt.Println("Ah, You would like to benchmark " , url, " What a lovely choice.")

	// Get rid of that pesky build/run command.
	newArgs := args[2:]

	// Struct to append to while reading from flags.
	var flags Spotter
    var err error
	// Iterate over the arguments.
	// Printing purely for testing purposes. 
	for i := 0 ; i < len(newArgs); i++ {
		if newArgs[i] == "-h" {
			fmt.Print("\tYou passed in a header flag\n")
			i++
			flags.Header = newArgs[i]
			fmt.Println("\tYou passed in ", newArgs[i], " as the parameter for header flag")
		} else if newArgs[i] == "-n" {
			fmt.Print("\tYou passed in a number of requests flag\n")
			i++
			flags.RequestNumber, err = strconv.Atoi(newArgs[i])
            if err != nil {
                return err
            }
			fmt.Println("\tYou passed in ", newArgs[i], " as the parameter for requests flag")
		} else if newArgs[i] == "-c" {
			fmt.Print("\tYou passed in a number of concurrent request flag\n")
			i++
			flags.ConcurrencyLevel, err = strconv.Atoi(newArgs[i])
            if err != nil {
                return err
            }
			fmt.Println("\tYou passed in ", newArgs[i], " as the parameter for concurrency flag")
		} else if newArgs[i] == "-m" {
			fmt.Print("\tYou passed in a request type flag\n")
			i++
			flags.RequestType = newArgs[i]
			fmt.Println("\tYou passed in ", newArgs[i], " as the parameter for request type flag")
		}
	}
    
    // Printing out just to make sure they are being added to the struct. 
    fmt.Println(flags)
    
	// Maybe pass this struct to another function to process it and then start executing.
    // Return a dummy thing for now. 
    return nil
}
