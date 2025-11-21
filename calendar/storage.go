package calendar

import "encoding/json"

func (c *Calendar) Save() error {
	data, marshErr := json.Marshal(c.calendarEvents)

	if marshErr != nil {
		return marshErr
	}

	storErr := c.storage.Save(data)

	return storErr
}

func (c *Calendar) Load() error {
	data, storErr := c.storage.Load()

	if storErr != nil {
		return storErr
	}

	unmarshErr := json.Unmarshal(data, &c.calendarEvents)

	return unmarshErr
}
