package main

import (
	"math/rand"
	"time"

	ray "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1024
	screenHeight = 600
	fps          = 120
	pipeWidth    = 100
	pipeGap      = 200
)

type GameState int

const (
	Title GameState = iota
	Menu
	Start
	End
)

var (
	gravity      float32 = 0.5
	jumpForce    float32 = 0.8
	BirdVelocity float32 = 0.5

	Overmessage string = "The Game is Over"

	Score    float32       = 0
	Best     float32       = 0
	Playtime time.Duration = 0
	GameOver               = false

	frameCounts float32 = 0

	PipesVelocity float32 = 0.7
	PipeHit       bool    = false
)

type Bird struct {
	BirdPosX *float32
	BirdPosY *float32

	BirdHeight int32
	birdWidth  int32
}

type Pipe struct {
	PosX   float32
	PosY   float32
	Height float32
}

func drawLine() {
	for i := 0; i < screenWidth; i += 50 {
		ray.DrawLine(int32(i), int32(screenHeight)-40, int32(i+50), int32(screenHeight)-40, ray.Black)
	}
}

type Flight interface {
	ResetBirdPos()
	BirdMovement()
	BirdDead()
}

// ResetBirdPos resets  the bird position after the game is Over
func (r *Bird) ResetBirdPos() {
	*r.BirdPosX = screenWidth/2 - float32(r.birdWidth)/2
	*r.BirdPosY = screenHeight/2 - float32(r.BirdHeight)/2 - 40
}

// BirdDead does renders something after the game is over
func (r *Bird) BirdDead(birdTexture ray.Texture2D, birdCord *Bird) {
	if GameOver {
		ray.DrawText(Overmessage, 0, 576, 60, ray.Green)
		BirdVelocity = 0
		PipesVelocity = 0
		gravity = 0.0
	} else {
		*birdCord.BirdPosX += BirdVelocity
	}
}

// DrawPipes draws pipes
func DrawPipes(pipesDown, pipesUp ray.Texture2D, pipes []Pipe) {
	for _, pipe := range pipes {
		ray.DrawTexture(pipesUp, int32(pipe.PosX), int32(pipe.PosY-pipe.Height), ray.White)
		ray.DrawTexture(pipesDown, int32(pipe.PosX), int32(pipe.PosY+pipeGap), ray.White)
	}
}

