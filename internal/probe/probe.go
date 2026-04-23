package probe

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// Status represents the health status of a gRPC service.
type Status int

const (
	StatusUnknown    Status = iota
	StatusServing
	StatusNotServing
	StatusUnreachable
)

func (s Status) String() string {
	switch s {
	case StatusServing:
		return "SERVING"
	case StatusNotServing:
		return "NOT_SERVING"
	case StatusUnreachable:
		return "UNREACHABLE"
	default:
		return "UNKNOWN"
	}
}

// Result holds the result of a single health probe.
type Result struct {
	Host      string
	Service   string
	Status    Status
	Latency   time.Duration
	Error     error
	CheckedAt time.Time
}

// Prober performs gRPC health checks against a target.
type Prober struct {
	DialTimeout  time.Duration
	CheckTimeout time.Duration
}

// NewProber creates a Prober with sensible defaults.
func NewProber() *Prober {
	return &Prober{
		DialTimeout:  2 * time.Second,
		CheckTimeout: 3 * time.Second,
	}
}

// Check performs a gRPC health check against host for the given service name.
func (p *Prober) Check(ctx context.Context, host, service string) Result {
	result := Result{
		Host:      host,
		Service:   service,
		CheckedAt: time.Now(),
	}

	dialCtx, cancel := context.WithTimeout(ctx, p.DialTimeout)
	defer cancel()

	conn, err := grpc.DialContext(dialCtx, host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		result.Status = StatusUnreachable
		result.Error = err
		return result
	}
	defer conn.Close()

	client := grpc_health_v1.NewHealthClient(conn)

	checkCtx, checkCancel := context.WithTimeout(ctx, p.CheckTimeout)
	defer checkCancel()

	start := time.Now()
	resp, err := client.Check(checkCtx, &grpc_health_v1.HealthCheckRequest{Service: service})
	result.Latency = time.Since(start)

	if err != nil {
		result.Status = StatusUnreachable
		result.Error = err
		return result
	}

	switch resp.Status {
	case grpc_health_v1.HealthCheckResponse_SERVING:
		result.Status = StatusServing
	case grpc_health_v1.HealthCheckResponse_NOT_SERVING:
		result.Status = StatusNotServing
	default:
		result.Status = StatusUnknown
	}

	return result
}
