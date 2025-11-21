package events

import (
	"errors"
	"time"

	"github.com/araddon/dateparse"
	"github.com/weitend/calendar-go/reminder"
	"github.com/weitend/calendar-go/utils"
	"github.com/weitend/calendar-go/validators"
)

type Event struct {
	ID       string             `json:"id"`
	Title    string             `json:"title"`
	StartAt  time.Time          `json:"start_at"`
	Priority Priority           `json:"priority"`
	Reminder *reminder.Reminder `json:"-"`
}

func NewEvent(title string, dateStr string, priority Priority) (*Event, error) {
	isValidTitle := validators.IsValidTitle(title)

	if !isValidTitle {
		return nil, errors.New("неверный формат заголовка")
	}

	t, err := dateparse.ParseAny(dateStr)

	if err != nil {
		return nil, err
	}

	if err := priority.Validate(); err != nil {
		return nil, err
	}

	return &Event{
		ID:       utils.GetNextId(),
		Title:    title,
		StartAt:  t,
		Priority: priority,
		Reminder: nil,
	}, nil
}

func (e *Event) Update(title string, dateStr string, priority Priority) error {
	isValidTitle := validators.IsValidTitle(title)
	time, dateErr := dateparse.ParseAny(dateStr)
	priorityErr := priority.Validate()

	if dateErr != nil {
		return dateErr
	}

	if !isValidTitle {
		return errors.New("некорректный заголовок при изменении события")

	}

	if priorityErr != nil {
		return priorityErr
	}

	e.Title = title
	e.StartAt = time
	e.Priority = priority

	return nil
}
