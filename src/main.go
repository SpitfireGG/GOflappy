// PERF: add levels , forgot that shit

package main

import (
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

	Paused = false
)

func main() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "GO flappy bird")
	rl.InitAudioDevice()

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

	pause := rl.LoadImage("sprites/pause.png")
	pauseTexture := rl.LoadTextureFromImage(pause)

	BirdLogo := rl.LoadImage("./sprites/message.png")
	BlTexture := rl.LoadTextureFromImage(BirdLogo)

	dead := rl.LoadImage("./sprites/gameover.png")
	deadTexture := rl.LoadTextureFromImage(dead)

	BirdUpTexture := rl.LoadTextureFromImage(BirdUp)
	// BirdDownTexture := rl.LoadTextureFromImage(BirdDown)

	PipeUpTexture := rl.LoadTextureFromImage(pipeUp)
	PipeDownTexture := rl.LoadTextureFromImage(pipeDown)

	die := rl.LoadSound("./audio/die.ogg")

	defer func() {
		rl.UnloadTexture(BirdUpTexture)
		rl.UnloadTexture(BlTexture)
		rl.UnloadTexture(logoTexture)
		rl.UnloadTexture(wallTexture)

		rl.UnloadTexture(PipeUpTexture)
		rl.UnloadTexture(PipeDownTexture)
		rl.UnloadTexture(deadTexture)
		rl.UnloadSound(die)
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

	game.CurrentState = game.Title
	delta := rl.GetFrameTime()

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

			// var newbvPosY float32
			var pauseYcord float32
			var pauseXcord float32

			isPaused := func() {
				if !Paused && rl.IsKeyPressed(rl.KeyBackspace) {
					Paused = true
				} else if Paused && rl.IsKeyPressed(rl.KeyBackspace) {
					Paused = false
				}
				if Paused {
					rl.DrawTexture(BirdUpTexture, int32(pauseXcord), int32(pauseYcord), rl.RayWhite)
					rl.DrawText("Paused", 50, 50, 40, rl.Red)
				}
			}

			// NOTE: This needs some fixes
			flight := func() {
				if !Paused {

					if rl.IsKeyDown(rl.KeySpace) && !GameOver {
						BirdUpTexture = rl.LoadTextureFromImage(BirdUp)
						*birdCord.BirdPosY -= game.JumpForce*delta + game.JumpForce
						pauseYcord = *birdCord.BirdPosY
					} else {
						*birdCord.BirdPosY += game.Gravity
						BirdUpTexture = rl.LoadTextureFromImage(BirdDown)
					}
					*birdCord.BirdPosX += game.BirdVelocity*delta + game.BirdVelocity
					pauseXcord = *birdCord.BirdPosX
				}
			}
			rl.DrawTexture(BirdUpTexture, int32(*birdCord.BirdPosX), int32(*birdCord.BirdPosY), rl.White)

			if *birdCord.BirdPosX >= float32(rl.GetScreenWidth()-50) || *birdCord.BirdPosY >= float32(rl.GetScreenHeight()-59) {
				GameOver = true
				game.CurrentState = game.EndGame
			}

			// NOTE : These are for bg and bird rendering

			rl.DrawTexture(wallTexture, initialX-wallTexture.Width/2, initialY-wallTexture.Height/2, rl.RayWhite)
			rl.DrawRectangle(int32(initialPosX), int32(initialPosY), BirdUpTexture.Width, BirdUpTexture.Height, rl.Green)
			rl.DrawLine(int32(initialPosX), int32(initialPosY)+10, int32(initialPosX)+100, int32(initialPosY)+10, rl.Red)

			rl.DrawTexture(BirdUpTexture, int32(*birdCord.BirdPosX), int32(*birdCord.BirdPosY), rl.White)

			rl.DrawTexture(pauseTexture, 10, 550, rl.RayWhite)
			rl.DrawRectangle(10, 550, pauseTexture.Width, pauseTexture.Height, rl.RayWhite)

			// fly and pausePlay function
			flight()
			isPaused()

			// FIX: Generate pipes across the screen

		// NOTE: pausing and so on

		// NOTE : add scoring

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
