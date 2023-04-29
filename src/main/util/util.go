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

func IntToComplexArray(arr []int) []complex128 {
	clx := make([]complex128, len(arr))
	for i := 0; i < len(arr); i++ {
		clx[i] = complex(float64(arr[i]), 0)
	}
	return clx
}

func ComplexToIntArray(arr []complex128) []int {
	realData := make([]int, len(arr))
	for i, c := range arr {
		realData[i] = int(real(c))
	}
	return realData
}
