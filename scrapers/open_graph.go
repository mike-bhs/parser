package scrapers

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const OgPrefix = "og:"

// pass to contructor only that data
// which should be saved inside object (dependencies etc)
// remove string dependecies

type OpenGraph struct{}

func (og OpenGraph) HasNecessaryData(doc *goquery.Document) bool {
	result := false

	// replace Each by EachWithBreak
	doc.Find("meta").Each(func(i int, el *goquery.Selection) {
		result = hasOgData(el) || result
	})

	return result
}

func hasOgData(element *goquery.Selection) bool {
	propValue, _ := element.Attr("property")

	if strings.HasPrefix(propValue, OgPrefix) {
		return true
	}

	return false
}

func (og OpenGraph) Perform(doc *goquery.Document, results map[string]string) {
	doc.Find("meta").Each(func(i int, el *goquery.Selection) {
		if hasOgData(el) {
			ogName, _ := el.Attr("property")
			ogData, _ := el.Attr("content")

			results[FormatMetaName(ogName)] = ogData
		}
	})
}
