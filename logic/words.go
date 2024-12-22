package logic

import (
	"math/rand"
	"time"
)

var words = []string{
	"golang", "goroutine", "channel", "struct", "interface",
	"pointer", "package", "import", "defer", "slice",
	"array", "map", "concurrency", "func", "module",
	"context", "sync", "select", "range", "variable",
}

func RandomWord() string {
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}
