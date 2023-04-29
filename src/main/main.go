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
	reverse    = flag.Bool("r", false, "reverse")
	pitchLow   = flag.Int("pl", 0, "pitch low")
	pitchRise  = flag.Int("pr", 0, "pitch rise")
	temp       = flag.Float64("t", 1.0, "temp")
	gpt        = flag.Float64("gpt", 1.0, "gpt")
)

func applyFlags(buf *[]int) {
	r := *reverse
	pl := *pitchLow
	pr := *pitchRise
	t := *temp
	g := *gpt
	if r {
		audio.Reverse(buf)
	}
	if pl != 0 {
		audio.PitchLow(buf, pl)
	}
	if pr != 0 {
		audio.PitchRise(buf, pr)
	}
	if t != 1.0 {
		audio.Temp(buf, t)
	}
	if g != 1.0 {
		audio.Gpt(buf, g)
	}
}

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

	applyFlags(&buf.Data)
	err = e.Write(buf)
	util.LogError(err)
}
