package solveproblem

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"spacedrep/internal/browser"
	"spacedrep/internal/storage"
	"spacedrep/models"
	"spacedrep/styles"
)

type model struct {
	problems []*models.Problem
	problem  *models.Problem
	input    textinput.Model
}

func InitialModel(problems []*models.Problem, problem *models.Problem) model {
	ti := textinput.New()
	ti.Placeholder = "0-5"
	ti.CharLimit = 1
	ti.Focus()

	return model{
		problems: problems,
		problem:  problem,
		input:    ti,
	}
}

func (m model) Init() tea.Cmd {
	return func() tea.Msg {
		browser.OpenURL(m.problem.URL)
		return nil
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, func() tea.Msg {
				return "list"
			}

		case "enter":
			switch m.input.Value() {
			case "0", "1", "2", "3", "4", "5":
				ease := int(m.input.Value()[0] - '0')

				models.Review(m.problem, ease)

				if err := storage.SaveProblems(m.problems); err != nil {
					return m, nil
				}

				return m, func() tea.Msg {
					return "list"
				}
			}
		}
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m model) View() string {
	header := styles.HeaderStyle.Render(
		fmt.Sprintf("Problem: %s", m.problem.Name),
	)

	help := "0 = again (reset)\n1 = very hard\n2 = hard\n3 = medium\n4 = easy\n5 = perfect\n\n"

	footer := styles.HelpStyle.Render("enter submit | esc cancel")

	content := header + "\n\n" +
		help +
		m.input.View() +
		"\n\n" +
		footer

	return styles.Style.Render(content)
}
