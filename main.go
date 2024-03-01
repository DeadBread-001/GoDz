package main

import (
	"fmt"
	"github.com/DeadBread-001/GoDz/tree/dz1part2/processing"
)

func main() {
	err := processing.ParseCmdLine()
	if err != nil {
		fmt.Println(err)
	}
}
