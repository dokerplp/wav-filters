package util

import (
	"log"
)

func LogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ShiftLow(data []complex128, shift int) []complex128 {
	mid := (len(data) + 1) / 2
	newData := []complex128{data[0]}
	newData = append(newData, data[shift+1:mid+1]...)
	newData = append(newData, make([]complex128, 2*shift-1)...)
	newData = append(newData, data[mid:len(data)-shift]...)
	return newData
}

func ShiftRise(data []complex128, shift int) []complex128 {
	mid := (len(data) + 1) / 2
	newData := []complex128{data[0]}
	newData = append(newData, make([]complex128, shift)...)
	newData = append(newData, data[1:mid-shift]...)
	newData = append(newData, data[mid+shift:]...)
	newData = append(newData, make([]complex128, shift)...)
	return newData
}
