package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	Options
	Data
}

func initialModel(opts Options) model {
	return model{
		Options: opts,
	}
}

func (m model) Init() tea.Cmd {
	return func() tea.Msg {
		data, err := LoadNotes(m.file)
		return Loader{
			Data:  *data,
			error: err,
		}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case Loader:
		if msg.error != nil {
			panic(fmt.Sprintf("error loading file: %v", msg.error))
		}

		m.Data = msg.Data
		return m, nil

	case tea.KeyPressMsg:

		switch msg.Code {

		case '\x1b', 'q':
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	var s strings.Builder

	s.WriteString("Notes\n\n")

	fmt.Fprintf(&s, "File: %s\n\n", m.file)

	for _, note := range m.Notes {
		fmt.Fprintf(&s, "- %s\n\n", note.Content)
	}

	v := tea.NewView(s.String())

	v.AltScreen = true

	return v
}
