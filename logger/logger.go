package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	Info   *log.Logger
	Error  *log.Logger
	System *log.Logger
}

var logger *Logger
var file *os.File
var filename = "app.log"

func Init() {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	file = f

	if err != nil {
		fmt.Println(err)
	}

	logger = &Logger{
		Info:   log.New(file, "INFO:", log.Ldate|log.Ltime|log.Lshortfile),
		Error:  log.New(file, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile),
		System: log.New(file, "SYSTEM:", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func Info(msg string) {
	logger.Info.Output(2, msg)
}

func Error(msg string) {
	logger.Error.Output(2, msg)
}

func System(msg string) {
	logger.System.Output(2, msg)
}

func Finish() {
	file.Close()
}
