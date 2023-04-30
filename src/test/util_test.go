package test

import (
	"main/src/main/util"
	"testing"
)

func TestShiftRise(t *testing.T) {
	data := complexArray(-1, -1, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 5, 5, 4, 4, 3, 3, 2, 2, 1, 1)
	want := complexArray(-1, -1, 0, 0, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 3, 3, 2, 2, 1, 1, 0, 0, 0, 0)
	shift := util.ShiftRise(data, 2)

	if !compareArrays(want, shift) {
		t.Fatalf("\nExpected: %v\nActual: %v", shift, want)
	}
}

func TestShiftLow(t *testing.T) {
	data := complexArray(-1, -1, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 5, 5, 4, 4, 3, 3, 2, 2, 1, 1)
	want := complexArray(-1, -1, 3, 3, 4, 4, 5, 5, 6, 6, 0, 0, 0, 0, 0, 0, 6, 6, 5, 5, 4, 4, 3, 3)
	shift := util.ShiftLow(data, 2)

	if !compareArrays(want, shift) {
		t.Fatalf("\nExpected: %v\nActual: %v", shift, want)
	}
}

func TestShuffleArray(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	cpy := data
	util.ShuffleArray(&cpy)

	if compareArrays(data, cpy) {
		t.Fatalf("Arrays shouldn't be equals")
	}
}
