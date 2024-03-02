package audio

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis" // Import the vorbis package for OGG files
)

var FlapSoundCtrl *beep.Ctrl

func LoadFlapSound(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	streamer, format, err := vorbis.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal(err)
	}

	FlapSoundCtrl = &beep.Ctrl{Streamer: beep.Loop(1, streamer), Paused: true}
	speaker.Play(FlapSoundCtrl)
}
