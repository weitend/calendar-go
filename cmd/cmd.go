package cmd

import (
	"github.com/weitend/calendar-go/calendar"
)

type Cmd struct {
	calendar *calendar.Calendar
	Logger   *Logger
}

func NewCmd(c *calendar.Calendar) *Cmd {
	return &Cmd{
		calendar: c,
		Logger: &Logger{
			messages: []string{},
		},
	}
}
