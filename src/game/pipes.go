package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Pipe struct {
	PipePosX1 int32
	PipePosX2 int32

	PipeHeight int32
	PipeWidth  int32

	PipesGap int32
	Speed    int32
}

var Pipes []Pipe

func CreatePipes(pipeUp, pipeDown rl.Texture2D, xpos1, xpos2 int32) {
	rl.DrawTexture(pipeUp, xpos1, xpos2+pipeUp.Height, rl.RayWhite)
	rl.DrawTexture(pipeDown, xpos1, xpos2+pipeDown.Height, rl.RayWhite)
}

/* func DrawPipes(pipesUp, pipesDown rl.Texture2D, pipes []Pipe) {
	for _, pipe := range pipes {
		rl.DrawTexture(pipesUp, int32(pipe.PipePosX1), int32(pipe.PipePosX2-pipe.PipeHeight), rl.White)
		rl.DrawTexture(pipesDown, int32(pipe.PipePosX1), int32(pipe.PipePosX2+pipe.PipesGap), rl.White)
	}
} */
