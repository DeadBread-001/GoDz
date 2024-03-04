package uniq

import (
	"errors"
	"flag"
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

func InitFlags(flags *Options) {
	flag.BoolVar(&flags.cFlag, "c", false, "Prefix lines by the number of occurrences")
	flag.BoolVar(&flags.dFlag, "d", false, "Only print duplicate lines, one for each group")
	flag.BoolVar(&flags.uFlag, "u", false, "Only print unique lines")
	flag.BoolVar(&flags.iFlag, "i", false, "Ignore differences in case when comparing")
	flag.IntVar(&flags.fFlag, "f", 0, "Avoid comparing the first N fields")
	flag.IntVar(&flags.sFlag, "s", 0, "Avoid comparing the first N characters")
}

// uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]].
func ExecuteUniq(flags Options, lines []string) ([]string, error) {
	var err error

	linesCopy := make([]string, len(lines))
	copy(linesCopy, lines)

	if flags.iFlag {
		linesCopy = IgnoreCase(linesCopy)
	}

	if flags.fFlag != 0 {
		linesCopy, err = IgnoreFields(linesCopy, flags.fFlag)
		if err != nil {
			return nil, err
		}
	}

	if flags.sFlag != 0 {
		linesCopy, err = IgnoreCharacters(linesCopy, flags.sFlag)
		if err != nil {
			return nil, err
		}
	}

	var suitableLines []bool

	switch {
	case (flags.cFlag != flags.dFlag) != flags.uFlag:
		if flags.cFlag {
			suitableLines = CountOfLines(linesCopy, lines)
		}

		if flags.dFlag {
			suitableLines = DuplicatedLines(linesCopy)
		}

		if flags.uFlag {
			suitableLines = UniqLines(linesCopy)
		}

	case !(flags.cFlag || flags.dFlag || flags.uFlag):
		suitableLines = StandartUniq(linesCopy)
	default:
		return nil, errParsing
	}

	lines, err = BuildResult(suitableLines, lines)
	if err != nil {
		return nil, err
	}

	return lines, nil
}
