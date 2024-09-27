package log

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

type logEntry struct {
	App       string `json:"app"`
	RequestID string `json:"request_id"`
	Level     string `json:"level"`
	Time      string `json:"time"`
	Message   string `json:"message"`
}

func InitLogger() {
	var path = "/app/logs"
	var writer *rotatelogs.RotateLogs
	var err error

	log.SetFlags(0)

	if strings.EqualFold(os.Getenv("ENV"), "dev") {
		log.SetOutput(os.Stdout)
		return
	}

	writer, err = rotatelogs.New(
		fmt.Sprintf("%s/app-%s.log", path, "%Y-%m-%d"),
		rotatelogs.WithLinkName(fmt.Sprintf("%s/link.log", path)),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		log.Fatalf("failed to initialize log file: %s", err)
	}

	log.SetOutput(writer)
}

func write(id, level, msg string) {
	var now = time.Now().Format("2006-01-02T15:04:05.000-0700")
	var entry, _ = json.Marshal(logEntry{
		App:       "monitoring-api",
		RequestID: id,
		Level:     level,
		Time:      now,
		Message:   msg,
	})
	log.Println(string(entry))
}

func Debugf(id, msg string, data ...any) {
	write(id, "DEBUG", fmt.Sprintf(msg, data...))
}

func Infof(id, msg string, data ...any) {
	write(id, "INFO", fmt.Sprintf(msg, data...))
}

func Warnf(id, msg string, data ...any) {
	write(id, "WARN", fmt.Sprintf(msg, data...))
}

func Errorf(id, msg string, data ...any) {
	write(id, "ERROR", fmt.Sprintf(msg, data...))
	os.Exit(3)
}
