package ui

import "time"

// NotifyLevel represents the severity of a notification.
type NotifyLevel int

const (
	NotifyInfo NotifyLevel = iota
	NotifyWarn
	NotifyError
)

// Notification is a timestamped message with a severity level.
type Notification struct {
	Message   string
	Level     NotifyLevel
	CreatedAt time.Time
}

// NotifyStore holds a capped, ordered list of notifications.
type NotifyStore struct {
	entries []Notification
	maxLen  int
}

// NewNotifyStore creates a NotifyStore with the given capacity.
// If maxLen is <= 0, it defaults to 50.
func NewNotifyStore(maxLen int) *NotifyStore {
	if maxLen <= 0 {
		maxLen = 50
	}
	return &NotifyStore{maxLen: maxLen}
}

// Push appends a new notification, evicting the oldest if at capacity.
func (s *NotifyStore) Push(msg string, level NotifyLevel) {
	n := Notification{
		Message:   msg,
		Level:     level,
		CreatedAt: time.Now(),
	}
	if len(s.entries) >= s.maxLen {
		s.entries = s.entries[1:]
	}
	s.entries = append(s.entries, n)
}

// All returns a copy of all stored notifications, oldest first.
func (s *NotifyStore) All() []Notification {
	out := make([]Notification, len(s.entries))
	copy(out, s.entries)
	return out
}

// Len returns the number of stored notifications.
func (s *NotifyStore) Len() int {
	return len(s.entries)
}

// Clear removes all notifications.
func (s *NotifyStore) Clear() {
	s.entries = nil
}

// LevelLabel returns a short string label for a NotifyLevel.
func LevelLabel(l NotifyLevel) string {
	switch l {
	case NotifyWarn:
		return "WARN"
	case NotifyError:
		return "ERR "
	default:
		return "INFO"
	}
}
