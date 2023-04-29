package main

import (
	"flag"
	"os"

	"main/src/main/audio"
	"main/src/main/util"

	"github.com/go-audio/wav"
)

var (
	inputFlag  = flag.String("i", "in.wav", "input file")
	outputFlag = flag.String("o", "out.wav", "output file")
)

func main() {
	flag.Parse()
	inputFile := *inputFlag
	outputFile := *outputFlag

	in, err := os.Open(inputFile)
	util.LogError(err)
	defer in.Close()

	out, err := os.Create(outputFile)
	util.LogError(err)
	defer out.Close()

	decoder := wav.NewDecoder(in)
	buf, err := decoder.FullPCMBuffer()
	util.LogError(err)

	e := wav.NewEncoder(out,
		buf.Format.SampleRate,
		int(decoder.BitDepth),
		buf.Format.NumChannels,
		int(decoder.WavAudioFormat))
	defer e.Close()

	rate := buf.Format.SampleRate
	audio.Pitch(&buf.Data, 100, rate)
	//reverse(&buf.Data)
	err = e.Write(buf)
	util.LogError(err)
}
