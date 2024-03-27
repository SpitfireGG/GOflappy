package game

var (
	BirdVelocity float32 = 0.5
	JumpForce    float32 = 0.5
	Gravity      float32 = 0.8
)

type Bird struct {
	BirdPosX *float32
	BirdPosY *float32

	BirdHeight float32
	BirdWidth  float32
}

type Flight interface {
	ResetBirdPos(screenWidth, screenHeight int32, birdWidth, birdHeight float32)
	FlyBird()
}

func (b *Bird) ResetBirdPos(screenWidth, screenHeight int32, birdWidth, birdHeight float32) {
	*b.BirdPosX = float32(screenWidth)/2 - b.BirdWidth
	*b.BirdPosY = float32(screenHeight)/2 - b.BirdHeight - 40
	BirdVelocity = 0.5
	JumpForce = 0.5
	Gravity = 0.8
}
