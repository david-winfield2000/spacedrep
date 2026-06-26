package main

import (
	"fmt"

	"spacedrep/internal/storage"
	"spacedrep/models"
	"spacedrep/screens/addproblem"
	"spacedrep/screens/deleteproblem"
	"spacedrep/screens/editproblem"
	"spacedrep/screens/listproblems"
	"spacedrep/screens/solveproblem"
	"spacedrep/screens/start"

	tea "github.com/charmbracelet/bubbletea"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type model struct {
	current  tea.Model
	config   *models.Config
	problems []*models.Problem
}

func initialModel(config *models.Config, problems []*models.Problem) model {
	var current tea.Model

	if config.RepoURL == "" {
		current = start.InitialModel(config)
	} else {
		current = listproblems.InitialModel(problems)
	}

	return model{
		current:  current,
		config:   config,
		problems: problems,
	}
}

func (m model) Init() tea.Cmd {
	if m.config.RepoURL == "" {
		return nil
	}

	return func() tea.Msg {
		_ = storage.SyncRepo(m.config.RepoURL)
		return repoSyncedMsg{}
	}
}

type repoSyncedMsg struct{}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case string:
		switch msg {
		case "add":
			m.current = addproblem.InitialModel()
			return m, m.current.Init()

		case "list":
			problems, err := storage.LoadProblems()
			if err != nil {
				return m, nil
			}

			m.problems = problems
			m.current = listproblems.InitialModel(m.problems)
			return m, m.current.Init()
		}

	case repoSyncedMsg:
		problems, _ := storage.LoadProblems()
		m.problems = problems
		m.current = listproblems.InitialModel(m.problems)
		return m, m.current.Init()

	case addproblem.AddProblemMsg:
		m.problems = append(m.problems, msg.Problem)
		check(storage.SaveProblems(m.problems))

		m.current = listproblems.InitialModel(m.problems)
		return m, m.current.Init()

	case listproblems.SolveProblemMsg:
		m.current = solveproblem.InitialModel(m.problems, msg.Problem)
		return m, m.current.Init()

	case listproblems.EditProblemMsg:
		m.current = editproblem.InitialModel(m.problems, msg.Problem)
		return m, m.current.Init()

	case listproblems.DeleteProblemMsg:
		m.current = deleteproblem.InitialModel(m.problems, msg.Problem)
		return m, m.current.Init()

	case deleteproblem.DoneMsg:
		m.problems = msg.Problems
		m.current = listproblems.InitialModel(m.problems)
		return m, nil
	}

	var cmd tea.Cmd
	m.current, cmd = m.current.Update(msg)

	return m, cmd
}

func (m model) View() string {
	return m.current.View()
}

func main() {
	config, err := storage.LoadConfig()
	check(err)

	problems, err := storage.LoadProblems()
	check(err)

	p := tea.NewProgram(
		initialModel(config, problems),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error starting program: %v\n", err)
	}
}
