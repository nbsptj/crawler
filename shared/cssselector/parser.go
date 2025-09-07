package cssselector

import (
	"bytes"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// ParseToDoc parse bytes to goquery Document
func ParseToDoc(data []byte) (*goquery.Document, bool) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		fmt.Printf("New Document From Reader error: %v\n", err)
		return nil, false
	}
	return doc, true
}
