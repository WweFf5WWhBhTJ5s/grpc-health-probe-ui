package ui

import "time"

// AlertLevel represents the severity of an alert.
type AlertLevel int

const (
	AlertInfo AlertLevel = iota
	AlertWarn
	AlertError
)

// Alert represents a transient notification shown in the UI.
type Alert struct {
	Message   string
	Level     AlertLevel
	CreatedAt time.Time
	TTL       time.Duration
}

// Expired reports whether the alert has outlived its TTL.
func (a Alert) Expired(now time.Time) bool {
	return now.After(a.CreatedAt.Add(a.TTL))
}

// AlertQueue holds a bounded list of active alerts.
type AlertQueue struct {
	alerts  []Alert
	maxSize int
}

// NewAlertQueue creates an AlertQueue with the given capacity.
func NewAlertQueue(maxSize int) *AlertQueue {
	if maxSize <= 0 {
		maxSize = 5
	}
	return &AlertQueue{maxSize: maxSize}
}

// Push adds a new alert, evicting the oldest if at capacity.
func (q *AlertQueue) Push(a Alert) {
	if len(q.alerts) >= q.maxSize {
		q.alerts = q.alerts[1:]
	}
	q.alerts = append(q.alerts, a)
}

// Prune removes all expired alerts relative to now.
func (q *AlertQueue) Prune(now time.Time) {
	active := q.alerts[:0]
	for _, a := range q.alerts {
		if !a.Expired(now) {
			active = append(active, a)
		}
	}
	q.alerts = active
}

// Active returns a copy of the current (non-expired) alerts.
func (q *AlertQueue) Active() []Alert {
	out := make([]Alert, len(q.alerts))
	copy(out, q.alerts)
	return out
}

// Len returns the number of queued alerts.
func (q *AlertQueue) Len() int { return len(q.alerts) }
