package cmd

import (
	"fmt"
	"sync"
	"time"
)

type LogType string

const (
	infoLog   LogType = "INFO"
	inputLog  LogType = "USER_INPUT"
	errLog    LogType = "ERROR"
	remindLog LogType = "REMIND"
)

type Logger struct {
	messages []string
}

func (l *Logger) Log() {
	for _, msg := range l.messages {
		fmt.Println(msg)
	}
}

func (l *Logger) Write(writeType LogType, t time.Time, msg string) {
	var m sync.Mutex
	m.Lock()
	defer m.Unlock()

	formattedTime := t.Format("2006-01-02 15:04:05")
	message := fmt.Sprintf("[%s] %s: %s", formattedTime, writeType, msg)

	l.messages = append(l.messages, message)
}
