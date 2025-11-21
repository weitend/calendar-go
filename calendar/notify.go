package calendar

func (c *Calendar) Notify(msg string) {
	c.Notification <- msg
}
