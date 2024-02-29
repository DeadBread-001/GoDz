package main

import (
	"fmt"
	"github.com/DeadBread-001/GoDz/tree/dz1part1/processing"
)

func main() {
	err := processing.ParseCmdLine()
	if err != nil {
		fmt.Printf("Error occured: %v", err)
	}
}
