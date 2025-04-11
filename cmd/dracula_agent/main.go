package main

import (
	"os"

	wsdmagent "github.com/seaung/Dracula/internal/WSDMAgent"
)

func main() {
	if err := wsdmagent.NewWsdmAgentCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
