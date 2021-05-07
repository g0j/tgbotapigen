package tgparser

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

var ErrNoResponse = errors.New("no response received")

// fetchDocument parses html content of reader and convert in to goquery.Document.
func (p *Parser) fetchDocument(r io.Reader) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("failed to create reader: %w", err)
	}

	return doc, nil
}

// fetchURL is fetchDocument wrapper for url parsing.
func (p *Parser) fetchURL() (*goquery.Document, error) {
	resp, err := http.Get(p.uri)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL %s: %w", p.uri, err)
	}

	if resp == nil {
		return nil, ErrNoResponse
	}

	defer resp.Body.Close()

	return p.fetchDocument(resp.Body)
}

// fetchFile is fetchDocument wrapper for file parsing.
func (p *Parser) fetchFile() (*goquery.Document, error) {
	f, err := os.OpenFile(p.uri, os.O_RDONLY, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch file %s: %w", p.uri, err)
	}

	defer f.Close()

	return p.fetchDocument(f)
}
