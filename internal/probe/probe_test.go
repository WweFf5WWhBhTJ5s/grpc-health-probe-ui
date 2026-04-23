package probe_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/user/grpc-health-probe-ui/internal/probe"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func startHealthServer(t *testing.T, serving bool) string {
	t.Helper()
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	healthSvc := health.NewServer()
	if serving {
		healthSvc.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	} else {
		healthSvc.SetServingStatus("", grpc_health_v1.HealthCheckResponse_NOT_SERVING)
	}
	grpc_health_v1.RegisterHealthServer(server, healthSvc)

	go server.Serve(lis) //nolint:errcheck
	t.Cleanup(server.Stop)

	return lis.Addr().String()
}

func TestCheck_Serving(t *testing.T) {
	addr := startHealthServer(t, true)
	p := probe.NewProber()
	result := p.Check(context.Background(), addr, "")

	if result.Status != probe.StatusServing {
		t.Errorf("expected SERVING, got %s (err: %v)", result.Status, result.Error)
	}
	if result.Latency <= 0 {
		t.Error("expected positive latency")
	}
}

func TestCheck_NotServing(t *testing.T) {
	addr := startHealthServer(t, false)
	p := probe.NewProber()
	result := p.Check(context.Background(), addr, "")

	if result.Status != probe.StatusNotServing {
		t.Errorf("expected NOT_SERVING, got %s", result.Status)
	}
}

func TestCheck_Unreachable(t *testing.T) {
	p := probe.NewProber()
	p.DialTimeout = 300 * time.Millisecond
	result := p.Check(context.Background(), "127.0.0.1:1", "")

	if result.Status != probe.StatusUnreachable {
		t.Errorf("expected UNREACHABLE, got %s", result.Status)
	}
	if result.Error == nil {
		t.Error("expected non-nil error for unreachable host")
	}
}

func TestStatusString(t *testing.T) {
	cases := []struct {
		status probe.Status
		want   string
	}{
		{probe.StatusServing, "SERVING"},
		{probe.StatusNotServing, "NOT_SERVING"},
		{probe.StatusUnreachable, "UNREACHABLE"},
		{probe.StatusUnknown, "UNKNOWN"},
	}
	for _, c := range cases {
		if got := c.status.String(); got != c.want {
			t.Errorf("Status(%d).String() = %q, want %q", c.status, got, c.want)
		}
	}
}
