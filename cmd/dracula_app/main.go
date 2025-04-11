package main

import (
	"os"

	app "github.com/seaung/Dracula/internal/App"
)

func main() {
	if err := app.NewAppCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
