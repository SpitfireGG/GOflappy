package game

type GameState int32

const (
	Title GameState = iota
	Menu
	EnterGame
	EndGame
)

var CurrentState = Title
