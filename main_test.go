package main

import (
	"testing"
)

func TestHasSuffixInArray(t *testing.T) {
	testCases := []struct {
		str  string
		arr  []string
		want bool
	}{
		{"image1.png", []string{"jpeg", "png"}, true},
		{"image1.png", []string{"JPEG", "PNG"}, true},
		{"image1.jpeg", []string{"asdf", "blahblah"}, false},
	}
	for _, tc := range testCases {
		if res := hasSuffixInArray(tc.str, tc.arr); res != tc.want {
			t.Errorf("\nhasSuffixInArray(%s, %s)\nwant %t\ngot %t", tc.str, tc.arr, tc.want, res)
		}
	}
}
