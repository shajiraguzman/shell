package builtins

import (
	"fmt"
	"os"
	"strings"
	//"strings"
)

func Export(input string) error {

	args := strings.Split(input, " ")

	//utilized chatGPT for the following string parsing to accomdoate for string variables or numerical values for export.
	// AI help allowed me to run both commands
	// 		1. export stringVal="hello world"
	// 		2. export num1=10

	if len(args) > 1 {
		// Join the remaining parts of the input to reconstruct the variable name and value
		varString := strings.Join(args[1:], " ")
		// Split the variable string by "=" to separate the variable name and its value
		parts := strings.SplitN(varString, "=", 2)
		if len(parts) == 2 {
			// Set the environment variable
			os.Setenv(parts[0], strings.Trim(parts[1], `"`))
		} else {
			fmt.Println("Invalid export syntax. Usage: export VAR=value")
		}
	} else {
		fmt.Println("Usage: export VAR=value")
	}

	return nil

}
