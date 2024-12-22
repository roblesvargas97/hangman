package main

import (
	"fmt"
	"net/http"

	"github.com/roblesvargas97/hangman/handlers"
)

func main() {
	// Registrar rutas de la API
	http.HandleFunc("/api/start", handlers.StartGame)
	http.HandleFunc("/api/guess", handlers.GuessLetter)
	http.HandleFunc("/api/state", handlers.GetGameState)

	// Ruta principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the hangman API! Use /api/start to start a game")
	})

	// Mensaje indicando que el servidor está corriendo
	fmt.Println("Server running at http://localhost:8080")

	// Iniciar el servidor
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
