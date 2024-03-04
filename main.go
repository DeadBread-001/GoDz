package main

import (
	"flag"
	"log"

	"github.com/DeadBread-001/GoDz/tree/dz1part1/inputoutput"
	"github.com/DeadBread-001/GoDz/tree/dz1part1/uniq"
)

func main() {
	var flags uniq.Options

	uniq.InitFlags(&flags)
	flag.Parse()

	inputFile := flag.Arg(0)
	outputFile := flag.Arg(1)

	lines, err := inputoutput.InputToSlice(inputFile)
	if err != nil {
		log.Fatalf("Error occured: %v", err)
	}

	lines, err = uniq.ExecuteUniq(flags, lines)
	if err != nil {
		log.Fatalf("Error occured: %v", err)
	}

	err = inputoutput.SliceToOutput(lines, outputFile)
	if err != nil {
		log.Fatalf("Error occured: %v", err)
	}
}
