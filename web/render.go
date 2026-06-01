package web

import (
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile("banner/" + filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) == 0 {
		return nil, err
	}
	font := make(map[rune][]string)

	for ch := ' '; ch <= '~'; ch++ {
		start := (int(ch) - 32) * 9
		font[ch] = lines[start+1 : start+9]
	}

	return font, nil
}

func Render(text string, banner map[rune][]string) string {
	var result strings.Builder
	word := strings.Split(text, "\n")
	for _, words := range word {
		if words == "" {
			result.WriteString("\n")
			continue
		}

		for row := 0; row < 8; row++ {
			for _, ch := range words {
				ascii, ok := banner[ch]
				if ok {
					result.WriteString(ascii[row])
				}
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}
