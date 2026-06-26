package addproblem

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"spacedrep/models"
	"spacedrep/styles"
)

type AddProblemMsg struct {
	Problem *models.Problem
}

type model struct {
	inputs []textinput.Model
	focus  int
}

func InitialModel() model {
	name := textinput.New()
	name.Placeholder = "Problem Name"
	name.Focus()

	url := textinput.New()
	url.Placeholder = "Problem URL"

	return model{
		inputs: []textinput.Model{name, url},
		focus:  0,
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
			name := m.inputs[0].Value()
			url := m.inputs[1].Value()

			return m, func() tea.Msg {
				return AddProblemMsg{
					Problem: models.NewProblem(name, url),
				}
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
		styles.HeaderStyle.Render("Add New Problem") + "\n\n" +
			m.inputs[0].View() + "\n" +
			m.inputs[1].View() + "\n\n" +
			styles.HelpStyle.Render("tab switch fields | enter save | esc cancel")

	return styles.Style.Render(s)
}
