package scrapers

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const OgPrefix = "og:"

type OpenGraph struct{}

func (og OpenGraph) HasNecessaryData(doc *goquery.Document) bool {
	return doc.Find("meta").EachWithBreak(func(i int, el *goquery.Selection) bool {
		return hasOgData(el)
	}).Length() > 0
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
