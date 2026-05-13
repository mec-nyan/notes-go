package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
)

func main() {
	// App setup:
	//
	// Get command line arguments, if any:
	opts, err := parseArgs()
	if err != nil {
		log.Fatalf("initialisation error: %v", err)
	}

	app := tea.NewProgram(initialModel(opts))

	if _, err := app.Run(); err != nil {
		log.Fatalf("Oops! Something went terribly wrong...: %v", err)
	}
}
