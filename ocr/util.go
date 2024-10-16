package ocr

import (
	"os"
	"regexp"
)

var regexTildy = regexp.MustCompile(`^~/`)

func handleTildy(s string) (string, error) {
	if regexTildy.MatchString(s) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		s = regexTildy.ReplaceAllLiteralString(s, homeDir+"/")
	}
	return s, nil
}
