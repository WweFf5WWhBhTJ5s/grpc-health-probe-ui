package ui

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/user/grpc-health-probe-ui/internal/probe"
)

// ExportRecord represents a single target's status snapshot for export.
type ExportRecord struct {
	Name      string    `json:"name"`
	Host      string    `json:"host"`
	Status    string    `json:"status"`
	LatencyMs float64   `json:"latency_ms"`
	Timestamp time.Time `json:"timestamp"`
}

// ExportSnapshot holds all target records at a point in time.
type ExportSnapshot struct {
	ExportedAt time.Time      `json:"exported_at"`
	Records    []ExportRecord `json:"records"`
}

// BuildSnapshot converts a slice of probe results into an ExportSnapshot.
func BuildSnapshot(results []probe.Result) ExportSnapshot {
	records := make([]ExportRecord, 0, len(results))
	for _, r := range results {
		records = append(records, ExportRecord{
			Name:      r.Target.Name,
			Host:      r.Target.Address,
			Status:    r.Status.String(),
			LatencyMs: float64(r.Latency.Milliseconds()),
			Timestamp: r.CheckedAt,
		})
	}
	return ExportSnapshot{
		ExportedAt: time.Now().UTC(),
		Records:    records,
	}
}

// WriteSnapshotJSON serialises snapshot to JSON and writes it to path.
// If path is "-", output is written to stdout.
func WriteSnapshotJSON(snapshot ExportSnapshot, path string) error {
	data, err := json.MarshalIndent(snapshot, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal snapshot: %w", err)
	}
	if path == "-" {
		_, err = fmt.Fprintln(os.Stdout, string(data))
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create export file: %w", err)
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}
