package common

import (
	"os"
	"unicode"
)

func FileExists(filename string) bool {
    _, err := os.Stat(filename)
    return !os.IsNotExist(err)
}

// Utility function to count words in a string
func CountWords(text string) int {
    return len(text)
}

func SplitText(text string, parts int) []string {
	if parts <= 0 {
		return []string{text} // Return the entire text if parts is invalid
	}

	// Clean text: remove spaces and newlines
	cleanedText := ""
	for _, r := range text {
		if !unicode.IsSpace(r) {
			cleanedText += string(r)
		}
	}

	// Calculate the length of each part
	textLength := len(cleanedText)
	partSize := textLength / parts
	remainder := textLength % parts

	var result []string
	start := 0

	for i := 0; i < parts; i++ {
		// Adjust the size for parts that need an extra character to account for the remainder
		end := start + partSize
		if i < remainder {
			end++
		}
		// Avoid out-of-bounds indexing
		if end > textLength {
			end = textLength
		}
		result = append(result, cleanedText[start:end])
		start = end
	}
	return result
}