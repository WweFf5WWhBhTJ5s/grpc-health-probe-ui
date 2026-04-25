package ui

import (
	"testing"
	"time"
)

var baseTime = time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)

func makeAlert(msg string, level AlertLevel, ttl time.Duration) Alert {
	return Alert{Message: msg, Level: level, CreatedAt: baseTime, TTL: ttl}
}

func TestNewAlertQueue_Defaults(t *testing.T) {
	q := NewAlertQueue(0)
	if q.maxSize != 5 {
		t.Fatalf("expected default maxSize 5, got %d", q.maxSize)
	}
}

func TestAlertQueue_Push(t *testing.T) {
	q := NewAlertQueue(3)
	q.Push(makeAlert("a", AlertInfo, time.Minute))
	q.Push(makeAlert("b", AlertWarn, time.Minute))
	if q.Len() != 2 {
		t.Fatalf("expected 2 alerts, got %d", q.Len())
	}
}

func TestAlertQueue_EvictsWhenFull(t *testing.T) {
	q := NewAlertQueue(2)
	q.Push(makeAlert("first", AlertInfo, time.Minute))
	q.Push(makeAlert("second", AlertInfo, time.Minute))
	q.Push(makeAlert("third", AlertInfo, time.Minute))

	if q.Len() != 2 {
		t.Fatalf("expected 2 alerts after eviction, got %d", q.Len())
	}
	if q.Active()[0].Message != "second" {
		t.Errorf("expected oldest surviving alert to be 'second'")
	}
}

func TestAlertExpired(t *testing.T) {
	a := makeAlert("x", AlertError, 5*time.Second)
	if a.Expired(baseTime.Add(3 * time.Second)) {
		t.Error("alert should not be expired yet")
	}
	if !a.Expired(baseTime.Add(6 * time.Second)) {
		t.Error("alert should be expired")
	}
}

func TestAlertQueue_Prune(t *testing.T) {
	q := NewAlertQueue(5)
	q.Push(makeAlert("old", AlertWarn, 1*time.Second))
	q.Push(makeAlert("fresh", AlertInfo, time.Minute))

	q.Prune(baseTime.Add(2 * time.Second))
	if q.Len() != 1 {
		t.Fatalf("expected 1 alert after prune, got %d", q.Len())
	}
	if q.Active()[0].Message != "fresh" {
		t.Errorf("expected 'fresh' alert to survive prune")
	}
}

func TestAlertQueue_ActiveReturnsCopy(t *testing.T) {
	q := NewAlertQueue(5)
	q.Push(makeAlert("msg", AlertInfo, time.Minute))
	copy1 := q.Active()
	copy1[0].Message = "mutated"
	if q.Active()[0].Message == "mutated" {
		t.Error("Active() should return a copy, not a reference")
	}
}
