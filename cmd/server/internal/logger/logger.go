package logger

import (
	"os"

	kitlog "github.com/go-kit/log"
)

type Logger interface {
	Log(keyvals ...any) error
}

func NewLogger() kitlog.Logger {
	return kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stdout))
}
