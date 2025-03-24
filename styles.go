package main

import "github.com/charmbracelet/lipgloss"

// define ur styles here

var InfoStl = lipgloss.NewStyle().
	Background(lipgloss.Color("#3d3d3d")). // space gray
	Foreground(lipgloss.Color("white")).   // niga white
	Bold(true).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("white"))

var WarningStl = lipgloss.NewStyle().
	Background(lipgloss.Color("#fad13c")). // candle yellow
	Foreground(lipgloss.Color("white")).   // niga white
	Bold(true).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("white"))

var ErrorStl = lipgloss.NewStyle().
	Background(lipgloss.Color("#fa3c3c")). // angry red
	Foreground(lipgloss.Color("white")).   // niga white
	Bold(true).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("white"))

var DebugStl = lipgloss.NewStyle().
	Background(lipgloss.Color("#3c91fa")). // Sky blue
	Foreground(lipgloss.Color("white")).   // niga white
	Bold(true).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("white"))
