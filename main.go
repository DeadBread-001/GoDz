package main

import (
	"errors"
	"fmt"
	"github.com/DeadBread-001/GoDz/tree/dz1part2/functions"
	"os"
	"strings"
)

var errParse = errors.New("error in parsing command line")

const expectedArgs = 2

func main() {
	if len(os.Args) != expectedArgs {
		fmt.Println(errParse)
	}

	expression := strings.Join(os.Args[1:], "")
	result, err := functions.Calculate(expression)

	if err != nil {
		fmt.Printf("error calculating result: %v\n", err)
	}
	fmt.Println(result)
}
