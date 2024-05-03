package builtins

import (
	"strings"
)

func Alias(input string, alias map[string]string) error {

	parts := strings.SplitN(input, "=", 2)
	aliasName := strings.TrimSpace(parts[0])

	// get alias name
	aliasNameParts := strings.Fields(aliasName)
	if len(aliasNameParts) > 1 {
		aliasName = aliasNameParts[1]
	}

	// remove double quotes from the command
	command := strings.Trim(parts[1], `"`)

	alias[aliasName] = command

	return nil
}
