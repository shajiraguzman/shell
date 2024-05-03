package builtins

import (
	"fmt"
	"os"
)

func Pwd() error {

	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mydir)

	return nil

}
