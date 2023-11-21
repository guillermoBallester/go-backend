package main

import (
	"github.com/gsasso/go-backend/src/server/internal/ticker"
)

func main() {

	ticker.Ticker()
	app := InitializeApp()
	if app != nil {
		app.Start()
	}

}
