package util

import (
	"strings"
	"testing"
)

func TestSpliceLoose(t *testing.T) {
	ls := []string{}
	ls = SpliceLoose(ls, 0, 1, []string{"1."})
	if s := strings.Join(ls, ""); s != "1." {
		t.Fatalf("got %q", s)
	}
	ls = SpliceLoose(ls, 0, 1, []string{"2."})
	if s := strings.Join(ls, ""); s != "2." {
		t.Fatalf("got %q", s)
	}
	ls = SpliceLoose(ls, 1, 0, []string{"3."})
	if s := strings.Join(ls, ""); s != "2.3." {
		t.Fatalf("got %q", s)
	}
	ls = SpliceLoose(ls, 1, 1, []string{"4."})
	if s := strings.Join(ls, ""); s != "2.4." {
		t.Fatalf("got %q", s)
	}
}
