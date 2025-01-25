package models

type Event struct {
	ID        int64
	UserID    int64
	PageURL   string
	EventType string
	Timestamp string
}
