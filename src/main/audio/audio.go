package audio

import (
	"github.com/mjibson/go-dsp/fft"
	"main/src/main/util"
	"math"
	"math/cmplx"
)

func Reverse(arr *[]int) {
	for i, j := 0, len(*arr)-1; i < len(*arr)/2; i, j = i+1, j-1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}

func Temp(arr *[]int, factor float64) {
	l := len(*arr)
	newArr := make([]int, int(math.Floor(float64(l)/factor)))
	newArr[0] = (*arr)[0]
	for i := 1; i < len(newArr); i++ {
		floor := int(math.Floor(float64(i) * factor))
		ceil := int(math.Ceil(float64(i) * factor))
		if ceil >= l {
			break
		}
		newArr[i] = ((*arr)[floor] + (*arr)[ceil]) / 2
	}
	*arr = newArr
}

func pitch(arr *[]int, factor int, f func([]complex128, int) []complex128) {
	clx := util.IntToComplexArray(*arr)

	fftData := fft.FFT(clx)
	newData := f(fftData, factor)
	ifftData := fft.IFFT(newData)
	realData := util.ComplexToIntArray(ifftData)

	*arr = realData
}

func PitchLow(arr *[]int, factor int) {
	pitch(arr, factor, util.ShiftLow)
}

func PitchRise(arr *[]int, factor int) {
	pitch(arr, factor, util.ShiftRise)
}

func Gpt(arr *[]int, factor float64) {
	audioData := *arr
	fftSize := 1024
	fftData := make([]complex128, fftSize)
	for i, v := range audioData {
		fftData[i%fftSize] = complex(float64(v), 0)
		if i%fftSize == fftSize-1 {
			fft.FFT(fftData)
			for j := range fftData {
				im := -2 * math.Pi * float64(j) * factor / float64(fftSize)
				fftData[j] *= complex(1, 0) * cmplx.Exp(complex(0, im))
			}
			fft.IFFT(fftData)
			for j := range fftData {
				audioData[i-fftSize+j+1] = int(real(fftData[j]))
			}
		}
	}
	*arr = audioData
}
