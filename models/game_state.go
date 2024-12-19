package models

type GameState struct {
	Word     string `json:"word"`
	Mistakes int    `json:"mistakes"`
	MaxTries int    `json:"maxTries"`
	GameWon  bool   `json:"gameWon"`
	GameLost bool   `json:"gameLost"`
	Message  string `json:"message"`
}
