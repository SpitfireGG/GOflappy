package game

import (
	"fmt"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PipeProps struct {
	UpperPos float32
	LowerPos float32
	Height   float32
	Width    float32
	Velocity float32
}

var (
	Score   int32
	Highest int32
)

var (
	SpaceBetweenPipes float32 = 110
	Pipes             []PipeProps
	dt                = rl.GetFrameTime()
	BirdPassedThePipe bool
)

type Piper interface {
	DrawPipes(upper, lower rl.Texture2D)
	InitPipes(upper, lower rl.Texture2D) *PipeProps
	UpdatePipePos(upper, lower rl.Texture2D)
	CheckBirdPass(b *Bird) (bool, error)
	SetScoring(xpos, ypos int32)
	Checkcollision()
}

func (pipe *PipeProps) InitPipes(windowHeight int32, upper, lower rl.Texture2D) *PipeProps {
	pipeHeight := float32(rl.GetRandomValue(windowHeight/2, (windowHeight - 200)))
	pipeX := float32(rl.GetScreenWidth()/2 + 200)
	pipeY := 0.0

	gap := rl.GetRandomValue(50, int32(SpaceBetweenPipes))

	pipePorperties := &PipeProps{
		UpperPos: pipeX + float32(gap),
		LowerPos: float32(pipeY) + pipeHeight + float32(gap),
		Height:   pipeHeight,
		Width:    float32(upper.Width),
		Velocity: 3,
	}
	Pipes = append(Pipes, *pipePorperties)
	return pipePorperties
}

func (pipe *PipeProps) DrawPipes(upper, lower rl.Texture2D, windowHeight float32) {
	for _, pipe := range Pipes {
		rl.DrawTexture(upper, int32(pipe.UpperPos), int32(pipe.LowerPos), rl.RayWhite)
		rl.DrawTexture(lower, int32(pipe.UpperPos), int32(pipe.LowerPos+pipe.Height+SpaceBetweenPipes), rl.RayWhite)
	}
}

func (pipe *PipeProps) UpdatePipePos(upper, lower rl.Texture2D) {
	for i, pipe := range Pipes {
		Pipes[i].UpperPos -= pipe.UpperPos + pipe.Velocity*dt

		if pipe.UpperPos <= float32(rl.GetScreenWidth()) {
			Pipes = append(Pipes[:i], Pipes[i+1:]...)
		}
	}
}

func (pipe *PipeProps) CheckBirdPass(b *Bird) bool {
	for _, pipe := range Pipes {
		if *b.BirdPosX >= pipe.UpperPos {
			BirdPassedThePipe = true
		}
	}
	return true
}

func (pipe *PipeProps) SetScoring(xpos, ypos int32) {
	if BirdPassedThePipe {
		Score += 5
		score := strconv.Itoa(int(Score))
		rl.DrawText(fmt.Sprintf("score : %s", score), xpos, ypos, 42, rl.Yellow)
	}
}

func (pipe *PipeProps) Checkcollision() {
}
