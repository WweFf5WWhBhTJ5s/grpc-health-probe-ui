package probe_test

import (
	"os"
	"testing"

	"github.com/user/grpc-health-probe-ui/internal/probe"
)

func writeTempConfig(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp("", "probe-config-*.json")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Remove(f.Name()) })
	if _, err := f.WriteString(content); err != nil {
		t.Fatal(err)
	}
	f.Close()
	return f.Name()
}

func TestLoadConfig_Valid(t *testing.T) {
	path := writeTempConfig(t, `{
		"interval_seconds": 10,
		"targets": [
			{"name": "my-svc", "host": "localhost:50051", "service": "MyService"}
		]
	}`)

	cfg, err := probe.LoadConfig(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cfg.Targets) != 1 {
		t.Fatalf("expected 1 target, got %d", len(cfg.Targets))
	}
	if cfg.Targets[0].Name != "my-svc" {
		t.Errorf("expected name 'my-svc', got %q", cfg.Targets[0].Name)
	}
	if cfg.IntervalSecs != 10 {
		t.Errorf("expected interval 10, got %d", cfg.IntervalSecs)
	}
}

func TestLoadConfig_MissingFile(t *testing.T) {
	_, err := probe.LoadConfig("/nonexistent/path.json")
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestLoadConfig_NoTargets(t *testing.T) {
	path := writeTempConfig(t, `{"interval_seconds": 5, "targets": []}`)
	_, err := probe.LoadConfig(path)
	if err == nil {
		t.Error("expected validation error for empty targets")
	}
}

func TestLoadConfig_DefaultName(t *testing.T) {
	path := writeTempConfig(t, `{
		"targets": [{"host": "localhost:9090"}]
	}`)
	cfg, err := probe.LoadConfig(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Targets[0].Name != "localhost:9090" {
		t.Errorf("expected name defaulted to host, got %q", cfg.Targets[0].Name)
	}
}

func TestValidate_BadInterval(t *testing.T) {
	cfg := &probe.Config{
		IntervalSecs: -1,
		Targets:      []probe.Target{{Host: "localhost:50051"}},
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for negative interval")
	}
}
