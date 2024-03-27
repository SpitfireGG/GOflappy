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
	pipesRemove  = make([]int, 0) // Store indices of pipes to remove
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
	for i := range pipes {
		rl.DrawTexture(pipes[i].upper, int32(pipes[i].upperPosX), int32(upperPosXY), rl.White)
		rl.DrawRectangle(int32(pipes[i].upperPosX), int32(upperPosXY), upper.Width, upper.Height, rl.Yellow)
		rl.DrawTexture(pipes[i].lower, int32(pipes[i].lowerPosX), int32(pipes[i].PosY+gapSize-characterFit), rl.White)
		rl.DrawRectangle(int32(pipes[i].lowerPosX), int32(pipes[i].PosY+gapSize-characterFit), lower.Width, lower.Height, rl.Yellow)

		pipes[i].upperPosX -= velocity*delta + velocity
		pipes[i].lowerPosX -= velocity*delta + velocity

		if pipes[i].upperPosX <= -50 || pipes[i].lowerPosX <= -50 {
			pipesRemove = append(pipesRemove, i) // Store index of pipe to remove
		}
	}
	for i := len(pipesRemove) - 1; i >= 0; i-- {
		pipes = append(pipes[:pipesRemove[i]], pipes[pipesRemove[i]+1:]...)
	}
	pipesRemove = pipesRemove[:0]

	return nil
}
