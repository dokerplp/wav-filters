package test

import (
	"main/src/main/audio"
	"testing"
)

func TestReverse(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	want := []int{8, 7, 6, 5, 4, 3, 2, 1}
	audio.Reverse(&data)

	if !compareArrays(data, want) {
		t.Fatalf("\nExpected: %v\nActual: %v", data, want)
	}
}
