package main

import (
	"os"

	"github.com/g0j/tgbotapigen/tgparser"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger().
		Level(zerolog.InfoLevel).
		Output(zerolog.NewConsoleWriter())

	logger.Info().Msg("starting Telegram BotAPI codegen")
	logger.Debug().Int("len", len(os.Args)).Msg("Parsing arguments")

	config, err := ParseFlags(os.Args[1:])
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to parse command line flags")
	}

	logger = logger.Level(config.LogLevel)

	logger.Debug().Msg("checking config")

	err = config.Check()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to check config")
	}

	tp := tgparser.New().SetLogger(logger.With().Str("module", "parser").Logger())
	if config.File != "" {
		tp.SetFile(config.File)
	} else {
		tp.SetURL(config.URL)
	}

	err = tp.Fetch()
	if err != nil {
		logger.Fatal().Err(err).Msg("parsing error")
	}

	if config.VersionMode {
		v, err := tp.APIVersion()
		if err != nil {
			logger.Fatal().Err(err).Msg("failed to detect api version")
		}

		logger.Info().Str("current", v).Msg("api version")

		return
	}
}
