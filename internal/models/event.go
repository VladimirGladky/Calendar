package models

import "time"

type Event struct {
	UserId  string    `json:"user_id"`
	EventID string    `json:"event_id"`
	Date    time.Time `json:"date"`
	Event   string    `json:"event"`
}
