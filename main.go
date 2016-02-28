package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide the necessary flags to use this ding dang tool, ya frank")
		return
	}
	newArgs := os.Args[1:]
	for i := range newArgs {
		fmt.Println("Argument ", i, " is: ", newArgs[i])
		if newArgs[i] == "-h" {
			fmt.Print("\tYou passed in a header flag\n")
		} else if newArgs[i] == "-n" {
			fmt.Print("\tYou passed in a number of requests flag\n")
		} else if newArgs[i] == "-c" {
			fmt.Print("\tYou passed in a number of concurrent request flag\n")
		} else if newArgs[i] == "-m" {
			fmt.Print("\tYou passed in a request type flag\n")
		}
	}

}
