package events

import "errors"

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

func (p Priority) Validate() error {
	switch p {
	case PriorityLow, PriorityHigh, PriorityMedium:
		return nil
	default:
		return errors.New("invalid priority")
	}
}
