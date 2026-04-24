package probe

import (
	"context"
	"sync"
	"time"
)

// PollResult holds the result of a single health check poll.
type PollResult struct {
	Target string
	Status Status
	Err    error
}

// Poller periodically checks health endpoints and emits results.
type Poller struct {
	prober   *Prober
	targets  []string
	interval time.Duration
	results  chan PollResult
	stopOnce sync.Once
	stop     chan struct{}
}

// NewPoller creates a Poller that checks the given targets at the given interval.
func NewPoller(cfg *Config, interval time.Duration) *Poller {
	targets := make([]string, len(cfg.Targets))
	for i, t := range cfg.Targets {
		targets[i] = t.Address
	}
	return &Poller{
		prober:   NewProber(cfg),
		targets:  targets,
		interval: interval,
		results:  make(chan PollResult, len(targets)),
		stop:     make(chan struct{}),
	}
}

// Results returns the channel on which poll results are delivered.
func (p *Poller) Results() <-chan PollResult {
	return p.results
}

// Start begins polling in the background. Call Stop to halt.
func (p *Poller) Start(ctx context.Context) {
	go func() {
		p.poll(ctx)
		ticker := time.NewTicker(p.interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				p.poll(ctx)
			case <-p.stop:
				return
			case <-ctx.Done():
				return
			}
		}
	}()
}

// Stop halts the poller.
func (p *Poller) Stop() {
	p.stopOnce.Do(func() { close(p.stop) })
}

func (p *Poller) poll(ctx context.Context) {
	var wg sync.WaitGroup
	for _, addr := range p.targets {
		wg.Add(1)
		go func(target string) {
			defer wg.Done()
			status, err := p.prober.Check(ctx, target)
			select {
			case p.results <- PollResult{Target: target, Status: status, Err: err}:
			default:
			}
		}(addr)
	}
	wg.Wait()
}
