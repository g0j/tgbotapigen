package tgparser

import (
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog"
)

type Parser struct {
	doc    *goquery.Document
	mode   ParsingMode
	uri    string
	logger zerolog.Logger
}

func New() *Parser {
	return &Parser{
		logger: zerolog.New(os.Stdout),
	}
}

func (p *Parser) SetLogger(logger zerolog.Logger) *Parser {
	p.logger = logger

	return p
}

func (p *Parser) SetFile(filename string) *Parser {
	p.uri = filename
	p.mode = ParsingModeFile

	p.logger.Debug().Str("uri", p.uri).Msg("parser file changed")

	return p
}

func (p *Parser) SetURL(url string) *Parser {
	p.uri = url
	p.mode = ParsingModeURL

	p.logger.Debug().Str("uri", p.uri).Msg("parser url changed")

	return p
}
