package logic

import "strings"

type Game struct {
	Word        string
	Guessed     []string
	Mistakes    int
	MaxMistakes int
}

func NewGame(word string) *Game {
	return &Game{
		Word:        word,
		Guessed:     []string{},
		Mistakes:    0,
		MaxMistakes: 6,
	}
}

func (g *Game) GuessLetter(letter string) bool {
	letter = strings.ToLower(letter)
	if strings.Contains(g.Word, letter) {
		g.Guessed = append(g.Guessed, letter)
		return true
	}
	g.Mistakes++
	return false
}

func (g *Game) GetWordState() string {
	var state strings.Builder
	for _, char := range g.Word {
		if contains(g.Guessed, string(char)) {
			state.WriteRune(char)
		} else {
			state.WriteRune('_')
		}
		state.WriteRune(' ')
	}
	return state.String()
}

func (g *Game) IsWon() bool {
	for _, char := range g.Word {
		if !contains(g.Guessed, string(char)) {
			return false
		}

	}
	return true
}

func (g *Game) IsLost() bool {
	return g.Mistakes >= g.MaxMistakes
}

func contains(slice []string, letter string) bool {
	for _, l := range slice {
		if l == letter {
			return true
		}
	}
	return false
}
