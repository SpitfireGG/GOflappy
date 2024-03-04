package game

var BirdPassedPipe = false

type Pipe struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
	Speed  float32
}

var PipeGap int32 = 200
