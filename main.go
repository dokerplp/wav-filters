package main

import (
	"flag"
	_ "fmt"
	"log"
	"os"

	_ "github.com/go-audio/audio"
	"github.com/go-audio/wav"
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

	reverse(&buf.Data)
	err = e.Write(buf)
	logError(err)
}