func main() {
	ray.InitWindow(int32(screenWidth), int32(screenHeight), "FlappyBird")
	defer ray.CloseWindow()

	ray.SetTargetFPS(fps)

	birdUp := ray.LoadImage("./assets/flappy-bird-assets/sprites/redbird-upflap.png")
	birdDown := ray.LoadImage("./assets/flappy-bird-assets/sprites/redbird-downflap.png")

	pipeUp := ray.LoadImage("./assets/flappy-bird-assets/sprites/pipeup.png")
	pipeDown := ray.LoadImage("./assets/flappy-bird-assets/sprites/pipedown.png")

	LightBg := ray.LoadImage("./assets/flappy-bird-assets/sprites/LightBg.jpg")
	DarkBg := ray.LoadImage("./assets/flappy-bird-assets/sprites/else.jpg")

	base := ray.LoadImage("./assets/flappy-bird-assets/sprites/base.png")

	message := ray.LoadImage("./assets/flappy-bird-assets/sprites/message.png")
	gameover := ray.LoadImage("./assets/flappy-bird-assets/sprites/gameover.png")

	birdTexture := ray.LoadTextureFromImage(birdUp)
	StartMessage := ray.LoadTextureFromImage(message)
	EndMessage := ray.LoadTextureFromImage(gameover)

	pipeUpTexture := ray.LoadTextureFromImage(pipeUp)
	pipeDownTexture := ray.LoadTextureFromImage(pipeDown)

	background := ray.LoadTextureFromImage(LightBg)
	darkbg := ray.LoadTextureFromImage(DarkBg)

	var DynamicBg ray.Texture2D

	defer func() {
		ray.UnloadTexture(birdTexture)
		ray.UnloadTexture(pipeUpTexture)
		ray.UnloadTexture(pipeDownTexture)
		ray.UnloadImage(birdUp)
		ray.UnloadImage(birdDown)
		ray.UnloadImage(pipeUp)
		ray.UnloadImage(pipeDown)
		ray.UnloadImage(base)
		ray.UnloadImage(LightBg)
		ray.UnloadImage(DarkBg)
	}()

	var initialPosX float32 = screenWidth/2 - float32(birdTexture.Width)/2
	var initialPosY float32 = screenHeight/2 - float32(birdTexture.Height)/2 - 40

	var initialX int32 = int32(screenWidth/2) - StartMessage.Width/2
	var initialY int32 = int32(screenHeight/2) - StartMessage.Height/2

	birdCord := Bird{
		BirdPosX: &initialPosX,
		BirdPosY: &initialPosY,

		BirdHeight: birdTexture.Height,
		birdWidth:  birdTexture.Width,
	}

	var pipes []Pipe
	currentState := Title
	DynamicBg = background

	for !ray.WindowShouldClose() {

		ray.BeginDrawing()
		ray.ClearBackground(ray.RayWhite)

		switch currentState {

		case Title:
			ray.DrawText("PRESS ENTER or TAP to JUMP to GAMEPLAY SCREEN", 120, 220,
				20, ray.DarkGreen)
			time.Sleep(time.Duration(frameCounts))
			frameCounts++
			if frameCounts > 120 {
				currentState = Menu
			}

		case Menu:
			if ray.IsKeyDown(ray.KeyEnter) || ray.IsMouseButtonDown(ray.MouseLeftButton) {
				currentState = Start
			} else {
				ray.DrawTexture(StartMessage, int32(initialX), int32(initialY), ray.RayWhite)
			}

		case Start:
			flight := func() {
				if ray.IsKeyDown(ray.KeySpace) && !GameOver {

					birdTexture = ray.LoadTextureFromImage(birdUp)

					*birdCord.BirdPosY -= BirdVelocity

					birdTexture = ray.LoadTextureFromImage(birdUp)
				} else {
					birdTexture = ray.LoadTextureFromImage(birdDown)

					*birdCord.BirdPosY += gravity
				}
				Playtime++

				if Playtime >= time.Second*1 {
					DynamicBg = darkbg
				} else {
					DynamicBg = background
				}
				ray.DrawTexture(DynamicBg, 0, 0, ray.RayWhite)

				ray.DrawTexture(birdTexture, int32(*birdCord.BirdPosX), int32(*birdCord.BirdPosY), ray.White)
				DrawPipes(pipeUpTexture, pipeDownTexture, pipes)
			}
			flight()

			if len(pipes) == 0 || pipes[len(pipes)-1].PosX <= screenWidth-pipeWidth*2 {

				pipeHeight := int32(rand.Intn(screenHeight/2)) + screenHeight/4
				pipes = append(pipes, Pipe{PosX: screenWidth, PosY: float32(pipeHeight), Height: float32(pipeHeight)})
			}

			// Update pipes position
			for i := range pipes {
				pipes[i].PosX -= PipesVelocity
			}

			// Check collision with pipes
			for _, pipe := range pipes {
				if *birdCord.BirdPosX+float32(birdCord.birdWidth) >= pipe.PosX && *birdCord.BirdPosX <= pipe.PosX+float32(pipeWidth) {
					if *birdCord.BirdPosY <= pipe.PosY || *birdCord.BirdPosY+float32(birdCord.BirdHeight) >= pipe.PosY+pipeGap {
						GameOver = true
					} else if *birdCord.BirdPosY >= float32(ray.GetScreenHeight()) || *birdCord.BirdPosX >= float32(ray.GetScreenWidth()) {
						GameOver = true
					}
				}
			}

			if GameOver {
				currentState = End
			}

		case End:
			if ray.IsKeyPressed(ray.KeyEnter) {
				GameOver = false
				Score = 0
				Best = 0
				frameCounts = 0
				pipes = nil
				Playtime = 0
				birdCord.ResetBirdPos()
				currentState = Menu
			} else {
				ray.DrawTexture(EndMessage, initialX, initialY, ray.RayWhite)
			}
		default:
			break
		}

		ray.DrawFPS(10, 10)

		ray.EndDrawing()
	}
}
