package game

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const obs_GENERATION_DELAY = time.Second * 2

type Obstacles struct {
	wizposx   float32
	wizposy   float32
	cloudposx float32
	cloudposy float32
}

var obs []Obstacles

func InitObs(wiz, cloud rl.Texture2D) ([]Obstacles, error) {
	obsatacles := Obstacles{
		wizposx:   float32(rl.GetRandomValue(1200, 1300)),
		wizposy:   float32(rl.GetRandomValue(400, 500)),
		cloudposx: float32(rl.GetRandomValue(1200, 1300)),
		cloudposy: float32(rl.GetRandomValue(10, 300)),
	}

	if time.Since(lastPipeTime) < obs_GENERATION_DELAY {
		return nil, nil
	}
	obs = append(obs, obsatacles)

	return obs, nil
}

func GenerateObs(wiz, cloud rl.Texture2D) {
	for i := range obs {
		rl.DrawTexture(cloud, int32(obs[i].cloudposx), int32(obs[i].cloudposy), rl.White)
	}
	for i := range obs {
		obs[i].wizposx -= 1
		obs[i].cloudposx -= 1
	}
}
