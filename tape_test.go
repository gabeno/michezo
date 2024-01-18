package main

import (
	"io"
	"testing"
)

func TestTape_Write(t *testing.T) {
	tempFile, clean := createTempFile(t, "12345")
	defer clean()

	tape := tape{tempFile}
	tape.Write([]byte("abc"))

	tempFile.Seek(0, 0)
	newFileContents, _ := io.ReadAll(tempFile)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
