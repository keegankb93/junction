package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

type routesModel struct {
	routes string
}

type routesMsg string

func getRoutes() tea.Msg {
	return routesMsg("These are my routes")
}

func (m routesModel) Init() tea.Cmd {
	return getRoutes
}

func (m routesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case routesMsg:
		m.routes = string(msg)
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m routesModel) View() tea.View {
	if m.routes == "" {
		return tea.NewView("Loading routes...\n")
	}

	return tea.NewView(fmt.Sprintf("%s\n", m.routes))
}
