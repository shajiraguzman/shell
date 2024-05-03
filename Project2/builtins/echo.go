package builtins

import (
	"os"
	"os/exec"
)

// executeCommand function from main defined here
type Execute func(name string, arg ...string) error

func Echo(execute Execute, input string, name string, args ...string) error {
	//fmt.Println("input: ", input)
	sh := os.Getenv("SHELL") // get shell name

	// Use echo to print the variable
	echoCmd := exec.Command(sh, "-c", input) //this works!

	echoCmd.Stderr = os.Stderr
	echoCmd.Stdout = os.Stdout

	//args = ""
	// Execute the command and return the error.
	return echoCmd.Run()

}
