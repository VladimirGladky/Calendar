package service

import (
	"Calendar/internal/errors"
	"Calendar/internal/models"
	"Calendar/internal/repository"
	"context"
	"github.com/google/uuid"
	"time"
)

type CalendarServiceInterface interface {
	CreateEvent(event *models.Event) (string, error)
	GetEventsForDay(userID string, dateStr string) ([]*models.Event, error)
	GetEventsForWeek(userID string, dateStr string) ([]*models.Event, error)
	GetEventsForMonth(userID string, dateStr string) ([]*models.Event, error)
	DeleteEvent(eventID string) error
	UpdateEvent(event *models.Event) error
}

type CalendarService struct {
	repo repository.CalendarRepositoryInterface
	ctx  context.Context
}

func NewCalendarService(ctx context.Context, repo repository.CalendarRepositoryInterface) *CalendarService {
	return &CalendarService{
		ctx:  ctx,
		repo: repo,
	}
}

func (s *CalendarService) CreateEvent(event *models.Event) (string, error) {
	if event.UserID == "" {
		return "", &errors.ValidationError{
			Field:   "user_id",
			Message: "can't be empty",
		}
	}
	if event.Event == "" {
		return "", &errors.ValidationError{
			Field:   "event",
			Message: "can't be empty",
		}
	}
	_, err := time.Parse(time.DateOnly, event.Date)
	if err != nil {
		return "", &errors.ValidationError{
			Field:   "date",
			Message: "format must be YYYY-MM-DD",
		}
	}
	id := uuid.New().String()
	event.EventID = id
	err = s.repo.CreateEvent(event)
	if err != nil {
		return "", &errors.BusinessError{
			Message: err.Error(),
		}
	}
	return id, nil
}

func (s *CalendarService) GetEventsForDay(userID string, dateStr string) ([]*models.Event, error) {
	if userID == "" {
		return nil, &errors.ValidationError{
			Field:   "user_id",
			Message: "can't be empty",
		}
	}
	if dateStr == "" {
		return nil, &errors.ValidationError{
			Field:   "date",
			Message: "can't be empty",
		}
	}
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, &errors.ValidationError{
			Field:   "date",
			Message: "format must be YYYY-MM-DD",
		}
	}
	events, err := s.repo.GetEventsForDay(userID, date)
	if err != nil {
		return nil, &errors.BusinessError{
			Message: err.Error(),
		}
	}

	if events == nil {
		return []*models.Event{}, nil
	}

	return events, nil
}

func (s *CalendarService) GetEventsForWeek(userID string, dateStr string) ([]*models.Event, error) {
	if userID == "" {
		return nil, &errors.ValidationError{
			Field:   "user_id",
			Message: "can't be empty",
		}
	}
	if dateStr == "" {
		return nil, &errors.ValidationError{
			Field:   "date",
			Message: "can't be empty",
		}
	}
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, &errors.ValidationError{
			Field:   "date",
			Message: "format must be YYYY-MM-DD",
		}
	}
	events, err := s.repo.GetEventsForWeek(userID, date)
	if err != nil {
		return nil, &errors.BusinessError{
			Message: err.Error(),
		}
	}

	if events == nil {
		return []*models.Event{}, nil
	}

	return events, nil
}

func (s *CalendarService) GetEventsForMonth(userID string, dateStr string) ([]*models.Event, error) {
	if userID == "" {
		return nil, &errors.ValidationError{
			Field:   "user_id",
			Message: "user id can't be empty",
		}
	}
	if dateStr == "" {
		return nil, &errors.ValidationError{
			Field:   "date",
			Message: "can't be empty",
		}
	}
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, &errors.ValidationError{
			Field:   "date",
			Message: "format must be YYYY-MM-DD",
		}
	}
	events, err := s.repo.GetEventsForMonth(userID, date)
	if err != nil {
		return nil, &errors.BusinessError{
			Message: err.Error(),
		}
	}

	if events == nil {
		return []*models.Event{}, nil
	}

	return events, nil
}

func (s *CalendarService) DeleteEvent(eventID string) error {
	if eventID == "" {
		return &errors.ValidationError{
			Field:   "event_id",
			Message: "event id can't be empty",
		}
	}
	err := s.repo.DeleteEvent(eventID)
	if err != nil {
		return &errors.BusinessError{
			Message: err.Error(),
		}
	}
	return nil
}

func (s *CalendarService) UpdateEvent(event *models.Event) error {
	if event.EventID == "" {
		return &errors.ValidationError{
			Field:   "event_id",
			Message: "event id can't be empty",
		}
	}
	if event.UserID == "" {
		return &errors.ValidationError{
			Field:   "user_id",
			Message: "user id can't be empty",
		}
	}
	_, err := time.Parse(time.DateOnly, event.Date)
	if err != nil {
		return &errors.ValidationError{
			Field:   "date",
			Message: "format must be YYYY-MM-DD",
		}
	}
	err = s.repo.UpdateEvent(event)
	if err != nil {
		return &errors.BusinessError{
			Message: err.Error(),
		}
	}
	return nil
}
