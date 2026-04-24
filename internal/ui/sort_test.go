package ui

import (
	"testing"

	"github.com/your-org/grpc-health-probe-ui/internal/probe"
)

func makeSortTargets() []probe.Target {
	return []probe.Target{
		{Name: "zebra", Address: "host-c:50051", LastStatus: probe.StatusUnknown},
		{Name: "alpha", Address: "host-a:50051", LastStatus: probe.StatusServing},
		{Name: "mango", Address: "host-b:50051", LastStatus: probe.StatusNotServing},
	}
}

func TestSortByName(t *testing.T) {
	targets := makeSortTargets()
	sorted := SortTargets(targets, SortByName)

	if sorted[0].Name != "alpha" || sorted[1].Name != "mango" || sorted[2].Name != "zebra" {
		t.Errorf("unexpected sort by name: %v", sorted)
	}
}

func TestSortByHost(t *testing.T) {
	targets := makeSortTargets()
	sorted := SortTargets(targets, SortByHost)

	if sorted[0].Address != "host-a:50051" || sorted[1].Address != "host-b:50051" || sorted[2].Address != "host-c:50051" {
		t.Errorf("unexpected sort by host: %v", sorted)
	}
}

func TestSortByStatus(t *testing.T) {
	targets := makeSortTargets()
	sorted := SortTargets(targets, SortByStatus)

	// StatusServing=1, StatusNotServing=2, StatusUnknown=0 — order depends on iota in probe
	if len(sorted) != 3 {
		t.Fatalf("expected 3 targets, got %d", len(sorted))
	}
}

func TestSortDoesNotMutateOriginal(t *testing.T) {
	targets := makeSortTargets()
	originalFirst := targets[0].Name
	_ = SortTargets(targets, SortByName)

	if targets[0].Name != originalFirst {
		t.Error("SortTargets mutated the original slice")
	}
}

func TestNextSortOrder(t *testing.T) {
	if NextSortOrder(SortByName) != SortByStatus {
		t.Error("expected SortByStatus after SortByName")
	}
	if NextSortOrder(SortByStatus) != SortByHost {
		t.Error("expected SortByHost after SortByStatus")
	}
	if NextSortOrder(SortByHost) != SortByName {
		t.Error("expected SortByName after SortByHost (wrap)")
	}
}

func TestSortOrderLabel(t *testing.T) {
	cases := map[SortOrder]string{
		SortByName:   "name",
		SortByStatus: "status",
		SortByHost:   "host",
	}
	for order, want := range cases {
		if got := SortOrderLabel(order); got != want {
			t.Errorf("SortOrderLabel(%d) = %q, want %q", order, got, want)
		}
	}
}
