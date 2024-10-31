package logger

import "log/slog"

func Setup() error {
	slog.SetDefault(slog.Default())
	return nil
}
