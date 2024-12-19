package main

import (
	"fmt"
	"net/http"

	"github.com/roblesvargas97/hangman/handlers"
)

func main() {

	http.HandleFunc("/api/start", handlers.StartGame)
	http.HandleFunc("/api/guess", handlers.GuessLetter)
	http.HandleFunc("/api/state", handlers.GetGameState)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the hangman API! Use /api/start to start a game")
	})

	fmt.Println("")

}
