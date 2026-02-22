package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("220"))
	docStyle    = lipgloss.NewStyle().Margin(1, 2)
	notesDir    string
)

