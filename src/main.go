// PERF: 60 - 55 constant fps
// NOTE :add levels and stages

package main

import (
	game "flappy/src/game"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth        = 1024
	ScreenHeight       = 576
	fps                = 60
	interpolationSpeed = 0.1
)

var (
	collision = false
	birdPosX  = ScreenWidth / 2
	birdPosY  = ScreenHeight / 2

	frameCounts = 0

	delta               = rl.GetFrameTime()
	interpolationFactor = delta * interpolationSpeed
)

func main() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "GO flappy bird")

	defer rl.CloseWindow()

	rl.SetTargetFPS(fps)

	BirdUp := rl.LoadImage("sprites/redbird-upflap.png")
	BirdDown := rl.LoadImage("sprites/redbird-downflap.png")

	pipeUp := rl.LoadImage("sprites/pipe-green.png")
	pipeDown := rl.LoadImage("sprites/pipe_down.png")

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

	attack := rl.LoadTextureFromImage(rl.LoadImage("./sprites/attacker.png"))
	cloud := rl.LoadTextureFromImage(rl.LoadImage("./sprites/cloud9.png"))

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

	birdX := Lerp(initialPosX, *birdCord.BirdPosX, interpolationSpeed)
	birdY := Lerp(initialPosY, *birdCord.BirdPosY, interpolationSpeed)

	birdCord.BirdPosX = &birdX
	birdCord.BirdPosY = &birdY

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

			// NOTE: pausing and so on, tried saving the cureeent pos but doesnot seem to need

			isPaused := func() {
				if !game.PAUSED && rl.IsKeyPressed(rl.KeyBackspace) {
					game.PAUSED = true
				} else if game.PAUSED && rl.IsKeyPressed(rl.KeyBackspace) {
					game.PAUSED = false
				}
				if game.PAUSED {
					rl.DrawText("paused", 50, 50, 40, rl.Red)
				}
			}

			// NOTE: This function needs some fixes, choppy movements
			// FIX : Added linear interpolation to make the bird movement even smoother

			flight := func() {
				if !game.PAUSED {
					if rl.IsKeyDown(rl.KeySpace) && !game.GAMEOVER {
						BirdUpTexture = rl.LoadTextureFromImage(BirdUp)
						*birdCord.BirdPosY -= game.JumpForce*delta + game.JumpForce
						*birdCord.BirdPosX += game.BirdVelocity * delta

					} else {
						*birdCord.BirdPosY += game.Gravity

						BirdUpTexture = rl.LoadTextureFromImage(BirdDown)
					}
				}
			}
			rl.DrawTexture(BirdUpTexture, int32(*birdCord.BirdPosX), int32(*birdCord.BirdPosY), rl.White)

			// NOTE : gameEnd
			if *birdCord.BirdPosX >= float32(rl.GetScreenWidth()-50) || *birdCord.BirdPosY >= float32(rl.GetScreenHeight()-59) {
				game.GAMEOVER = true
				game.CurrentState = game.EndGame
			}

			// NOTE : rendering and stuffs
			rl.DrawTexture(wallTexture, initialX-wallTexture.Width/2, initialY-wallTexture.Height/2, rl.RayWhite)
			// rl.DrawRectangle(int32(*birdCord.BirdPosX), int32(*birdCord.BirdPosY), BirdUpTexture.Width, BirdUpTexture.Height, rl.Blue)
			rl.DrawTexture(BirdUpTexture, int32(*birdCord.BirdPosX), int32(*birdCord.BirdPosY), rl.White)

			if !game.PAUSED {
				// FIXME: Generate pipes across the screen

				game.InitObs(attack, cloud)
				game.GenerateObs(attack, cloud)

				game.MakePipes(PipeDownTexture, PipeUpTexture, ScreenHeight, ScreenWidth)
				game.DrawAndUpdatePipes(PipeDownTexture, PipeUpTexture, ScreenHeight, ScreenWidth)
			}

			flight()
			isPaused()

			// check for collision
			BirdsRec := rl.Rectangle{
				X:      *birdCord.BirdPosX,
				Y:      *birdCord.BirdPosY,
				Height: float32(BirdUpTexture.Height),
				Width:  float32(BirdUpTexture.Width),
			}
			collision = game.CheckCollision(BirdsRec)
			if collision {
				rl.DrawText("COLLISION detected!", 10, 200, 24, rl.Black)
				game.GAMEOVER = true
				game.CurrentState = game.EndGame
			}

		case game.EndGame:
			if rl.IsKeyDown(rl.KeySpace) {

				game.GAMEOVER = false
				game.PAUSED = false
				game.Gravity = 0
				game.JumpForce = 0
				game.BirdVelocity = 0
				frameCounts = 0
				game.PIPES = nil
				game.OBS = nil

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

// the interpolation Factorial needs to be 0 to 1
func Lerp(start, end, factor float32) float32 {
	return start + (end-start)*factor
}
