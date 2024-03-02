package main

import (
	"flappy/src/audio"
	game "flappy/src/game"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1024
	ScreenHeight = 576
	fps          = 120
)

var (
	birdPosX = ScreenWidth / 2
	birdPosY = ScreenHeight / 2

	frameCounts = 0
	GameOver    = false
)

func main() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "GO flappy bird")

	defer rl.CloseWindow()
	defer rl.CloseAudioDevice()

	rl.SetTargetFPS(fps)

	BirdUp := rl.LoadImage("sprites/redbird-upflap.png")
	BirdDown := rl.LoadImage("sprites/redbird-downflap.png")

	pipeUp := rl.LoadImage("sprites/pipe-green.png")
	pipeDown := rl.LoadImage("sprites/pipe-red.png")

	wall := rl.LoadImage("./sprites/wall.png")
	wallTexture := rl.LoadTextureFromImage(wall)

	logo := rl.LoadImage("sprites/logo.png")
	logoTexture := rl.LoadTextureFromImage(logo)

	BirdLogo := rl.LoadImage("./sprites/message.png")
	BlTexture := rl.LoadTextureFromImage(BirdLogo)

	dead := rl.LoadImage("./sprites/gameover.png")
	deadTexture := rl.LoadTextureFromImage(dead)

	BirdUpTexture := rl.LoadTextureFromImage(BirdUp)

	PipeUpTexture := rl.LoadTextureFromImage(pipeUp)
	PipeDownTexture := rl.LoadTextureFromImage(pipeDown)

	defer func() {
		rl.UnloadTexture(BirdUpTexture)
		rl.UnloadTexture(BlTexture)
		rl.UnloadTexture(logoTexture)
		rl.UnloadTexture(wallTexture)

		rl.UnloadTexture(PipeUpTexture)
		rl.UnloadTexture(PipeDownTexture)
		rl.UnloadTexture(deadTexture)
	}()

	var initialPosX float32 = ScreenWidth/2 - float32(BirdUpTexture.Width)/2
	var initialPosY float32 = ScreenHeight/2 - float32(BirdUpTexture.Height)/2 - 40

	var initialX int32 = int32(ScreenWidth / 2)
	var initialY int32 = int32(ScreenHeight / 2)

	birdCord := game.Bird{
		BirdHeight: float32(BirdUpTexture.Height),
		BirdWidth:  float32(BirdUpTexture.Width),
		BirdPosX:   &initialPosX,
		BirdPosY:   &initialPosY,
	}

	/* pipeProps := game.Pipe{
		PipePosX1:  ScreenWidth,
		PipePosX2:  ScreenHeight,
		PipeWidth:  pipeUp.Width,
		PipeHeight: pipeUp.Height,
		Speed:      5,
		PipesGap:   200,
	}
	*/

	game.CurrentState = game.Title

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		switch game.CurrentState {
		case game.Title:
			rl.DrawTexture(logoTexture, 60, -150, rl.RayWhite)
			rl.DrawText("Wait for 2 seconds....", initialX-160, 2*initialY-50, 32, rl.Beige)
			frameCounts++
			if frameCounts > 120 {
				game.CurrentState = game.Menu
			}
		case game.Menu:
			if rl.IsKeyDown(rl.KeyEnter) || rl.IsMouseButtonDown(rl.MouseLeftButton) {
				game.CurrentState = game.EnterGame
			} else {
				rl.DrawTexture(BlTexture, int32(initialX)-BlTexture.Width/2, int32(initialY)-BlTexture.Height/2, rl.RayWhite)
			}
		case game.EnterGame:

			game.Gravity = 0.5
			game.JumpForce = 0.7
			game.BirdVelocity = 0.6
			frameCounts = 0

			flight := func() {
				if rl.IsKeyDown(rl.KeySpace) && !GameOver {

					*birdCord.BirdPosX += game.BirdVelocity
					*birdCord.BirdPosY -= game.JumpForce
					BirdUpTexture = rl.LoadTextureFromImage(BirdUp)

					if audio.FlapSoundCtrl.Paused {
						audio.FlapSoundCtrl.Paused = false
					} else {
						if !audio.FlapSoundCtrl.Paused {
							audio.FlapSoundCtrl.Paused = true
						}
					}

				} else {
					BirdUpTexture = rl.LoadTextureFromImage(BirdDown)
					*birdCord.BirdPosY += game.Gravity
				}

				if *birdCord.BirdPosX >= ScreenWidth || *birdCord.BirdPosY >= ScreenHeight {
					GameOver = true

					game.CurrentState = game.EndGame

				}
			}
			flight()
			audio.LoadFlapSound("./audio/swoosh.ogg")

			rl.DrawTexture(wallTexture, initialX-wallTexture.Width/2, initialY-wallTexture.Height/2, rl.RayWhite)
			rl.DrawTexture(BirdUpTexture, int32(*birdCord.BirdPosX), int32(*birdCord.BirdPosY), rl.White)

		case game.EndGame:

			if rl.IsKeyDown(rl.KeySpace) {

				GameOver = false
				game.Gravity = 0
				game.JumpForce = 0
				game.BirdVelocity = 0
				frameCounts = 0
				birdCord.ResetBirdPos(ScreenWidth, ScreenHeight, birdCord.BirdWidth, birdCord.BirdHeight)
				game.CurrentState = game.Menu
			} else {
				rl.DrawTexture(deadTexture, initialX-deadTexture.Width/2, initialY-deadTexture.Height/2, rl.RayWhite)
				rl.DrawText("Enter space to play again or ESC to quit", 170, 400, 32, rl.Green)
			}

			if rl.IsKeyPressed(rl.KeyEscape) {
				game.CurrentState = game.Menu
			}

		default:
			break

		}
		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}
