package util

import (
	"log"
	"os"
	"strings"
)

func ReadInput(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(content))
}

func SplitInput(input string) []string {
	return strings.Split(input, "\n")
}
