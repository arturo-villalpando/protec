package config

import (
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
)

func Logger() *zerolog.Logger {
	buildInfo, _ := debug.ReadBuildInfo()

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Int("pid", os.Getpid()).
		Str("go_version", buildInfo.GoVersion).
		Logger().
		Sample(&zerolog.BasicSampler{N: 10})

	return &logger
}
