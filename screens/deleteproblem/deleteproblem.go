package deleteproblem

import (
	tea "github.com/charmbracelet/bubbletea"

	"spacedrep/internal/storage"
	"spacedrep/models"
	"spacedrep/styles"
)

type model struct {
	problems []*models.Problem
	problem  *models.Problem
}

type DoneMsg struct {
	Problems []*models.Problem
}

func InitialModel(problems []*models.Problem, problem *models.Problem) model {
	return model{
		problems: problems,
		problem:  problem,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "enter":
			for i, p := range m.problems {
				if p == m.problem {
					m.problems = append(m.problems[:i], m.problems[i+1:]...)
					break
				}
			}

			_ = storage.SaveProblems(m.problems)

			return m, func() tea.Msg {
				return DoneMsg{Problems: m.problems}
			}

		case "esc":
			return m, func() tea.Msg {
				return DoneMsg{Problems: m.problems}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	return styles.Style.Render(
		styles.HeaderStyle.Render("Delete this problem?") +
			"\n\n" +
			m.problem.Name +
			"\n\n" +
			styles.HelpStyle.Render("enter confirm | esc cancel"),
	)
}
