package ascii_art_web

import (
	"os"
	"strings"
)

const ART_CHAR_HEIGHT = 8 // Height of each ASCII art character

// ArgProcessor processes an input string and generates its ASCII art representation
// based on a provided template file.
func ArgProcessor(input, templatePath string) (string, error) {
	output := ""

	if strings.Count(input, "\r\n") == len(input) {
		return input, nil
	}

	data, err := os.ReadFile(templatePath)
	if err != nil {
		return "", err
	}

	content := strings.ReplaceAll(string(data), "\r", "")
	bannerContent := strings.Split(content[1:], "\n\n")

	artMap := AsciiArtGenerator(bannerContent)
	argWords := strings.Split(input, "\r\n")

	for _, word := range argWords {
		if word == "" {
			output += "\n"
			continue
		}

		for l := 0; l < ART_CHAR_HEIGHT; l++ {
			for _, c := range word {
				output += artMap[c][l]
			}
			output += "\n"
		}
	}
	return output, nil
}

// AsciiArtGenerator creates a mapping of characters to their ASCII art representation
// by reading from the provided banner content.
func AsciiArtGenerator(bannerContent []string) map[rune][]string {
	characterMatrix := make(map[rune][]string)

	for i, artChar := range bannerContent {
		characterMatrix[rune(32+i)] = strings.Split(artChar, "\n")
	}
	return characterMatrix
}
