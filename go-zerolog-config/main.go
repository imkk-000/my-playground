package main

import (
	"os"

	"github.com/rs/zerolog"
)

func cally(logger zerolog.Logger) {
	mylogger := logger.With().Str("service", "cally").Logger()
	mylogger.Info().Msg("call by cally")
	logger.Info().Msg("call by cally")
}

func main() {
	logger := zerolog.New(os.Stdout).
		With().Timestamp().Logger()
	zerolog.TimestampFieldName = "custom_time"
	logger.Info().Msg("Hi")

	cally(logger)

	logger.Info().Msg("World")
}
