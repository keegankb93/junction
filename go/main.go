package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	choices []menuChoice
	current page
	cursor  int
	routes  routesModel
	status  string
}

type page int

const (
	pageHome page = iota
	pageRoutes
)

type menuChoice struct {
	label page
	title string
}

func initialModel() model {
	return model{
		choices: []menuChoice{
			{label: pageRoutes, title: "Routes"},
			{title: "Migrations"},
			{title: "Tests"},
			{title: "Generators"},
			{title: "Logs"},
			{title: "Jobs"},
			{title: "Console"},
			{title: "Command Palette"},
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
		case "esc":
			if m.current != pageHome {
				m.current = pageHome
				m.status = "Choose an option."
				return m, nil
			}
		}
	}

	switch m.current {
	case pageRoutes:
		updated, cmd := m.routes.Update(msg)
		m.routes = updated.(routesModel)
		return m, cmd
	}
	// This doesn't really need to be separate, we should also make handler functions to clean
	// this function up a bit
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			choice := m.choices[m.cursor]
			if choice.label == pageRoutes {
				m.current = pageRoutes
				return m, m.routes.Init()
			}

			m.status = fmt.Sprintf("%s selected.", choice.title)
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	switch m.current {
	case pageRoutes:
		return tea.NewView("Junction / Routes\n\n" + m.routes.View().Content + "\nPress esc to go back. Press q to quit.\n")
	}

	s := "Junction\n"
	s += "Rails development hub\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice.title)
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
