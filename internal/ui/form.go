package ui

import (
	"46elks-tui/internal/utils"
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// NewModel returns a new Model for the TUI
func NewModel() Model {
	m := Model{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)

	// Create the form and set its fields with validation
	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Sender ID or phone number").
				Validate(utils.ValidateFrom).
				Value(&outgoing.From),

			huh.NewInput().
				Title("Receiver phone number").
				Validate(utils.ValidateTo).
				Value(&outgoing.To),

			huh.NewText().
				Title("Message").
				Validate(utils.ValidateMessage).
				Lines(3).
				Value(&outgoing.Text),

			huh.NewMultiSelect[string]().
				Options(
					huh.NewOption("dryrun", "yes"),
					huh.NewOption("flashsms", "yes"),
				).
				Title("Options").
				Limit(4).
				Value(&options),

			huh.NewConfirm().
				Key("done").
				Title("Are you ready to send?").
				Validate(func(v bool) error {
					if !v {
						return fmt.Errorf("well, finish up then")
					}
					return nil
				}).
				Affirmative("Yes").
				Negative("No"),
		),
	).
		WithShowHelp(false).
		WithShowErrors(false)

	return m
}
