package main

import (
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

var ErrNoResponse = errors.New("no response received")

// APIDocument parses html content of reader and convert in to goquery.Document.
func APIDocument(r io.Reader) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// APIDocumentFromURL is APIDocument wrapper for url parsing.
func APIDocumentFromURL(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, ErrNoResponse
	}

	defer resp.Body.Close()

	return APIDocument(resp.Body)
}

// APIDocumentFromFile is APIDocument wrapper for file parsing.
func APIDocumentFromFile(filename string) (*goquery.Document, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	return APIDocument(f)
}
