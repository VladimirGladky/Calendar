package transport

import (
	"Calendar/internal/config"
	"Calendar/internal/errors"
	"Calendar/internal/models"
	"Calendar/internal/service"
	"Calendar/pkg/logger"
	"context"
	errors1 "errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type CalendarServer struct {
	ctx context.Context
	cfg *config.Config
	srv service.CalendarServiceInterface
}

func NewCalendarServer(ctx context.Context, cfg *config.Config, srv service.CalendarServiceInterface) *CalendarServer {
	return &CalendarServer{
		ctx: ctx,
		cfg: cfg,
		srv: srv,
	}
}

func (s *CalendarServer) Run() error {
	router := gin.Default()
	router.Use(s.Logger())
	logger.GetLoggerFromCtx(s.ctx).Info("gin framework is running")
	api := router.Group("/api/v1")
	{
		api.POST("/create_event", s.createEventHandler())
		api.POST("/update_event", s.updateEventHandler())
		api.POST("/delete_event", s.deleteEventHandler())
		api.GET("/events_for_day", s.getEventsForDayEventHandler())
		api.GET("/events_for_week", s.getEventsForWeekEventHandler())
		api.GET("/events_for_month", s.getEventsForMonthEventHandler())
	}
	return router.Run(s.cfg.Host + ":" + s.cfg.Port)
}

func (s *CalendarServer) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.GetLoggerFromCtx(s.ctx).Info("Middleware:", zap.Any("method:", c.Request.Method), zap.Any("path:", c.Request.URL.Path), zap.Any("time:", time.Now()))
	}
}

func (s *CalendarServer) createEventHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error1"})
				return
			}
		}()
		if c.Request.Method != http.MethodPost {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
			return
		}
		var request *models.Event
		if err := c.ShouldBindJSON(&request); err != nil {
			s.handleError(c, &errors.ValidationError{
				Field:   "request_body",
				Message: "invalid JSON format",
			})
			return
		}
		id, err := s.srv.CreateEvent(request)
		if err != nil {
			s.handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": "Event created successfully", "id": id})
	}
}

func (s *CalendarServer) updateEventHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error1"})
				return
			}
		}()
		if c.Request.Method != http.MethodPost {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
			return
		}
		var request *models.Event
		if err := c.ShouldBindJSON(&request); err != nil {
			s.handleError(c, &errors.ValidationError{
				Field:   "request_body",
				Message: "invalid JSON format",
			})
			return
		}
		err := s.srv.UpdateEvent(request)
		if err != nil {
			s.handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": "Event updated successfully"})
	}
}

func (s *CalendarServer) deleteEventHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error1"})
				return
			}
		}()
		if c.Request.Method != http.MethodPost {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
			return
		}
		var request *models.ID
		if err := c.ShouldBindJSON(&request); err != nil {
			s.handleError(c, &errors.ValidationError{
				Field:   "request_body",
				Message: "invalid JSON format",
			})
			return
		}
		err := s.srv.DeleteEvent(request.ID)
		if err != nil {
			s.handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": "Event deleted successfully"})
	}
}

func (s *CalendarServer) getEventsForDayEventHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error1"})
				return
			}
		}()
		if c.Request.Method != http.MethodGet {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
			return
		}
		userID := c.Query("user_id")
		date := c.Query("date")
		events, err := s.srv.GetEventsForDay(userID, date)
		if err != nil {
			s.handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"events": events})
	}
}

func (s *CalendarServer) getEventsForWeekEventHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error1"})
				return
			}
		}()
		if c.Request.Method != http.MethodGet {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
			return
		}
		userID := c.Query("user_id")
		date := c.Query("date")
		events, err := s.srv.GetEventsForWeek(userID, date)
		if err != nil {
			s.handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"events": events})
	}
}

func (s *CalendarServer) getEventsForMonthEventHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error1"})
				return
			}
		}()
		if c.Request.Method != http.MethodGet {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
			return
		}
		userID := c.Query("user_id")
		date := c.Query("date")
		events, err := s.srv.GetEventsForMonth(userID, date)
		if err != nil {
			s.handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"events": events})
	}
}

type ErrorResponse struct {
	Error   string            `json:"error"`
	Message string            `json:"message,omitempty"`
	Details map[string]string `json:"details,omitempty"`
}

func (s *CalendarServer) handleError(c *gin.Context, err error) {
	var validationErr *errors.ValidationError
	var businessErr *errors.BusinessError

	switch {
	case errors1.As(err, &validationErr):
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: validationErr.Error(),
			Details: map[string]string{validationErr.Field: validationErr.Message},
		})
	case errors1.As(err, &businessErr):
		c.JSON(http.StatusServiceUnavailable, ErrorResponse{
			Error:   "business_error",
			Message: businessErr.Error(),
		})
	default:
		logger.GetLoggerFromCtx(s.ctx).Error("Internal server error", zap.Any("error", err))
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "internal_server_error",
			Message: "An unexpected error occurred",
		})
	}
}
