package ui

import (
	"testing"

	"github.com/jhump/grpc-health-probe-ui/internal/probe"
)

func makeFilterTargets() []probe.Target {
	return []probe.Target{
		{Name: "auth-service", Host: "localhost:50051"},
		{Name: "payment-service", Host: "payments.internal:50052"},
		{Name: "user-api", Host: "localhost:50053"},
	}
}

func TestFilterTargets_EmptyQuery(t *testing.T) {
	targets := makeFilterTargets()
	result := FilterTargets(targets, "")
	if len(result) != len(targets) {
		t.Errorf("expected %d targets, got %d", len(targets), len(result))
	}
}

func TestFilterTargets_ByName(t *testing.T) {
	targets := makeFilterTargets()
	result := FilterTargets(targets, "auth")
	if len(result) != 1 || result[0].Name != "auth-service" {
		t.Errorf("expected auth-service, got %+v", result)
	}
}

func TestFilterTargets_ByHost(t *testing.T) {
	targets := makeFilterTargets()
	result := FilterTargets(targets, "payments.internal")
	if len(result) != 1 || result[0].Name != "payment-service" {
		t.Errorf("expected payment-service, got %+v", result)
	}
}

func TestFilterTargets_CaseInsensitive(t *testing.T) {
	targets := makeFilterTargets()
	result := FilterTargets(targets, "AUTH")
	if len(result) != 1 {
		t.Errorf("expected 1 result for case-insensitive match, got %d", len(result))
	}
}

func TestFilterTargets_MultipleMatches(t *testing.T) {
	targets := makeFilterTargets()
	result := FilterTargets(targets, "localhost")
	if len(result) != 2 {
		t.Errorf("expected 2 results for 'localhost', got %d", len(result))
	}
}

func TestFilterTargets_NoMatch(t *testing.T) {
	targets := makeFilterTargets()
	result := FilterTargets(targets, "zzznomatch")
	if len(result) != 0 {
		t.Errorf("expected 0 results, got %d", len(result))
	}
}

func TestFilterQuery_Toggle(t *testing.T) {
	fq := &FilterQuery{}
	fq.Toggle()
	if !fq.Active {
		t.Error("expected Active=true after first toggle")
	}
	fq.Query = "abc"
	fq.Toggle()
	if fq.Active || fq.Query != "" {
		t.Error("expected Active=false and Query='' after second toggle")
	}
}

func TestFilterQuery_AppendAndBackspace(t *testing.T) {
	fq := &FilterQuery{Active: true}
	fq.Append('a')
	fq.Append('b')
	if fq.Query != "ab" {
		t.Errorf("expected 'ab', got %q", fq.Query)
	}
	fq.Backspace()
	if fq.Query != "a" {
		t.Errorf("expected 'a' after backspace, got %q", fq.Query)
	}
}

func TestFilterQuery_BackspaceEmpty(t *testing.T) {
	fq := &FilterQuery{Active: true, Query: ""}
	fq.Backspace() // should not panic
	if fq.Query != "" {
		t.Errorf("expected empty query, got %q", fq.Query)
	}
}
