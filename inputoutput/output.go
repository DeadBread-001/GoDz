package inputoutput

import (
	"fmt"
	"os"
	"strings"
)

func SliceToOutput(lines []string, outputFile string) error {
	if outputFile != "" {
		file, err := os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("failed creating file: %w", err)
		}
		defer file.Close()

		_, err = fmt.Fprintln(file, strings.Join(lines, "\n"))
		if err != nil {
			return fmt.Errorf("failed writing to file: %w", err)
		}

		return nil
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return nil
}
