package builtins

import (
	"fmt"
	"regexp"
	"strings"
	//"strings"
)

func Alias(input string, alias map[string][]string) error {

	findAlias := regexp.MustCompile(`alias (\w+)="([^"]+)"`)
	a := findAlias.FindStringSubmatch(input)

	if len(a) != 3 {
		fmt.Println("Error creating alias")
	}

	//fmt.Println("alias: ", a[1])
	//fmt.Println("command: ", a[2])

	parseCommand := strings.Split(a[2], " ")
	parseCommand = parseCommand[0:]
	/*
		for i := range parseCommand {
			fmt.Printf("%d |", i)

			fmt.Printf("%s|", parseCommand[i])

			fmt.Println()
		}
		fmt.Println("the command", parseCommand) */

	alias[a[1]] = parseCommand

	return nil
}
