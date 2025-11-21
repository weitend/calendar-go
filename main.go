package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/weitend/calendar-go/calendar"
	"github.com/weitend/calendar-go/cmd"
	"github.com/weitend/calendar-go/logger"
	"github.com/weitend/calendar-go/reminder"
	"github.com/weitend/calendar-go/storage"
)

func main() {
	logger.Init()
	defer logger.Finish()

	jsonStorage := storage.NewJsonStorage("calendar.json")
	// zipStorage := storage.NewZipStorage("calendar.zip")

	c := calendar.NewCalendar(jsonStorage)

	loadErr := c.Load()
	if loadErr != nil {
		fmt.Println("Ошибка загрузки данных:", loadErr)
		return
	}
	defer func() {
		saveErr := c.Save()
		if saveErr != nil {
			fmt.Println("Ошибка загрузки данных:", saveErr)
			return
		}
	}()

	err := c.AddEventReminder("9fedf744-34e6-41e4-8f57-df3d599084ab", "", time.Now(), 5*time.Second)

	if errors.Is(err, reminder.ErrEmptyMessage) {
		fmt.Println("can't set reminder with empty message")
	} else if err != nil {
		fmt.Println(err)
	}

	cli := cmd.NewCmd(c)
	cli.Run()
}
