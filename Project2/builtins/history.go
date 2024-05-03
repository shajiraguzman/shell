package builtins

import (
	"fmt"
	//"strings"
)

func History(commands [][]string, args ...string) error {

	//  ls -l | grep main

	for i := range commands {
		fmt.Printf("%d ", i)
		for j := range commands[i] {
			fmt.Printf("%s ", commands[i][j])
			//commands[i][j] = args[j]
			//commands[i] = append(commands[i], args[j])
		}
		fmt.Println()
	}

	return nil
}
