package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	choices []string
	cursor  int
	status  string
}

func initialModel() model {
	return model{
		choices: []string{
			"Routes",
			"Migrations",
			"Tests",
			"Generators",
			"Logs",
			"Jobs",
			"Console",
			"Command Palette",
		},
		status: "Choose a workflow to wire up next.",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			m.status = fmt.Sprintf("%s selected. Rails command integration comes next.", m.choices[m.cursor])
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	s := "Junction\n"
	s += "Rails development hub\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += fmt.Sprintf("\n%s\n", m.status)
	s += "\nUse j/k or arrows to move. Press enter to select. Press q to quit.\n"

	return tea.NewView(s)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
