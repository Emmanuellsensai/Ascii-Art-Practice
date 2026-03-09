package ascii

import (
	"fmt"
	"os"
	"strings"
)

func ReadBanner(file string) ([]string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	return lines, nil
}

func BuildAsciiMap(lines []string) map[rune][]string {
	asciiMap := make(map[rune][]string)

	char := 32

	for i := 0; i < len(lines); i += 9 {
		asciiMap[rune(char)] = lines[i : i+8]
		char++
	}

	return asciiMap
}

func PrintAscii(text string, asciiMap map[rune][]string) {

	lines := strings.Split(text, "\\n")

	for _, line := range lines {

		if line == "" {
			fmt.Println()
			continue
		}

		for row := 0; row < 8; row++ {

			for _, char := range line {
				fmt.Print(asciiMap[char][row])
			}

			fmt.Println()
		}
	}
}