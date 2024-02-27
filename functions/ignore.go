package functions

import (
	"errors"
	"strings"
)

var errNum = errors.New("invalid input number")

func IgnoreCase(lines []string) []string {
	for i := range lines {
		lines[i] = strings.ToLower(lines[i])
	}

	return lines
}

func IgnoreFields(lines []string, num int) ([]string, error) {
	if num < 0 {
		return nil, errNum
	}

	for i := range lines {
		words := strings.Split(lines[i], " ")
		if num > len(words) {
			lines[i] = ""
		} else {
			lines[i] = strings.Join(words[num:], " ")
		}
	}

	return lines, nil
}

func IgnoreCharacters(lines []string, num int) ([]string, error) {
	if num < 0 {
		return nil, errNum
	}

	for i := range lines {
		if num > len((lines)[i]) {
			lines[i] = ""
		} else {
			lines[i] = lines[i][num:]
		}
	}

	return lines, nil
}
