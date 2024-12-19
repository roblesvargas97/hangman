package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/roblesvargas97/hangman/logic"
)

var game *logic.Game

func StartGame(w http.ResponseWriter, r *http.Request) {

	game = logic.NewGame("golang")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Game started!",
		"word":    game.GetWordState(),
	})

}

func GuessLetter(w http.ResponseWriter, r *http.Request) {

	//Validate if the game is not started
	if game != nil {

		http.Error(w, "Game not started. Call /api/start first.", http.StatusBadRequest)
		return
	}

	letter := r.URL.Query().Get("letter")

	if letter == "" || len(letter) != 1 {
		http.Error(w, "Invalid input. Provide a single letter", http.StatusBadRequest)
		return
	}

	result := game.GuessLetter(letter)

	response := map[string]interface{}{
		"word":     game.GetWordState(),
		"mistakes": game.Mistakes,
		"maxTries": game.MaxMistakes,
		"correct":  result,
		"gameWon":  game.IsWon(),
		"gameLost": game.IsLost(),
	}

	if game.IsWon() {
		response["message"] = "You won! 🎉"
	} else if game.IsLost() {
		response["message"] = "Game over! The word was:" + game.Word
	} else {
		response["message"] = "Keep guessing!"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

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
