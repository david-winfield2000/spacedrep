package editproblem

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"spacedrep/internal/storage"
	"spacedrep/models"
	"spacedrep/styles"
)

type EditProblemMsg struct {
	Problems []*models.Problem
}

type model struct {
	problems []*models.Problem
	problem  *models.Problem
	inputs   []textinput.Model
	focus    int
}

func InitialModel(problems []*models.Problem, p *models.Problem) model {
	name := textinput.New()
	name.SetValue(p.Name)
	name.Focus()
	name.Placeholder = "Problem Name"

	tag := textinput.New()
	tag.SetValue(p.Tag)
	tag.Placeholder = "Problem Tag"

	url := textinput.New()
	url.SetValue(p.URL)
	url.Placeholder = "Problem URL"

	return model{
		problems: problems,
		problem:  p,
		inputs:   []textinput.Model{name, tag, url},
		focus:    0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "tab":
			m.inputs[m.focus].Blur()
			m.focus = (m.focus + 1) % len(m.inputs)
			m.inputs[m.focus].Focus()

		case "enter":
			m.problem.Name = m.inputs[0].Value()
			m.problem.Tag = m.inputs[1].Value()
			m.problem.URL = m.inputs[2].Value()

			_ = storage.SaveProblems(m.problems)

			return m, func() tea.Msg {
				return "list"
			}

		case "esc":
			return m, func() tea.Msg {
				return "list"
			}
		}
	}

	var cmds []tea.Cmd
	for i := range m.inputs {
		var cmd tea.Cmd
		m.inputs[i], cmd = m.inputs[i].Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	s :=
		styles.HeaderStyle.Render("Edit Problem") + "\n\n" +
			m.inputs[0].View() + "\n" +
			m.inputs[1].View() + "\n" +
			m.inputs[2].View() + "\n\n" +
			styles.HelpStyle.Render("tab switch fields | enter save | esc cancel")

	return styles.Style.Render(s)
}
