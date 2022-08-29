package uploader

import (
	"os"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestUpload(t *testing.T) {

	file, err := os.Open("test.txt")

	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = Upload(BayFiles, file)
	if err != nil {
		t.Fatal(err.Error())
	}
}
