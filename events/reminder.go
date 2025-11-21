package events

import (
	"fmt"
	"time"

	"github.com/weitend/event-calendar/reminder"
)

func (e *Event) AddReminder(message string, at time.Time, d time.Duration, notify func(msg string)) error {
	r, err := reminder.NewReminder(message, at, d)

	if err != nil {
		return fmt.Errorf("can't add reminder to event: %w", err)
	}

	r.Start(notify)

	return nil
}

func (e *Event) RemoveReminder() {
	e.Reminder.Stop()
	e.Reminder = nil
}
