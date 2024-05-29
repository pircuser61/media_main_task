package config

import (
	"log/slog"
	"os"
)

func GetAddr() string {
	return getEnv("EXCHANGE_ADDR", "127.0.0.1:8088")
}

func GetLogLevel() slog.Level {
	ll := slog.LevelError
	switch getEnv("LOG_LEVEL", "DEBUG") {
	case "DEBUG":
		ll = slog.LevelDebug
	case "WARN":
		ll = slog.LevelWarn
	case "INFO":
		ll = slog.LevelInfo
	}
	return ll
}

func getEnv(name string, defaultValue string) string {
	val, ok := os.LookupEnv(name)
	if ok {
		return val
	}
	return defaultValue
}
