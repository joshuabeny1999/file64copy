package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestPrintHelp(t *testing.T) {
	// Capture the standard output
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printHelp()

	// Restore the standard output
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	// Check the output
	if !strings.Contains(string(out), "Usage: file64copy [OPTIONS] FILE") {
		t.Errorf("printHelp() = %s; want 'Usage: file64copy [OPTIONS] FILE'", string(out))
	}
}

func TestPrintVersion(t *testing.T) {
	// Capture the standard output
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printVersion()

	// Restore the standard output
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	// Check the output
	if !strings.Contains(string(out), "Version: development") {
		t.Errorf("printVersion() = %s; want 'Version: development'", string(out))
	}
}

func TestEncodeFile(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "example.*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	_, err = encodeFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerateCommand(t *testing.T) {
	result := generateCommand("Test encoded content", "test.txt")
	expected := `echo "Test encoded content" | openssl base64 -d -A -out test.txt`
	if result != expected {
		t.Errorf("generateCommand() = %s; want '%s'", result, expected)
	}
}
