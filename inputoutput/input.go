package inputoutput

import (
	"bufio"
	"fmt"
	"os"
)

func InputToSlice(inputFile string) ([]string, error) {
	var lines []string

	if inputFile != "" {
		file, err := os.Open(inputFile)
		if err != nil {
			return nil, fmt.Errorf("failed inputing: %w", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Ошибка при сканировании:", err)
				break
			}
			lines = append(lines, scanner.Text())
		}

		return lines, nil
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println("Ошибка при сканировании:", err)
			break
		}
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
