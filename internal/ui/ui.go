package ui

import (
	"46elks-tui/internal/models"
	"46elks-tui/internal/utils"
	"46elks-tui/pkg/elks"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var (
	red      = lipgloss.AdaptiveColor{Light: "#FE5F86", Dark: "#FE5F86"}
	indigo   = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	green    = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
	outgoing = models.OutgoingSMS{}
	options  []string
)

const (
	statusNormal state = iota
	stateDone
	maxWidth = 80
)

func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}
	s.Base = lg.NewStyle().
		Padding(1, 4, 0, 1)
	s.HeaderText = lg.NewStyle().
		Foreground(indigo).
		Bold(true).
		Padding(0, 1, 0, 2)
	s.Status = lg.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(indigo).
		PaddingLeft(1).
		MarginTop(1)
	s.StatusHeader = lg.NewStyle().
		Foreground(green).
		Bold(true)
	s.Highlight = lg.NewStyle().
		Foreground(lipgloss.Color("212"))
	s.ErrorHeaderText = s.HeaderText.Copy().
		Foreground(red)
	s.Help = lg.NewStyle().
		Foreground(lipgloss.Color("240"))
	return &s
}

// Init initializes the Tea program
func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

// Update updates the TUI, and returns a command to run
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = utils.Minimum(msg.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd

	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	// If the form is completed, send the SMS
	if m.form.State == huh.StateCompleted {
		// Process the options
		processOptions(&outgoing, options)

		// Create 46elks client
		client := elks.CreateClient()

		// Send the SMS
		responseMsg, err := sendSMS(client, &outgoing)
		if err != nil {
			fmt.Println("Error sending SMS:", err)
		} else {
			fmt.Println(responseMsg)
		}

		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}

// View renders the TUI
func (m Model) View() string {
	s := m.styles

	switch m.form.State {
	case huh.StateCompleted:
		// Show the user what they entered and the API response from 46elks
		var b strings.Builder
		_, _ = fmt.Fprintf(&b, "Outgoing:\n%v\n\n", outgoing)
		return s.Status.Copy().Margin(0, 1).Padding(1, 2).Width(48).Render(b.String()) + "\n\n"
	default:
		// Form (left side)
		v := strings.TrimSuffix(m.form.View(), "\n\n")
		form := m.lg.NewStyle().Margin(1, 0).Render(v)

		// Status (right side)
		var status string
		{
			const statusWidth = 28
			statusMarginLeft := m.width - statusWidth - lipgloss.Width(form) - s.Status.GetMarginRight()
			status = s.Status.Copy().
				Height(lipgloss.Height(form)).
				Width(statusWidth).
				MarginLeft(statusMarginLeft).
				Render(s.StatusHeader.Render("Preview") + "\n" +
					outgoing.From +
					outgoing.To +
					outgoing.Text)
		}

		errors := m.form.Errors()
		header := m.appBoundaryView("46Elks TUI")
		if len(errors) > 0 {
			header = m.appErrorBoundaryView(m.errorView())
		}
		body := lipgloss.JoinHorizontal(lipgloss.Top, form, status)

		footer := m.appBoundaryView(m.form.Help().ShortHelpView(m.form.KeyBinds()))
		if len(errors) > 0 {
			footer = m.appErrorBoundaryView("")
		}

		return s.Base.Render(header + "\n" + body + "\n\n" + footer)
	}
}
