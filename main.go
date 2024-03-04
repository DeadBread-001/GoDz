package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/DeadBread-001/GoDz/tree/dz1part2/functions"
)

var errParse = errors.New("error in parsing command line")

const expectedArgs = 2

func main() {
	if len(os.Args) != expectedArgs {
		log.Fatalf("Error: %v", errParse)
	}

	expression := strings.Join(os.Args[1:], "")
	result, err := functions.Calculate(expression)

	if err != nil {
		log.Fatalf("error calculating result: %v\n", err)
	}
	fmt.Println(result)
}
