package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/c-bata/go-prompt"
	"github.com/google/shlex"
	"github.com/weitend/event-calendar/events"
	"github.com/weitend/event-calendar/logger"
)

type Command string

const (
	addCmd    Command = "add"
	listCmd   Command = "list"
	removeCmd Command = "remove"
	helpCmd   Command = "help"
	exitCmd   Command = "exit"
	updateCmd Command = "update"
	logCmd    Command = "log"
)

var suggestions = []prompt.Suggest{
	{Text: string(addCmd), Description: "Добавить событие"},
	{Text: string(listCmd), Description: "Показать все события"},
	{Text: string(removeCmd), Description: "Удалить событие"},
	{Text: string(helpCmd), Description: "Показать справку"},
	{Text: string(exitCmd), Description: "Выйти из программы"},
	{Text: string(updateCmd), Description: "Изменить событие"},
	{Text: string(logCmd), Description: "Посмотреть историю ввода и вывода"},
}

func (c *Cmd) executor(input string) {
	parts, err := shlex.Split(input)

	if err != nil {
		msg := fmt.Sprintf("Ошибка при парсинге input: %s", err)

		fmt.Println(msg)
		c.Logger.Write(errLog, time.Now(), msg)
		logger.Error(msg)
	}

	if len(parts) < 1 {
		return
	}

	cmd := strings.ToLower(parts[0])

	switch Command(cmd) {
	case addCmd:
		if len(parts) < 4 {
			msg := "Формат: add \"название события\" \"дата и время\" \"приоритет\""

			fmt.Println(msg)
			c.Logger.Write(infoLog, time.Now(), msg)
			logger.Info(msg)

			return
		}

		title := parts[1]
		date := parts[2]
		priority := events.Priority(parts[3])

		e, err := c.calendar.AddEvent(title, date, priority)

		if err != nil {
			msg := fmt.Sprintf("Ошибка добавления: %s", err)

			fmt.Println(msg)
			c.Logger.Write(errLog, time.Now(), msg)
			logger.Error(msg)
		} else {
			msg := fmt.Sprintf("Событие: %s добавлено", e.Title)

			fmt.Println(msg)
			c.Logger.Write(infoLog, time.Now(), msg)
			logger.Info(msg)

			c.calendar.Save()
		}
	case removeCmd:
		if len(parts) < 2 {
			msg := "Формат: remove \"ID события\""

			fmt.Println(msg)
			c.Logger.Write(infoLog, time.Now(), msg)
			logger.Info(msg)

			return
		}

		id := parts[1]

		err := c.calendar.DeleteEvent(id)

		if err != nil {
			msg := fmt.Sprintf("Ошибка удаления: %s", err)

			fmt.Println(msg)
			c.Logger.Write(errLog, time.Now(), msg)
			logger.Error(msg)

			return
		} else {
			msg := fmt.Sprintf("Событие с ID %s успешно удалено", id)

			fmt.Println(msg)
			c.Logger.Write(infoLog, time.Now(), msg)
			logger.Info(msg)

			c.calendar.Save()
		}
	case updateCmd:
		if len(parts) < 5 {
			msg := "Формат: update \"ID события\" \"название события\" \"дата и время события\" \"приоритет события\""

			fmt.Println(msg)
			c.Logger.Write(infoLog, time.Now(), msg)
			logger.Info(msg)

			return
		}

		id := parts[1]
		title := parts[2]
		date := parts[3]
		priority := parts[4]

		err := c.calendar.EditEvent(id, title, date, events.Priority(priority))

		if err != nil {
			msg := fmt.Sprintf("Ошибка изменения: %s", err)

			fmt.Println(msg)
			c.Logger.Write(errLog, time.Now(), msg)
			logger.Error(msg)
		} else {
			msg := fmt.Sprintf("Событие с ID %s успешно изменено", id)

			fmt.Println(msg)
			c.Logger.Write(infoLog, time.Now(), msg)
			logger.Info(msg)

			c.calendar.Save()
		}
	case listCmd:
		events := c.calendar.GetEvents()

		if len(events) == 0 {
			msg := "Ни одного события еще не добавлено!"

			fmt.Println(msg)
			c.Logger.Write(infoLog, time.Now(), msg)
			logger.Info(msg)

			return
		}

		var b strings.Builder

		for _, event := range events {
			fmt.Fprintln(&b)
			fmt.Fprintln(&b, "----------")
			fmt.Fprintln(&b, "ID:", event.ID)
			fmt.Fprintln(&b, "Название:", event.Title)
			fmt.Fprintln(&b, "Когда:", event.StartAt.Format("2006-01-02 15:04:05"))
			fmt.Fprintln(&b, "Приоритет:", event.Priority)
			fmt.Fprintln(&b, "----------")
		}

		listOutput := b.String()

		fmt.Print(listOutput)
		c.Logger.Write(infoLog, time.Now(), listOutput)
		logger.Info(listOutput)
	case exitCmd:
		err := c.calendar.Save()

		if err != nil {

			msg := fmt.Sprintf("Ошибка сохранения календаря: %s", err)

			fmt.Println(msg)
			c.Logger.Write(errLog, time.Now(), msg)
			logger.Error(msg)
		}

		close(c.calendar.Notification)
		os.Exit(0)
	case helpCmd:

		var b strings.Builder

		fmt.Fprintln(&b, "Список всех возможных команд:")
		for _, cmd := range suggestions {
			fmt.Fprintln(&b, cmd.Text, "-", cmd.Description)
		}

		cmdsOutput := b.String()

		fmt.Print(cmdsOutput)
		c.Logger.Write(infoLog, time.Now(), cmdsOutput)
		logger.Info(cmdsOutput)
	case logCmd:
		c.Logger.Log()
	default:
		var b strings.Builder

		fmt.Fprintln(&b, "Неизвестная команда:", input)
		fmt.Fprintln(&b, "Введите 'help' для списка команд")

		infoOutput := b.String()

		fmt.Print(infoOutput)
		c.Logger.Write(infoLog, time.Now(), infoOutput)
		logger.Info(infoOutput)
	}
	c.Logger.Write(inputLog, time.Now(), input)
	logger.System("Введен текст: " + input)
}

func (c *Cmd) completer(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func (c *Cmd) Run() {
	p := prompt.New(c.executor, c.completer, prompt.OptionPrefix("> "))

	go func() {
		for msg := range c.calendar.Notification {
			msg := fmt.Sprintf("Напоминание сработало: %s", msg)

			fmt.Println(msg)
			c.Logger.Write(remindLog, time.Now(), msg)
			logger.System(msg)
		}
	}()

	p.Run()
}
