package uniq

import (
	"errors"
	"strconv"
)

var errProcess = errors.New("selection funcs error")

func BuildResult(suitableLines []bool, lines []string) ([]string, error) {
	if len(lines) != len(suitableLines) {
		return nil, errProcess
	}

	var resultLines []string

	for i := range lines {
		if suitableLines[i] {
			resultLines = append(resultLines, lines[i])
		}
	}

	return resultLines, nil
}

func StandartUniq(lines []string) []bool {
	if len(lines) == 0 {
		return []bool{}
	}

	uniqueLines := make([]bool, len(lines))
	uniqueLines[0] = true

	for i := 1; i < len(lines); i++ {
		uniqueLines[i] = lines[i] != lines[i-1]
	}

	return uniqueLines
}

func DuplicatedLines(lines []string) []bool {
	if len(lines) == 0 {
		return []bool{}
	}

	duplicatedLines := make([]bool, len(lines))
	currIdx := 0

	for i := 1; i < len(lines); i++ {
		if lines[i] == lines[currIdx] {
			duplicatedLines[currIdx] = true
		} else {
			currIdx = i
		}
	}

	return duplicatedLines
}

func UniqLines(lines []string) []bool {
	if len(lines) == 0 {
		return []bool{}
	}

	uniqLines := make([]bool, len(lines))
	cnt := 1

	for i := 1; i < len(lines); i++ {
		if lines[i] != lines[i-1] {
			uniqLines[i-1] = cnt == 1
			cnt = 1
		} else {
			cnt++
		}
	}

	uniqLines[len(lines)-1] = cnt == 1

	return uniqLines
}

func CountOfLines(lines []string, originalLines []string) []bool {
	if len(lines) == 0 {
		return []bool{}
	}

	countedLines := make([]bool, len(lines))
	cnt := 1

	for i := 1; i < len(lines); i++ {
		if lines[i] != lines[i-1] {
			originalLines[i-1] = strconv.Itoa(cnt) + " " + originalLines[i-1]
			countedLines[i-1] = true
			cnt = 1
		} else {
			cnt++
		}
	}

	originalLines[len(lines)-1] = strconv.Itoa(cnt) + " " + originalLines[len(lines)-1]
	countedLines[len(lines)-1] = true

	return countedLines
}
