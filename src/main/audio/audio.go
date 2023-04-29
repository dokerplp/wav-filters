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
	newArr := make([]int, int(math.Floor(float64(len(*arr))/factor)))
	newArr[0] = (*arr)[0]
	for i := 1; i < len(newArr); i++ {
		floor := int(math.Floor(float64(i) * factor))
		ceil := int(math.Ceil(float64(i) * factor))
		newArr[i] = ((*arr)[floor] + (*arr)[ceil]) / 2
	}
	*arr = newArr
}

func Pitch(arr *[]int, factor float64, rate int) {
	clx := make([]complex128, len(*arr))
	arrLen := len(*arr)
	for i := 0; i < arrLen; i++ {
		clx[i] = complex(float64((*arr)[i]), 0)
	}

	fftData := fft.FFT(clx)

	newData := util.ShiftRise(fftData, 50000)

	ifftData := fft.IFFT(newData)

	realData := make([]int, arrLen)
	for i, c := range ifftData {
		realData[i] = int(real(c))
	}

	*arr = realData
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
