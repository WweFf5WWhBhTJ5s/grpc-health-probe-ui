package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/example/grpc-health-probe-ui/internal/probe"
)

// PollResultMsg carries a single poll result into the Bubble Tea update loop.
type PollResultMsg probe.PollResult

// TickMsg triggers a refresh of the last-updated timestamp in the view.
type TickMsg time.Time

// WaitForResult returns a Cmd that reads the next result from the poller channel.
func WaitForResult(results <-chan probe.PollResult) tea.Cmd {
	return func() tea.Msg {
		return PollResultMsg(<-results)
	}
}

// Tick returns a Cmd that fires after the given duration.
func Tick(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

// Update handles incoming messages and updates the model accordingly.
// It is intended to be composed into the main Bubble Tea Update method.
func Update(m Model, msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case PollResultMsg:
		m.SetStatus(msg.Target, probe.Status(probe.PollResult(msg).Status))
		m.LastUpdated = time.Now()
		return m, nil

	case TickMsg:
		m.LastUpdated = time.Time(msg)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}
