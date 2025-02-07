package glog

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type LogLevel struct {
	Name  string
	style lipgloss.Style
}

type Logger struct {
	Name         string
	CurLogLevel  string
	LevelsToShow []LogLevel
	Output       io.Writer
	FmtConfig    []string
}

func itemExists(slice []LogLevel, item string) bool {
	for _, v := range slice {
		if v.Name == item {
			return true
		}
	}
	return false
}
func (Logger Logger) Println(level LogLevel, msg string) {
	var final_output string
	if itemExists(Logger.LevelsToShow, level.Name) {
		for _, cfg := range Logger.FmtConfig {
			switch cfg {
			case "{level}":
				final_output += level.Name
			case "{msg}":
				final_output += msg
			case "{name}":
				final_output += Logger.Name
			default:
				final_output += cfg
			}
		}

		fmt.Fprintln(Logger.Output, level.style.Render(final_output))

		if strings.Contains(level.Name, "fatal") || strings.Contains(level.Name, "FATAL") {
			os.Exit(1)
		}
	}
}
