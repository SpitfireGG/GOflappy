package game

type GameState int32

const (
	Title GameState = iota
	Menu
	EnterGame
	EndGame
)

var PAUSED bool = false

var CurrentState = Title

var GAMEOVER bool = false
