package reminder

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Reminder struct {
	Message  string
	At       time.Time
	Duration time.Duration
	IsSended bool
	timer    *time.Timer
}

var ErrEmptyMessage = errors.New("message is empty")

func NewReminder(message string, at time.Time, d time.Duration) (*Reminder, error) {
	if len(strings.TrimSpace(message)) == 0 {
		return nil, fmt.Errorf("can't create reminder: %w", ErrEmptyMessage)
	}

	return &Reminder{
		Message:  message,
		At:       at,
		IsSended: false,
		Duration: d,
	}, nil
}

func (r *Reminder) Start(notify func(msg string)) {
	notifyWrapper := func() {
		if r.IsSended {
			return
		}

		notify(r.Message)

		r.IsSended = true
	}

	r.timer = time.AfterFunc(r.Duration, notifyWrapper)
}

func (r *Reminder) Stop() {
	r.timer.Stop()
}
