package scrapers

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const TwitterPrefix = "twitter:"

type TwitterCard struct{}

func (tw TwitterCard) HasNecessaryData(doc *goquery.Document) bool {
	result := false

	doc.Find("meta").Each(func(i int, el *goquery.Selection) {
		result = hasTwitterData(el) || result
	})

	return result
}

func (tw TwitterCard) Perform(doc *goquery.Document, results map[string]string) {
	doc.Find("meta").Each(func(i int, el *goquery.Selection) {
		if hasTwitterData(el) {
			twName, _ := el.Attr("name")
			twData, _ := el.Attr("content")

			results[FormatMetaName(twName)] = twData
		}
	})
}

func hasTwitterData(element *goquery.Selection) bool {
	propValue, _ := element.Attr("name")

	if strings.HasPrefix(propValue, TwitterPrefix) {
		return true
	}

	return false
}
