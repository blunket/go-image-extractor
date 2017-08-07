package main

import (
	"testing"
)

func TestA(t *testing.T) {
	t.Log(`hasSuffixInArray("image1.png", ["jpeg", "png"]) should be true`)

	expect := true

	s, arr := "image1.png", []string{"jpeg", "png"}
	if res := hasSuffixInArray(s, arr); res != expect {
		t.Errorf("Expected: %b", expect)
		t.Errorf("     Got: %b", res)
	}
}

func TestCases(t *testing.T) {
	t.Log(`hasSuffixInArray("image1.png", ["JPEG", "PNG"]) should be true`)

	expect := true

	s, arr := "image1.png", []string{"JPEG", "PNG"}
	if res := hasSuffixInArray(s, arr); res != expect {
		t.Errorf("Expected: %b", expect)
		t.Errorf("     Got: %b", res)
	}
}

func TestB(t *testing.T) {
	t.Log(`hasSuffixInArray("image1.jpg", ["blahblah", "zzz"]) should be false`)

	expect := false

	s, arr := "image1.jpg", []string{"blahblah", "zzz"}
	if res := hasSuffixInArray(s, arr); res != expect {
		t.Errorf("Expected: %b", expect)
		t.Errorf("     Got: %b", res)
	}
}
