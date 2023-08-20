package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Info message") // {"time":"2023-08-21T01:37:49.760783+09:00","level":"INFO","msg":"Info message"}
}
