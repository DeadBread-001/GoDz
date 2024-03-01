package processing

import (
	"errors"
	"fmt"
	"github.com/DeadBread-001/GoDz/tree/dz1part2/functions"
	"os"
	"strings"
)

var errParse = errors.New("error in parsing command line")

const expectedArgs = 2

func ParseCmdLine() error {
	if len(os.Args) != expectedArgs {
		return errParse
	}

	expression := strings.Join(os.Args[1:], "")
	result, err := functions.Calculate(expression)

	if err != nil {
		return fmt.Errorf("error calculating result: %v", err)
	}
	fmt.Println(result)

	return nil
}
