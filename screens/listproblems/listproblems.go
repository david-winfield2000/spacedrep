package listproblems

import (
	"spacedrep/internal/storage"
	"spacedrep/models"
	"spacedrep/styles"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	problems []*models.Problem
	cursor   int
	showAll  bool
}

type SolveProblemMsg struct {
	Problem *models.Problem
}

type EditProblemMsg struct {
	Problem *models.Problem
}

type DeleteProblemMsg struct {
	Problem *models.Problem
}

func InitialModel(problems []*models.Problem) model {
	return model{
		problems: problems,
		cursor:   0,
		showAll:  false,
	}
}

func (m model) visibleProblems() []*models.Problem {
	if m.showAll {
		return m.problems
	}
	return models.GetDueProblems(m.problems)
}

func (m model) clampCursor() {
	visible := m.visibleProblems()
	if len(visible) == 0 {
		m.cursor = 0
		return
	}
	if m.cursor >= len(visible) {
		m.cursor = len(visible) - 1
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	visible := m.visibleProblems()

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "q":
			return m, tea.Sequence(
				func() tea.Msg {
					_ = storage.CommitPush()
					return nil
				},
				func() tea.Msg {
					return tea.QuitMsg{}
				},
			)

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(visible)-1 {
				m.cursor++
			}

		case "g":
			m.cursor = 0

		case "G":
			if len(visible) > 0 {
				m.cursor = len(visible) - 1
			}

		case "enter":
			if len(visible) > 0 {
				return m, func() tea.Msg {
					return SolveProblemMsg{
						Problem: visible[m.cursor],
					}
				}
			}

		case "d":
			if len(visible) > 0 {
				return m, func() tea.Msg {
					return DeleteProblemMsg{
						Problem: visible[m.cursor],
					}
				}
			}

		case "a":
			return m, func() tea.Msg { return "add" }

		case "e":
			if len(visible) > 0 {
				return m, func() tea.Msg {
					return EditProblemMsg{
						Problem: visible[m.cursor],
					}
				}
			}

		case "t":
			m.showAll = !m.showAll
			m.cursor = 0
		}
	}

	if len(visible) == 0 {
		m.cursor = 0
	} else if m.cursor >= len(visible) {
		m.cursor = len(visible) - 1
	}

	return m, nil
}

func (m model) View() string {
	visible := m.visibleProblems()

	filter := "Due"
	if m.showAll {
		filter = "All"
	}

	s := styles.HeaderStyle.Render("Spaced Repetition") + "\n\n" + filter + " Problems:\n\n"

	if len(visible) == 0 {
		s += "No problems\n"
	} else {
		for i, p := range visible {
			cursor := " "
			if i == m.cursor {
				cursor = ">"
			}
			s += cursor + " " + p.Name + "\n"
		}
	}

	s += styles.HelpStyle.Render("\nj/k move | g jump to top | G jump to bottom | enter solve | a add | e edit | d delete | t toggle filter | q quit")

	return styles.Style.Render(s)
}
