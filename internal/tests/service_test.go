package tests

import (
	"Calendar/internal/errors"
	"Calendar/internal/models"
	"Calendar/internal/repository"
	"Calendar/internal/service"
	"context"
	errors1 "errors"
	"testing"
	"time"
)

type MockRepository struct {
	repository.CalendarRepositoryInterface
}

func (m *MockRepository) CreateEvent(event *models.Event) error {
	return nil
}

func TestCalendarService_CreateEvent(t *testing.T) {
	ctx := context.Background()
	repo := &MockRepository{}
	srv := service.NewCalendarService(ctx, repo)

	tests := []struct {
		name  string
		event *models.Event
		err   error
	}{
		{
			name: "valid event",
			event: &models.Event{
				UserID:  "1",
				EventID: "1",
				Event:   "event",
				Date:    "2025-09-29",
			},
			err: nil,
		},
		{
			name: "invalid event1",
			event: &models.Event{
				UserID: "",
				Event:  "event",
				Date:   "2025-09-29",
			},
			err: &errors.ValidationError{
				Field:   "user_id",
				Message: "can't be empty",
			},
		},
		{
			name: "invalid event2",
			event: &models.Event{
				UserID: "1",
				Event:  "",
				Date:   "2025-09-29",
			},
			err: &errors.ValidationError{
				Field:   "event",
				Message: "can't be empty",
			},
		},
		{
			name: "invalid event3",
			event: &models.Event{
				UserID: "1",
				Event:  "event",
				Date:   "2020:31:05",
			},
			err: &errors.ValidationError{
				Field:   "date",
				Message: "format must be YYYY-MM-DD",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.CreateEvent(tt.event)
			if tt.err == nil {
				if err != nil {
					t.Errorf("error = %v, wantErr %v", err, tt.err)
				}
				return
			}
			var target *errors.ValidationError
			if errors1.As(err, &target) {
				if target.Field != tt.err.(*errors.ValidationError).Field ||
					target.Message != tt.err.(*errors.ValidationError).Message {
					t.Errorf("error = %v, wantErr %v", err, tt.err)
				}
			} else {
				t.Errorf("error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}

func (m *MockRepository) DeleteEvent(eventID string) error {
	return nil
}

func TestCalendarService_DeleteEvent(t *testing.T) {
	ctx := context.Background()
	repo := &MockRepository{}
	srv := service.NewCalendarService(ctx, repo)

	tests := []struct {
		name    string
		eventID string
		err     error
	}{
		{
			name:    "valid eventID",
			eventID: "1",
			err:     nil,
		},
		{
			name:    "invalid eventID",
			eventID: "",
			err: &errors.ValidationError{
				Field:   "event_id",
				Message: "event id can't be empty",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := srv.DeleteEvent(tt.eventID)
			if tt.err == nil {
				if err != nil {
					t.Errorf("error = %v, wantErr %v", err, tt.err)
				}
				return
			}

			var target *errors.ValidationError
			if errors1.As(err, &target) {
				if target.Field != tt.err.(*errors.ValidationError).Field ||
					target.Message != tt.err.(*errors.ValidationError).Message {
					t.Errorf("error = %v, wantErr %v", err, tt.err)
				}
			} else {
				t.Errorf("error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}

func (m *MockRepository) UpdateEvent(event *models.Event) error {
	return nil
}

func TestCalendarService_UpdateEvent(t *testing.T) {
	ctx := context.Background()
	repo := &MockRepository{}
	srv := service.NewCalendarService(ctx, repo)

	tests := []struct {
		name  string
		event *models.Event
		err   error
	}{
		{
			name: "valid event",
			event: &models.Event{
				UserID:  "1",
				EventID: "1",
				Event:   "event",
				Date:    "2025-09-29",
			},
			err: nil,
		},
		{
			name: "invalid event1",
			event: &models.Event{
				UserID:  "",
				EventID: "1",
				Event:   "event",
				Date:    "2025-09-29",
			},
			err: &errors.ValidationError{
				Field:   "user_id",
				Message: "user id can't be empty",
			},
		},
		{
			name: "invalid event2",
			event: &models.Event{
				UserID:  "1",
				EventID: "1",
				Event:   "",
				Date:    "2025-09-29",
			},
			err: &errors.ValidationError{
				Field:   "event",
				Message: "can't be empty",
			},
		},
		{
			name: "invalid event3",
			event: &models.Event{
				UserID:  "1",
				EventID: "1",
				Event:   "event",
				Date:    "2020:31:05",
			},
			err: &errors.ValidationError{
				Field:   "date",
				Message: "format must be YYYY-MM-DD",
			},
		},
		{
			name: "invalid event4",
			event: &models.Event{
				UserID:  "1",
				EventID: "",
				Event:   "event",
				Date:    "2020:31:05",
			},
			err: &errors.ValidationError{
				Field:   "event_id",
				Message: "event id can't be empty",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := srv.UpdateEvent(tt.event)
			if tt.err == nil {
				if err != nil {
					t.Errorf("error = %v, wantErr %v", err, tt.err)
				}
				return
			}

			var target *errors.ValidationError
			if errors1.As(err, &target) {
				if target.Field != tt.err.(*errors.ValidationError).Field ||
					target.Message != tt.err.(*errors.ValidationError).Message {
					t.Errorf("error = %v, wantErr %v", err, tt.err)
				}
			} else {
				t.Errorf("error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}

func (m *MockRepository) GetEventsForDay(userID string, date time.Time) ([]*models.Event, error) {
	return nil, nil
}

func TestCalendarService_GetEventsForDay(t *testing.T) {
	ctx := context.Background()
	repo := &MockRepository{}
	srv := service.NewCalendarService(ctx, repo)

	tests := []struct {
		name   string
		userID string
		date   string
		err    error
	}{
		{
			name:   "valid",
			userID: "1",
			date:   "2025-09-29",
			err:    nil,
		},
		{
			name:   "invalid1",
			userID: "1",
			date:   "20250929",
			err: &errors.ValidationError{
				Field:   "date",
				Message: "format must be YYYY-MM-DD",
			},
		},
		{
			name:   "invalid2",
			userID: "",
			date:   "2025-09-29",
			err: &errors.ValidationError{
				Field:   "user_id",
				Message: "can't be empty",
			},
		},
		{
			name:   "invalid3",
			userID: "1",
			date:   "",
			err: &errors.ValidationError{
				Field:   "date",
				Message: "can't be empty",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.GetEventsForDay(tt.userID, tt.date)
			if tt.err == nil {
				if err != nil {
					t.Errorf("error = %v, wantErr %v", err, tt.err)
				}
				return
			}
			var target *errors.ValidationError
			if errors1.As(err, &target) {
				if target.Field != tt.err.(*errors.ValidationError).Field ||
					target.Message != tt.err.(*errors.ValidationError).Message {
					t.Errorf("error = %v, wantErr %v", err, tt.err)
				}
			} else {
				t.Errorf("error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}
