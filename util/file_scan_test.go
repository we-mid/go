package util

import (
	"bytes"
	"testing"
)

func TestFileTail(t *testing.T) {
	lines, err := FileTail("./testdata/lines.txt", 4)
	if err != nil {
		t.Fatal("FileTail: err is", err)
	}
	if bs := bytes.Join(lines, []byte{}); !bytes.Equal(bs, []byte("4.5.6.7.")) {
		t.Fatalf("FileTail: got %q", bs)
	}
}
