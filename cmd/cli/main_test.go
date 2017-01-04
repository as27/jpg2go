package main

import (
	"path/filepath"
	"testing"
)

func TestMakePath(t *testing.T) {
	srcFolder = "b"
	relResizeFolder = "resize2"
	in := filepath.Join("b", "a", "bild.jpg")
	got := makePath(in)
	expect := filepath.Join("b", "resize2", "a", "bild.jpg")
	if got != expect {
		t.Error("MakePath() returns wrong value!")
	}
}

func TestFileFilter(t *testing.T) {
	srcFolder = "b"
	relResizeFolder = "resize2"
	resizeFolder = filepath.Join(srcFolder, relResizeFolder)
	tests := []struct {
		str    string
		expect bool
	}{
		{
			filepath.Join("a", "b", "pic.png"),
			false,
		},
		{
			filepath.Join("a", "b", "pic.PNG"),
			false,
		},
		{
			filepath.Join("a", "b", "pic.jpg"),
			true,
		},
		{
			filepath.Join("a", "b", "pic.JPG"),
			true,
		},
		{
			filepath.Join("a", "b", "pic.jpeg"),
			true,
		},
		{
			filepath.Join("a", "b", "pic.JPEG"),
			true,
		},
	}
	for _, test := range tests {
		got := fileFilter(test.str)
		if got != test.expect {
			t.Errorf("Expect: %v, Got: %v", test.expect, got)
		}
	}
}
