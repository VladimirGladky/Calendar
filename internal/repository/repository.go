package repository

import (
	"Calendar/internal/models"
	"Calendar/pkg/logger"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"time"
)

type CalendarRepositoryInterface interface {
	CreateEvent(event *models.Event) error
	GetEventsForDay(userID string, date time.Time) ([]*models.Event, error)
	GetEventsForWeek(userID string, date time.Time) ([]*models.Event, error)
	GetEventsForMonth(userID string, date time.Time) ([]*models.Event, error)
	DeleteEvent(eventId string) error
	UpdateEvent(event *models.Event) error
}

type CalendarRepository struct {
	ctx context.Context
	db  *pgx.Conn
}

func NewCalendarRepository(ctx context.Context, db *pgx.Conn) *CalendarRepository {
	return &CalendarRepository{
		ctx: ctx,
		db:  db,
	}
}

func (r *CalendarRepository) CreateEvent(event *models.Event) error {
	date, _ := time.Parse(time.DateOnly, event.Date)
	_, err := r.db.Exec(r.ctx,
		"INSERT INTO events (event_id, user_id,event, date) "+
			"VALUES ($1, $2, $3, $4)",
		event.EventID,
		event.UserID,
		event.Event,
		date,
	)
	if err != nil {
		logger.GetLoggerFromCtx(r.ctx).Error("error creating event", zap.Error(err))
		return fmt.Errorf("error creating event: %w", err)
	}
	return nil
}

func (r *CalendarRepository) GetEventsForDay(userID string, date time.Time) ([]*models.Event, error) {
	var events []*models.Event

	rows, err := r.db.Query(r.ctx,
		"SELECT event_id, user_id, event, date "+
			"FROM events WHERE user_id = $1 AND date = $2",
		userID,
		date,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting events for day: %w", err)
	}
	defer rows.Close()

	var d time.Time

	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.EventID, &event.UserID, &event.Event, &d)
		if err != nil {
			return nil, err
		}
		event.Date = d.Format(time.DateOnly)
		events = append(events, &event)
	}

	return events, nil
}

func (r *CalendarRepository) GetEventsForWeek(userID string, date time.Time) ([]*models.Event, error) {
	var events []*models.Event
	endDate := date.AddDate(0, 0, 6).Add(23*time.Hour + 59*time.Minute + 59*time.Second)
	rows, err := r.db.Query(r.ctx,
		"SELECT user_id, event_id, date, event "+
			"FROM events WHERE user_id = $1 AND date BETWEEN $2 AND $3",
		userID,
		date,
		endDate,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting events for week: %w", err)
	}
	var d time.Time
	defer rows.Close()
	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.UserID, &event.EventID, &d, &event.Event)
		if err != nil {
			return nil, err
		}
		event.Date = d.Format(time.DateOnly)
		events = append(events, &event)
	}
	return events, nil
}

func (r *CalendarRepository) GetEventsForMonth(userID string, date time.Time) ([]*models.Event, error) {
	var events []*models.Event
	endDate := date.AddDate(0, 1, 0)
	rows, err := r.db.Query(r.ctx,
		"SELECT user_id, event_id, date, event "+
			"FROM events WHERE user_id = $1 AND date BETWEEN $2 AND $3",
		userID,
		date,
		endDate,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting events for week: %w", err)
	}
	var d time.Time
	defer rows.Close()
	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.UserID, &event.EventID, &d, &event.Event)
		if err != nil {
			return nil, err
		}
		event.Date = d.Format(time.DateOnly)
		events = append(events, &event)
	}
	return events, nil
}

func (r *CalendarRepository) DeleteEvent(eventID string) error {
	res, err := r.db.Exec(r.ctx,
		"DELETE FROM events WHERE event_id = $1",
		eventID,
	)
	if err != nil {
		logger.GetLoggerFromCtx(r.ctx).Error("error deleting event", zap.Any("error:", err))
		return fmt.Errorf("error deleting event: %w", err)
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("no event found with id: %s", eventID)
	}
	return nil
}

func (r *CalendarRepository) UpdateEvent(event *models.Event) error {
	res, err := r.db.Exec(r.ctx,
		"UPDATE events SET user_id = $1, event_id = $2, event = $3, date = $4 WHERE event_id = $5",
		event.UserID,
		event.EventID,
		event.Event,
		event.Date,
		event.EventID,
	)
	if err != nil {
		logger.GetLoggerFromCtx(r.ctx).Error("error updating event", zap.Any("error:", err))
		return fmt.Errorf("error updating event: %w", err)
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("no event found with id: %s", event.EventID)
	}
	return nil
}
