package util

import (
	"os"
	"regexp"
	"strings"
)

var regexTildy = regexp.MustCompile(`^~/`)

func Untildify(s string) (string, error) {
	if regexTildy.MatchString(s) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		s = regexTildy.ReplaceAllLiteralString(s, homeDir+"/")
	}
	return s, nil
}

func Tildify(s string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	if strings.HasPrefix(s, homeDir+"/") {
		s = strings.Replace(s, homeDir+"/", "~/", 1)
	}
	return s, nil
}
