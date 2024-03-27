package game

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	obs_GENERATION_DELAY = time.Second * 2
	veloc                = 1
)

var removeObs = make([]int, 0)

type Obstacles struct {
	wizposx   float32
	wizposy   float32
	cloudposx float32
	cloudposy float32
}

var OBS []Obstacles

func InitObs(wiz, cloud rl.Texture2D) error {
	if !GAMEOVER && !PAUSED {
		obsatacles := Obstacles{
			wizposx:   float32(rl.GetRandomValue(1200, 1300)),
			wizposy:   float32(rl.GetRandomValue(400, 500)),
			cloudposx: float32(rl.GetRandomValue(1200, 1300)),
			cloudposy: float32(rl.GetRandomValue(10, 500)),
		}

		if time.Since(lastPipeTime) < obs_GENERATION_DELAY {
			return nil
		}
		OBS = append(OBS, obsatacles)
	}

	return nil
}

var delta = rl.GetFrameTime()

func GenerateObs(wiz, cloud rl.Texture2D) {
	if !GAMEOVER && !PAUSED {

		for i := range OBS {
			rl.DrawTexture(cloud, int32(OBS[i].cloudposx), int32(OBS[i].cloudposy), rl.White)
		}
		for i := range OBS {
			OBS[i].wizposx -= veloc
			OBS[i].cloudposx -= veloc*delta + veloc

			if OBS[i].cloudposx <= 0 {
				removeObs = append(removeObs, i)
			}
		}
	}
}
