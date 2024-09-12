package util

import (
	"strings"
	"testing"
)

func TestTimeParseWithExtraTZ(t *testing.T) {
	// tim, err := TimeParseWithExtraTZ("2006/01/02 15:04", "2024/09/09 06:25", "+0600")
	tim, err := TimeParseWithExtraTZ("2006/01/02 15:04", "2024/09/09 06:25", "+0800")
	if err != nil {
		t.Fatal("got err=", err)
	}
	// if tim.Hour() != 6 || !strings.Contains(tim.String(), "+0600") {
	if tim.Hour() != 6 || !strings.Contains(tim.String(), "+0800") {
		t.Fatal("got: tim=", tim)
	}
}
