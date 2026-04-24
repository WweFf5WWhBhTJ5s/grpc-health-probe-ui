package probe_test

import (
	"context"
	"testing"
	"time"

	"github.com/example/grpc-health-probe-ui/internal/probe"
)

func TestPoller_ReceivesResults(t *testing.T) {
	srv := startHealthServer(t, true)

	cfg := &probe.Config{
		Targets: []probe.Target{
			{Name: "svc", Address: srv},
		},
	}

	poller := probe.NewPoller(cfg, 50*time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	poller.Start(ctx)
	defer poller.Stop()

	select {
	case result := <-poller.Results():
		if result.Target != srv {
			t.Errorf("expected target %q, got %q", srv, result.Target)
		}
		if result.Status != probe.StatusServing {
			t.Errorf("expected SERVING, got %v", result.Status)
		}
		if result.Err != nil {
			t.Errorf("unexpected error: %v", result.Err)
		}
	case <-ctx.Done():
		t.Fatal("timed out waiting for poll result")
	}
}

func TestPoller_StopsCleanly(t *testing.T) {
	srv := startHealthServer(t, true)

	cfg := &probe.Config{
		Targets: []probe.Target{
			{Name: "svc", Address: srv},
		},
	}

	poller := probe.NewPoller(cfg, 10*time.Millisecond)
	ctx := context.Background()

	poller.Start(ctx)
	time.Sleep(30 * time.Millisecond)
	poller.Stop()
	// Calling Stop twice must not panic.
	poller.Stop()
}

func TestPoller_MultipleTargets(t *testing.T) {
	srv1 := startHealthServer(t, true)
	srv2 := startHealthServer(t, false)

	cfg := &probe.Config{
		Targets: []probe.Target{
			{Name: "a", Address: srv1},
			{Name: "b", Address: srv2},
		},
	}

	poller := probe.NewPoller(cfg, 50*time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	poller.Start(ctx)
	defer poller.Stop()

	seen := map[string]bool{}
	for len(seen) < 2 {
		select {
		case r := <-poller.Results():
			seen[r.Target] = true
		case <-ctx.Done():
			t.Fatalf("timed out; only saw targets: %v", seen)
		}
	}
}
