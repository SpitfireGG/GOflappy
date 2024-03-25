package game

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	velocity              float32 = 1
	PIPE_GENERATION_DELAY         = time.Second * 2
	gapSize                       = 300
	characterFit                  = 50
)

var (
	pipes        = []PipeProperties{}
	spawnTimer   = 0
	lastPipeTime time.Time
	upperPosXY   = 0
)

type PipeProperties struct {
	upper     rl.Texture2D
	lower     rl.Texture2D
	upperPosX float32
	lowerPosX float32
	PosY      float32
}

func MakePipes(upper, lower rl.Texture2D, screenHeight, screenWidth float32) []PipeProperties {
	randomHeight := rl.GetRandomValue(130, upper.Height/2)
	pipeX := rl.GetScreenWidth() + 150

	pipeInstance := PipeProperties{
		upper:     upper,
		lower:     lower,
		upperPosX: float32(pipeX),
		lowerPosX: float32(pipeX),
		PosY:      float32(randomHeight),
	}
	if time.Since(lastPipeTime) < PIPE_GENERATION_DELAY {
		return nil
	}

	pipes = append(pipes, pipeInstance)
	lastPipeTime = time.Now()

	return pipes
}

func DrawAndUpdatePipes(upper, lower rl.Texture2D, screenHeight, screenWidth float32) error {
	delta := rl.GetFrameTime()
	for _, pipe := range pipes {
		rl.DrawTexture(pipe.upper, int32(pipe.upperPosX), int32(upperPosXY), rl.White)
		rl.DrawTexture(pipe.lower, int32(pipe.lowerPosX), int32(pipe.PosY+gapSize-characterFit), rl.White)
	}

	for i := range pipes {
		pipes[i].upperPosX -= velocity*delta + velocity
		pipes[i].lowerPosX -= velocity*delta + velocity
	}
	return nil
}
