package main

import (
	"flag"
	_ "fmt"
	_ "github.com/go-audio/audio"
	"log"
	"math"
	"math/cmplx"
	"os"

	"github.com/go-audio/wav"
	"github.com/mjibson/go-dsp/fft"
)

var (
	inputFlag  = flag.String("i", "in.wav", "input file")
	outputFlag = flag.String("o", "out.wav", "output file")
)

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func reverse(arr *[]int) {
	for i, j := 0, len(*arr)-1; i < len(*arr)/2; i, j = i+1, j-1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}

func temp(arr *[]int, factor float64) {
	newArr := make([]int, int(math.Floor(float64(len(*arr))/factor)))
	newArr[0] = (*arr)[0]
	for i := 1; i < len(newArr); i++ {
		floor := int(math.Floor(float64(i) * factor))
		ceil := int(math.Ceil(float64(i) * factor))
		newArr[i] = ((*arr)[floor] + (*arr)[ceil]) / 2
	}
	*arr = newArr
}

func shiftLow(data []complex128, shift int) []complex128 {
	mid := (len(data) + 1) / 2
	newData := []complex128{data[0]}
	newData = append(newData, data[shift+1:mid+1]...)
	newData = append(newData, make([]complex128, 2*shift-1)...)
	newData = append(newData, data[mid:len(data)-shift]...)
	return newData
}

func shiftRise(data []complex128, shift int) []complex128 {
	mid := (len(data) + 1) / 2
	newData := []complex128{data[0]}
	newData = append(newData, make([]complex128, shift)...)
	newData = append(newData, data[1:mid-shift]...)
	newData = append(newData, data[mid+shift:]...)
	newData = append(newData, make([]complex128, shift)...)
	return newData
}

func pitch(arr *[]int, factor float64, rate int) {
	clx := make([]complex128, len(*arr))
	arrLen := len(*arr)
	for i := 0; i < arrLen; i++ {
		clx[i] = complex(float64((*arr)[i]), 0)
	}

	fftData := fft.FFT(clx)

	newData := shiftRise(fftData, 50000)

	ifftData := fft.IFFT(newData)

	realData := make([]int, arrLen)
	for i, c := range ifftData {
		realData[i] = int(real(c))
	}

	*arr = realData
}

// cool but smth different, factor -1000, -10000, -100000 is awesome
func gpt(arr *[]int, factor float64) {
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

func main() {
	flag.Parse()
	inputFile := *inputFlag
	outputFile := *outputFlag

	in, err := os.Open(inputFile)
	logError(err)
	defer in.Close()

	out, err := os.Create(outputFile)
	logError(err)
	defer out.Close()

	decoder := wav.NewDecoder(in)
	buf, err := decoder.FullPCMBuffer()
	logError(err)

	// setup the encoder and write all the frames
	e := wav.NewEncoder(out,
		buf.Format.SampleRate,
		int(decoder.BitDepth),
		buf.Format.NumChannels,
		int(decoder.WavAudioFormat))
	defer e.Close()

	rate := buf.Format.SampleRate
	pitch(&buf.Data, 100, rate)
	//reverse(&buf.Data)
	err = e.Write(buf)
	logError(err)
}

func test() {
	arr := []complex128{
		complex(-1, -1),
		complex(1, 1),
		complex(2, 2),
		complex(3, 3),
		complex(4, 4),
		complex(5, 5),
		complex(6, 6),
		complex(5, 5),
		complex(4, 4),
		complex(3, 3),
		complex(2, 2),
		complex(1, 1)}

	sht := shiftRise(arr, 2)

	log.Println(sht)
}
