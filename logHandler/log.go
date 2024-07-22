package loghandler

import (
	"log/slog"
	"os"
)

var LogFile *os.File
var err error
var AppLogger *slog.Logger

func OpenLogFile() {
	LogFile, err = os.OpenFile("logfile.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		// Handle error opening the file
		panic(err)
	}

}

func Log() {
	jsonHandler := slog.NewJSONHandler(LogFile, &slog.HandlerOptions{})

	AppLogger = slog.New(jsonHandler)

	AppLogger.Info("Log Started")
}
