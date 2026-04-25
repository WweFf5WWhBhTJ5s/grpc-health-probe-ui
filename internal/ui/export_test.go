package ui

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/user/grpc-health-probe-ui/internal/probe"
)

func makeResults() []probe.Result {
	return []probe.Result{
		{
			Target:    probe.Target{Name: "auth", Address: "localhost:50051"},
			Status:    probe.StatusServing,
			Latency:   12 * time.Millisecond,
			CheckedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Target:    probe.Target{Name: "payment", Address: "localhost:50052"},
			Status:    probe.StatusNotServing,
			Latency:   5 * time.Millisecond,
			CheckedAt: time.Date(2024, 1, 1, 0, 0, 1, 0, time.UTC),
		},
	}
}

func TestBuildSnapshot_RecordCount(t *testing.T) {
	snap := BuildSnapshot(makeResults())
	if len(snap.Records) != 2 {
		t.Fatalf("expected 2 records, got %d", len(snap.Records))
	}
}

func TestBuildSnapshot_FieldMapping(t *testing.T) {
	snap := BuildSnapshot(makeResults())
	r := snap.Records[0]
	if r.Name != "auth" {
		t.Errorf("expected name 'auth', got %q", r.Name)
	}
	if r.Host != "localhost:50051" {
		t.Errorf("expected host 'localhost:50051', got %q", r.Host)
	}
	if r.Status != "SERVING" {
		t.Errorf("expected status 'SERVING', got %q", r.Status)
	}
	if r.LatencyMs != 12 {
		t.Errorf("expected latency_ms 12, got %v", r.LatencyMs)
	}
}

func TestBuildSnapshot_EmptyResults(t *testing.T) {
	snap := BuildSnapshot(nil)
	if len(snap.Records) != 0 {
		t.Errorf("expected empty records, got %d", len(snap.Records))
	}
}

func TestWriteSnapshotJSON_File(t *testing.T) {
	snap := BuildSnapshot(makeResults())
	dir := t.TempDir()
	path := filepath.Join(dir, "out.json")

	if err := WriteSnapshotJSON(snap, path); err != nil {
		t.Fatalf("WriteSnapshotJSON error: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read file error: %v", err)
	}

	var decoded ExportSnapshot
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if len(decoded.Records) != 2 {
		t.Errorf("expected 2 decoded records, got %d", len(decoded.Records))
	}
}

func TestWriteSnapshotJSON_InvalidPath(t *testing.T) {
	snap := BuildSnapshot(makeResults())
	err := WriteSnapshotJSON(snap, "/nonexistent/dir/out.json")
	if err == nil {
		t.Error("expected error for invalid path, got nil")
	}
}
