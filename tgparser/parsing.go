package tgparser

import (
	"errors"
)

const SelectorAPIVersion = "#dev_page_content>p>strong"

type ParsingMode int8

const (
	ParsingModeFile ParsingMode = iota + 1
	ParsingModeURL
)

var (
	ErrSelectorNotFound      = errors.New("selector not found")
	ErrUnexpectedParsingMode = errors.New("unexpected parsing mode")
)

// Fetch fetches current URI and create parseable document from it.
func (p *Parser) Fetch() (err error) {
	p.logger.Debug().Str("uri", p.uri).Msg("parsing document")

	switch p.mode {
	case ParsingModeFile:
		p.doc, err = p.fetchFile()
	case ParsingModeURL:
		p.doc, err = p.fetchURL()
	default:
		return ErrUnexpectedParsingMode
	}

	return
}

// APIVersion parses api version from first api change notice.
func (p *Parser) APIVersion() (string, error) {
	v := p.doc.Find(SelectorAPIVersion).First().Text()
	if v == "" {
		return v, ErrSelectorNotFound
	}

	return v, nil
}
