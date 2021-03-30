package validate

import (
	"testing"
)

func TestFileInput1(t *testing.T) {
	filename := "test.htm"
	msg, err := checkFilename(filename)
	if err == nil {
		t.Fatalf(`CheckFilename(test.htm)= %q, %v, want err msg "required .html"`, msg, err)
	}

}

func TestFileInput2(t *testing.T) {
	filename := "fileNotExist.html"
	msg, err := checkFilename(filename)
	if err == nil {
		t.Fatalf(`CheckFilename(test.htm)= %q, %v, want err msg"file not exist"`, msg, err)
	}

}
