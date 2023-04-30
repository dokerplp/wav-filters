package audio

import (
	"github.com/mjibson/go-dsp/fft"
	"log"
	"main/src/main/util"
	"math"
	"math/cmplx"
)

func Noise(arr *[]int) {
	newData := make([]int, 10)
	for _, v := range *arr {
		if v >= 0 {
			newData = append(newData, v, -v)
		}
	}
	*arr = newData
}

func Volume(arr *[]int, factor int) {
	if factor > 100 || factor < 0 {
		log.Fatal("Incorrect value, must be between 0 and 100")
	}
	for i := range *arr {
		(*arr)[i] = (*arr)[i] * factor / 100
	}
}

func Shuffle(arr *[]int) {
	clx := util.IntToComplexArray(*arr)

	fftData := fft.FFT(clx)
	realData := util.RealPartOfComplexArray(fftData)
	util.ShuffleArray(&realData)

	for i, v := range realData {
		fftData[i] = complex(v, imag(fftData[i]))
	}

	ifftData := fft.IFFT(fftData)
	*arr = util.ComplexToIntArray(ifftData)
}

func RaiseAmplitude(arr *[]int) {
	clx := util.IntToComplexArray(*arr)

	fftData := fft.FFT(clx)
	l := len(fftData)
	for i, v := range fftData {
		fftData[i] = complex(float64(i)/float64(l), imag(v))
	}

	ifftData := fft.IFFT(fftData)
	*arr = util.ComplexToIntArray(ifftData)
}

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

func Shift(arr *[]int, factor float64) {
	clx := util.IntToComplexArray(*arr)

	fftData := fft.FFT(clx)

	newData := make([]complex128, len(fftData))
	for i, v := range fftData {
		newData[i] = complex(real(v), imag(v)*factor)
	}

	ifftData := fft.IFFT(newData)
	*arr = util.ComplexToIntArray(ifftData)
}

func base(arr *[]int, factor int, f func([]complex128, int) []complex128) {
	clx := util.IntToComplexArray(*arr)

	fftData := fft.FFT(clx)
	newData := f(fftData, factor)
	ifftData := fft.IFFT(newData)
	*arr = util.ComplexToIntArray(ifftData)
}

func Low(arr *[]int, factor int) {
	base(arr, factor, util.ShiftLow)
}

func Rise(arr *[]int, factor int) {
	base(arr, factor, util.ShiftRise)
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
