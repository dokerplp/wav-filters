package main

import (
	"flag"
	"log"
	"os"

	"main/src/main/audio"
	"main/src/main/util"

	"github.com/go-audio/wav"
)

var (
	inputFlag  = flag.String("i", "in.wav", "input file")
	outputFlag = flag.String("o", "out.wav", "output file")

	reverse = flag.Bool("r", false, "reverse")
	volume  = flag.Int("vol", 100, "volume")
	pitch   = flag.Float64("pitch", 1.0, "pitch shift")
	tempo   = flag.Float64("tempo", 1.0, "increase / decrease tempo")
	shift   = flag.Float64("shift", 1.0, "phase shift")
	pt      = flag.Float64("pt", 1.0, "pitch & tempo")

	low            = flag.Int("low", 0, "smthng low")
	rise           = flag.Int("rise", 0, "smthng rise")
	gpt            = flag.Float64("gpt", 1.0, "chatgpt strange method")
	shuffle        = flag.Bool("sh", false, "shuffle amplitudes")
	raiseAmplitude = flag.Bool("ra", false, "dynamically changing amplitude")
	noise          = flag.Bool("n", false, "make sound noisy")
)

func applyFlags(buf *[]int) {
	if *reverse {
		audio.Reverse(buf)
	}
	if *shuffle {
		audio.Shuffle(buf)
	}
	if *raiseAmplitude {
		audio.RaiseAmplitude(buf)
	}
	if *noise {
		audio.Noise(buf)
	}
	if *tempo != 1.0 {
		audio.Temp(buf, *tempo)
	}
	if *shift != 1.0 {
		audio.Shift(buf, *shift)
	}
	if *pt != 1.0 && *pitch == 1.0 {
		audio.Temp(buf, 1 / *pt)
	}
	if *low != 0 {
		audio.Low(buf, *low)
	}
	if *rise != 0 {
		audio.Rise(buf, *rise)
	}
	if *gpt != 1.0 {
		audio.Gpt(buf, *gpt)
	}
	if *volume != 100 {
		audio.Volume(buf, *volume)
	}
}

func pitchShift() float64 {
	if *pt != 1.0 && *pitch != 1.0 {
		log.Fatal("pt and pitch flags ")
	} else if *pt != 1.0 {
		return *pt
	} else if *pitch != 1.0 {
		return *pitch
	}
	return 1.0
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

	ps := pitchShift()
	newSampleRate := int(float64(buf.Format.SampleRate) * ps)
	e := wav.NewEncoder(out,
		newSampleRate,
		int(decoder.BitDepth),
		buf.Format.NumChannels,
		int(decoder.WavAudioFormat))
	defer e.Close()

	applyFlags(&buf.Data)
	err = e.Write(buf)
	util.LogError(err)
}
