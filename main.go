package main

import (
	"os"

	"github.com/PuerkitoBio/goquery"
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

	var doc *goquery.Document

	if config.File != "" {
		logger.Info().Str("file", config.File).Msg("fetching file")
		doc, err = APIDocumentFromFile(config.File)
	} else {
		logger.Info().Str("url", config.URL).Msg("fetching url")
		doc, err = APIDocumentFromURL(config.URL)
	}

	if err != nil {
		logger.Fatal().Err(err).Msg("failed to fetch api")
	}

	if config.VersionMode {
		v, err := ParseAPIVersion(doc)
		if err != nil {
			logger.Fatal().Err(err).Msg("failed to detect api version")
		}

		logger.Info().Str("ver", v).Msg("api version")

		return
	}
}
