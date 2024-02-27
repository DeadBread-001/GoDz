package processing

import (
	"errors"
	"flag"
	"github.com/DeadBread-001/GoDz/tree/dz1part1/functions"
	"github.com/DeadBread-001/GoDz/tree/dz1part1/inputoutput"
)

var errParsing = errors.New("формат команды: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")

type Options struct {
	cFlag bool
	dFlag bool
	uFlag bool
	iFlag bool
	fFlag int
	sFlag int
}

func initFlags(flags *Options) {
	flag.BoolVar(&flags.cFlag, "c", false, "Prefix lines by the number of occurrences")
	flag.BoolVar(&flags.dFlag, "d", false, "Only print duplicate lines, one for each group")
	flag.BoolVar(&flags.uFlag, "u", false, "Only print unique lines")
	flag.BoolVar(&flags.iFlag, "i", false, "Ignore differences in case when comparing")
	flag.IntVar(&flags.fFlag, "f", 0, "Avoid comparing the first N fields")
	flag.IntVar(&flags.sFlag, "s", 0, "Avoid comparing the first N characters")
}

// uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]].
func executeUniq(flags Options, lines []string) ([]string, error) {
	var err error

	linesCopy := make([]string, len(lines))
	copy(linesCopy, lines)

	if flags.iFlag {
		linesCopy = functions.IgnoreCase(linesCopy)
	}

	if flags.fFlag != 0 {
		linesCopy, err = functions.IgnoreFields(linesCopy, flags.fFlag)
		if err != nil {
			return nil, err
		}
	}

	if flags.sFlag != 0 {
		linesCopy, err = functions.IgnoreCharacters(linesCopy, flags.sFlag)
		if err != nil {
			return nil, err
		}
	}

	var suitableLines []bool

	switch {
	case (flags.cFlag != flags.dFlag) != flags.uFlag:
		if flags.cFlag {
			suitableLines = functions.CountOfLines(linesCopy, lines)
		}

		if flags.dFlag {
			suitableLines = functions.DuplicatedLines(linesCopy)
		}

		if flags.uFlag {
			suitableLines = functions.UniqLines(linesCopy)
		}

	case !(flags.cFlag || flags.dFlag || flags.uFlag):
		suitableLines = functions.StandartUniq(linesCopy)
	default:
		return nil, errParsing
	}

	lines, err = functions.BuildResult(suitableLines, lines)
	if err != nil {
		return nil, err
	}

	return lines, nil
}

func ParseCmdLine() error {
	var flags Options

	initFlags(&flags)
	flag.Parse()

	inputFile := flag.Arg(0)
	outputFile := flag.Arg(1)

	lines, err := inputoutput.InputToSlice(inputFile)
	if err != nil {
		return err
	}

	lines, err = executeUniq(flags, lines)
	if err != nil {
		return err
	}

	err = inputoutput.SliceToOutput(lines, outputFile)
	if err != nil {
		return err
	}

	return nil
}
