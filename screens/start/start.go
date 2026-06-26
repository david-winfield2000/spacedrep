package start

import (
	"spacedrep/internal/storage"
	"spacedrep/models"
	"spacedrep/styles"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type RepoReadyMsg struct{}

type model struct {
	input  textinput.Model
	config *models.Config
}

func InitialModel(config *models.Config) model {
	ti := textinput.New()
	ti.Placeholder = "Enter github repo URL for backing up your save data"
	ti.Focus()

	return model{
		input:  ti,
		config: config,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "esc":
			return m, tea.Quit

		case "enter":
			m.config.RepoURL = m.input.Value()

			if err := storage.SaveConfig(m.config); err != nil {
				panic(err)
			}

			return m, func() tea.Msg {
				_ = storage.SyncRepo(m.config.RepoURL)
				return "list"
			}
		}
	}

	return m, cmd
}

func (m model) View() string {
	s := styles.HeaderStyle.Render("Welcome to SpacedRep!")
	s += "\n\n"
	s += "Please enter your github repo URL for backing up your save data:\n\n"
	s += m.input.View()
	s += "\n\n"
	s += styles.HelpStyle.Render("enter submit | esc quit\n")
	return styles.Style.Render(s)
}
