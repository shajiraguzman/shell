package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/shajiraguzman/shell/Project2/builtins"
)

var commands [][]string             //store all commands
var alias = make(map[string]string) //store all aliases created

func main() {
	exit := make(chan struct{}, 2) // buffer this so there's no deadlock.
	runLoop(os.Stdin, os.Stdout, os.Stderr, exit)
}

func runLoop(r io.Reader, w, errW io.Writer, exit chan struct{}) {
	var (
		input    string
		err      error
		readLoop = bufio.NewReader(r)
	)
	for {
		select {
		case <-exit:
			_, _ = fmt.Fprintln(w, "exiting gracefully...")
			return
		default:
			if err := printPrompt(w); err != nil {
				_, _ = fmt.Fprintln(errW, err)
				continue
			}
			if input, err = readLoop.ReadString('\n'); err != nil {
				_, _ = fmt.Fprintln(errW, err)
				continue
			}
			if err = handleInput(w, input, exit); err != nil {
				_, _ = fmt.Fprintln(errW, err)
			}
		}
	}
}

func printPrompt(w io.Writer) error {
	// Get current user.
	// Don't prematurely memoize this because it might change due to `su`?
	u, err := user.Current()
	if err != nil {
		return err
	}
	// Get current working directory.
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	// /home/User [Username] $
	_, err = fmt.Fprintf(w, "%v [%v] $ ", wd, u.Username)

	return err
}

func storeCommand(input string) {

	args := strings.Split(input, " ")
	args = args[0:] //, args[1:]
	commands = append(commands, args[0:])

}

func handleInput(w io.Writer, input string, exit chan<- struct{}) error {

	//sh := os.Getenv("SHELL") //fetch default shell

	input = strings.TrimSpace(input)
	//check if command is an alias first
	if value, ok := alias[input]; ok {
		storeCommand(input)
		input = value
	} else {
		storeCommand(input)
	}

	args := strings.Split(input, " ")
	name, args := args[0], args[1:]

	// Check for built-in commands.
	// New builtin commands should be added here. Eventually this should be refactored to its own func.

	switch name {
	case "cd":
		return builtins.ChangeDirectory(args...)
	case "env":
		return builtins.EnvironmentVariables(w, args...)
	case "history":
		return builtins.History(commands)
	case "alias":
		return builtins.Alias(input, alias)
	case "export":
		return builtins.Export(input)
	case "echo":
		return builtins.Echo(executeCommand, input, name, args...)
	case "pwd":
		return builtins.Pwd()
	case "exit":
		exit <- struct{}{}
		return nil
	}

	return executeCommand(name, args...)

}

func executeCommand(name string, arg ...string) error {
	// Otherwise prep the command
	cmd := exec.Command(name, arg...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
}
