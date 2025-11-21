package calendar

import (
	"errors"

	"github.com/weitend/calendar-go/events"
	"github.com/weitend/calendar-go/storage"
)

type Calendar struct {
	calendarEvents map[string]*events.Event
	Notification   chan string
	storage        storage.Store
}

func NewCalendar(s storage.Store) *Calendar {
	return &Calendar{
		calendarEvents: make(map[string]*events.Event),
		Notification:   make(chan string),
		storage:        s,
	}
}

func (c *Calendar) AddEvent(title string, dateStr string, priority events.Priority) (*events.Event, error) {
	e, err := events.NewEvent(title, dateStr, priority)

	if err != nil {
		return nil, err
	}

	c.calendarEvents[e.ID] = e

	return e, nil
}

func (c *Calendar) GetEvents() map[string]*events.Event {
	return c.calendarEvents
}

func (c *Calendar) DeleteEvent(id string) error {
	if c.calendarEvents[id] == nil {
		return errors.New("такого события не существует")
	}

	delete(c.calendarEvents, id)

	return nil
}

func (c *Calendar) EditEvent(id string, title string, dateStr string, priority events.Priority) error {

	e, exist := c.calendarEvents[id]

	if !exist {
		return errors.New("event with key " + id + " not found")
	}

	err := e.Update(title, dateStr, priority)
	if err != nil {
		return err
	}

	return nil
}
