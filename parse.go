package main

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
)

const SelectorAPIVersion = "#dev_page_content>p>strong"

var ErrSelectorNotFound = errors.New("selector not found")

// ParseAPIVersion parses api version from first api change notice.
func ParseAPIVersion(doc *goquery.Document) (string, error) {
	v := doc.Find(SelectorAPIVersion).First().Text()
	if v == "" {
		return v, ErrSelectorNotFound
	}

	return v, nil
}
