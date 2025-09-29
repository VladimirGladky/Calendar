package errors

import "fmt"

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s - %s", e.Field, e.Message)
}

type NotFoundError struct {
	Resource string
	ID       string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with id %s not found", e.Resource, e.ID)
}

type BusinessError struct {
	Message string
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("business error: %s", e.Message)
}
