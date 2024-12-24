package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/roblesvargas97/hangman/logic"
	"github.com/roblesvargas97/hangman/models"
)

var game *logic.Game
var guessChannel = make(chan string)
var responseChannel = make(chan models.GameState)

func init() {

	go func() {

		for letter := range guessChannel {

			if game != nil {

				result := game.GuessLetter(letter)
				gameState := models.GameState{
					Word:     game.GetWordState(),
					Mistakes: game.Mistakes,
					MaxTries: game.MaxMistakes,
					GameWon:  game.IsWon(),
					GameLost: game.IsLost(),
				}

				if game.IsWon() {
					gameState.Message = "You won"
				} else if game.IsLost() {
					gameState.Message = "Game over! The word was: " + game.Word

				} else if result {
					gameState.Message = "Correct Guess!"
				} else {
					gameState.Message = "Incorrect guess. Try again!"
				}

				responseChannel <- gameState

			}
		}
	}()
}

func StartGame(w http.ResponseWriter, r *http.Request) {

	word := logic.RandomWord()
	game = logic.NewGame(word)

	gameState := models.GameState{
		Word:     game.GetWordState(),
		Mistakes: game.Mistakes,
		MaxTries: game.MaxMistakes,
		GameWon:  false,
		GameLost: false,
		Message:  "Game started!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  "success",
		Message: "Game started succesfully!",
		Data:    gameState,
	})

}

func GuessLetter(w http.ResponseWriter, r *http.Request) {
	if game == nil {
		http.Error(w, "Game not started. Call /api/start first", http.StatusBadRequest)
		return
	}

	letter := r.URL.Query().Get("letter")

	if letter == "" || len(letter) != 1 {
		http.Error(w, "Invalid input, Provide a single letter", http.StatusBadRequest)
		return
	}

	guessChannel <- letter

	gameState := <-responseChannel

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  "success",
		Message: "Letter processed successfully!",
		Data:    gameState,
	})

}

func GetGameState(w http.ResponseWriter, r *http.Request) {

	if game == nil {
		http.Error(w, "Game not started. Call /api/start first.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"word":     game.GetWordState(),
		"mistakes": game.Mistakes,
		"maxTries": game.MaxMistakes,
		"gameWon":  game.IsWon(),
		"gameLost": game.IsLost(),
	})

}
