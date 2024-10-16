package util

import (
	"os"
	"testing"
)

func TestTildify(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("os.UserHomeDir: err is", err)
	}
	s, err := Tildify(homeDir + "/foo/bar") // prepend
	if err != nil {
		t.Fatal("Tildify: err is", err)
	}
	w := "~/foo/bar"
	if s != w {
		t.Fatalf("Tildify: want %q, got %q", w, s)
	}

	s, _ = Tildify(homeDir + "foo/bar") // without slash
	w = homeDir + "foo/bar"             // not touched
	if s != w {
		t.Fatalf("Tildify: want %q, got %q", w, s)
	}

	s, _ = Tildify("foo/bar/" + homeDir) // append
	w = "foo/bar/" + homeDir             // not touched
	if s != w {
		t.Fatalf("Tildify: want %q, got %q", w, s)
	}

	s, _ = Tildify(homeDir + "/foo/bar/" + homeDir) // prepend & append
	w = "~/foo/bar/" + homeDir                      // only process prefix
	if s != w {
		t.Fatalf("Tildify: want %q, got %q", w, s)
	}
}

func TestUntildify(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("os.UserHomeDir: err is", err)
	}
	s, err := Untildify("~/foo/bar") // prepend
	if err != nil {
		t.Fatal("Tildify: err is", err)
	}
	w := homeDir + "/foo/bar"
	if s != w {
		t.Fatalf("Untildify: want %q, got %q", w, s)
	}

	s, _ = Untildify("~foo/bar") // without slash
	w = "~foo/bar"               // not touched
	if s != w {
		t.Fatalf("Tildify: want %q, got %q", w, s)
	}

	s, _ = Untildify("foo/bar/~/") // append
	w = "foo/bar/~/"               // not touched
	if s != w {
		t.Fatalf("Tildify: want %q, got %q", w, s)
	}

	s, _ = Untildify("~/foo/bar/~/") // prepend & append
	w = homeDir + "/foo/bar/~/"      // only process prefix
	if s != w {
		t.Fatalf("Tildify: want %q, got %q", w, s)
	}
}
