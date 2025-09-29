package models

type Event struct {
	UserID  string `json:"user_id"`
	EventID string `json:"event_id"`
	Date    string `json:"date"`
	Event   string `json:"event"`
}
