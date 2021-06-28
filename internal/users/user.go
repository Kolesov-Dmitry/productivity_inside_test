package users

import "users.org/pkg/time"

// User represents a user
type User struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	LastName      string         `json:"last_name"`
	Age           int            `json:"age"`
	RecordingDate time.Timestamp `json:"recording_date"`
}
