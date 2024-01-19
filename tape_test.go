package poker_test

import (
	"io"
	"testing"

	"github.com/gabeno/poker/v1"
)

func TestTape_Write(t *testing.T) {
	tempFile, clean := poker.CreateTempFile(t, "12345")
	defer clean()

	tape := poker.Tape{tempFile}
	tape.Write([]byte("abc"))

	tempFile.Seek(0, 0)
	newFileContents, _ := io.ReadAll(tempFile)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
