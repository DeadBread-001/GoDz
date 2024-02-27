package main

import (
	"github.com/DeadBread-001/GoDz/tree/dz1part1/processing"
	"log"
)

func main() {
	err := processing.ParseCmdLine()
	if err != nil {
		log.Printf("Error occured: %v", err) //nolint:misspell
	}
}
