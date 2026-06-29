package styles

import (
	lipgloss "github.com/charmbracelet/lipgloss"
)

var Style = lipgloss.NewStyle().
	Padding(1, 2).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#7E22CE")) // purple

var HeaderStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FBBF24")) // gold

var HelpStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#22D3EE")) // cyan

var TagStyle = lipgloss.NewStyle().
	Italic(true).
	Foreground(lipgloss.Color("#777777")) // gray
