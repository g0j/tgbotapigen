package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/rs/zerolog"
)

var (
	ErrUndefinedAPIURI = errors.New("either api url or api file must be specified")
	ErrTooManyURI      = errors.New("only one api uri must be specified")
)

// Telegram BotAPI default url.
const botAPIURL = "https://core.telegram.org/bots/api"

// Config stores values received from command line arguments.
type Config struct {
	File        string
	URL         string
	LogLevel    zerolog.Level
	VersionMode bool
}

// Check vlidates required fields of config.
func (c *Config) Check() error {
	switch {
	case c.File == "" && c.URL == "":
		return ErrUndefinedAPIURI
	case c.File != "" && c.URL != "":
		return ErrTooManyURI
	}

	return nil
}

var defaultConfig = Config{
	File:        "",
	URL:         botAPIURL,
	LogLevel:    zerolog.InfoLevel,
	VersionMode: true,
}

// ParseFlags parses command line arguments and returns it in Config struct
// Prints usage if it can't parse arguments.
func ParseFlags(args []string) (*Config, error) {
	conf := defaultConfig

	flagset := flag.NewFlagSet("tgbotapigen", flag.CommandLine.ErrorHandling())
	flagset.StringVar(
		&conf.File,
		"f",
		defaultConfig.File,
		"use pre-downloaded html file instead of getting it from Telegram",
	)
	flagset.StringVar(
		&conf.URL,
		"u",
		defaultConfig.URL,
		"use this url instead of default Telegram BotAPI url",
	)
	flagset.BoolVar(
		&conf.VersionMode,
		"v",
		defaultConfig.VersionMode,
		"parse api version and exit",
	)
	flagset.Func(
		"ll",
		"defines logging level",
		func(levelStr string) (err error) {
			conf.LogLevel, err = zerolog.ParseLevel(levelStr)

			return
		},
	)

	if err := flagset.Parse(args); err != nil {
		return nil, fmt.Errorf("failed to parse flags: %w", err)
	}

	return &conf, nil
}
