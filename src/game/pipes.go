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

var PIPES = []PipeProperties{}
var (
	pipesRemove  = make([]int, 0) // Store indices of PIPES to remove
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

	if !PAUSED {
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

		PIPES = append(PIPES, pipeInstance)
	}
	lastPipeTime = time.Now()

	return PIPES
}

func DrawAndUpdatePipes(upper, lower rl.Texture2D, screenHeight, screenWidth float32) error {
	delta := rl.GetFrameTime()

	if !PAUSED && !GAMEOVER {
		for i := range PIPES {
			rl.DrawTexture(PIPES[i].upper, int32(PIPES[i].upperPosX), int32(upperPosXY), rl.White)
			// rl.DrawRectangle(int32(PIPES[i].upperPosX), int32(upperPosXY), upper.Width, upper.Height, rl.Yellow)
			rl.DrawTexture(PIPES[i].lower, int32(PIPES[i].lowerPosX), int32(PIPES[i].PosY+gapSize-characterFit), rl.White)
			// rl.DrawRectangle(int32(PIPES[i].lowerPosX), int32(PIPES[i].PosY+gapSize-characterFit), lower.Width, lower.Height, rl.Yellow)

			PIPES[i].upperPosX -= velocity*delta + velocity
			PIPES[i].lowerPosX -= velocity*delta + velocity

			if PIPES[i].upperPosX <= -50 || PIPES[i].lowerPosX <= -50 {
				pipesRemove = append(pipesRemove, i) // Store index of pipe to remove
			}
		}
		for i := len(pipesRemove) - 1; i >= 0; i-- {
			PIPES = append(PIPES[:pipesRemove[i]], PIPES[pipesRemove[i]+1:]...)
		}
		pipesRemove = pipesRemove[:0]
	}

	return nil
}

type pipesRec struct {
	X      float32
	Y      float32
	Height float32
	Width  float32
}

/*
CheckCollision checks for collision using the raylib-go builtin functoin
	the rectangle structs for both the PIPES and bird needs to be passed
		the bird struct is in aotheer module , so pass by reference here
			hope it works
*/

// Checks collision between two textures fuck yeah, the collision detection is working well
func CheckCollision(birdRec rl.Rectangle) bool {
	for i := range PIPES {
		if !PAUSED && !GAMEOVER {
			PipesCordsUpper := rl.Rectangle{
				X:      PIPES[i].upperPosX,
				Y:      float32(upperPosXY),
				Height: float32(PIPES[i].upper.Height),
				Width:  float32(PIPES[i].upper.Width),
			}

			PipesCordsLower := rl.Rectangle{
				X:      PIPES[i].lowerPosX,
				Y:      float32(PIPES[i].PosY + gapSize - characterFit),
				Height: float32(PIPES[i].lower.Height),
				Width:  float32(PIPES[i].lower.Width),
			}
			if rl.CheckCollisionRecs(birdRec, PipesCordsUpper) || rl.CheckCollisionRecs(birdRec, PipesCordsLower) {
				return true
			}
		}
	}
	return false
}

/*
inside the main.go where i created the rec for the bird
			BirdsRec := rl.Rectangle{
				X:      *birdCord.BirdPosX,
				Y:      *birdCord.BirdPosY,
				Height: float32(BirdUpTexture.Height),
				Width:  float32(BirdUpTexture.Width),
			}

			collision = game.CheckCollision(BirdsRec)
			if collision {
				fmt.Println("collision detectino")
			}
*/
