package calendar

import (
	"errors"
	"fmt"
	"time"
)

func (c *Calendar) AddEventReminder(id string, message string, at time.Time, d time.Duration) error {
	e := c.calendarEvents[id]

	if e == nil {
		return errors.New("такого события не существует")
	}

	if e.Reminder != nil {
		return errors.New("напоминание у события уже существует")
	}

	err := e.AddReminder(message, at, d, c.Notify)

	if err != nil {
		return fmt.Errorf("can't set reminder: %w", err)
	}

	return nil
}

func (c *Calendar) RemoveEventReminder(id string) {
	e := c.calendarEvents[id]

	if e == nil {
		return
	}

	e.RemoveReminder()
}
