package main

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	filename string
}

func NewLogger(filename string) *Logger {
	return &Logger{filename: filename}
}

func (l *Logger) write(level, msg string) error {
	f, err := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	line := fmt.Sprintf("[%s] %-5s %s\n",
		time.Now().Format("2006-01-02 15:04:05"), level, msg)
	_, err = f.WriteString(line)
	return err
}

func (l *Logger) Info(msg string)  { l.write("INFO", msg) }
func (l *Logger) Warn(msg string)  { l.write("WARN", msg) }
func (l *Logger) Error(msg string) { l.write("ERROR", msg) }

func main() {
	log := NewLogger("some app.log")

	log.Info("One, two three, GO!")
	log.Info("Attention! Breaking News!")
	log.Warn("Something went wrong...")
	log.Warn("HEY! LISTEN!")
	log.Error("Cannot reach some file I...actually don't really know about")
	log.Error("Cannot reach some file oh MY GOD LEAVE ME ALONE")
	log.Info("Almost finished!...")
	log.Info("Finished")

	content, _ := os.ReadFile("app.log")
	fmt.Print(string(content))
}
