package logging

import (
	"log/slog"
	"os"
)

func Logger(file *os.File, filename string, str string) (err error) {
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		logger.Warn("Something wrong with line 5\n cannot proceed further   !")
	}
	slog.Info("Logging to file has been done")
	return nil
}
