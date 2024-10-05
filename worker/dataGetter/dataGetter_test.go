package datagetter

import (
	"os"
	"testing"
)

func TestHTMLGet(t *testing.T) {
	ticker := "AAPL"
	err := getHTML(ticker)
	if err != nil {
		t.Fatalf("Failed to get HTML: %v", err)
	}

	// Check if the file was created
	filename := ticker + ".html"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatalf("Expected file %s to be created, but it does not exist", filename)
	}

	// Clean up the file after test
	defer os.Remove(filename)
}
