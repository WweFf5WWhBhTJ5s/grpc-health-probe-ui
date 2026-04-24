// main is the entry point for grpc-health-probe-ui, a terminal dashboard
// for monitoring gRPC service health endpoints across multiple hosts.
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/your-org/grpc-health-probe-ui/internal/probe"
	"github.com/your-org/grpc-health-probe-ui/internal/ui"
)

func main() {
	configPath := flag.String("config", "", "Path to config file (default: grpc-health-probe-ui.yaml in current directory)")
	interval := flag.Duration("interval", 10*time.Second, "Poll interval for health checks")
	timeout := flag.Duration("timeout", 5*time.Second, "Timeout per health check request")
	flag.Parse()

	cfg, err := probe.LoadConfig(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading config: %v\n", err)
		os.Exit(1)
	}

	// Apply CLI overrides to config defaults where provided.
	if *interval != 10*time.Second {
		cfg.Interval = *interval
	}
	if *timeout != 5*time.Second {
		cfg.Timeout = *timeout
	}

	prober := probe.NewProber(probe.ProberConfig{
		Timeout: cfg.Timeout,
	})

	poller := probe.NewPoller(cfg.Targets, prober, cfg.Interval)

	model := ui.NewModel(cfg.Targets, poller.Results())

	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),
	)

	// Start polling in the background; it will be stopped when the program exits.
	poller.Start()
	defer poller.Stop()

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error running UI: %v\n", err)
		os.Exit(1)
	}
}
