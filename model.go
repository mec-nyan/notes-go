package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type model struct {
	Options
	Data
	Focused int
	Size
}

type Size struct {
	Width  int
	Height int
}

func initialModel(opts Options) model {
	return model{
		Options: opts,
		Size:    Size{80, 25},
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

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

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

		case 'j':
			m.Focused++
			if m.Focused == len(m.Notes) {
				m.Focused = 0
			}

		case 'k':
			m.Focused--
			if m.Focused < 0 {
				m.Focused = len(m.Notes) - 1
			}
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Yellow)

	var s strings.Builder

	s.WriteString(titleStyle.Render("  Notes"))

	noteStyle := lipgloss.NewStyle().
		Width(m.Width-2).
		Foreground(lipgloss.Blue).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Blue).
		Padding(0, 2).
		Margin(1, 1, 0, 1)

	focusedNoteStyle := noteStyle.
		BorderForeground(lipgloss.Green).
		Foreground(lipgloss.Green)

	for i, note := range m.Notes {
		if i == m.Focused {
			s.WriteString(focusedNoteStyle.Render(note.Content))
		} else {
			s.WriteString(noteStyle.Render(note.Content))
		}
	}

	v := tea.NewView(s.String())

	v.AltScreen = true

	return v
}
