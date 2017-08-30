package scrapers

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const TwitterPrefix = "twitter:"

type TwitterCard struct{}

func (tw TwitterCard) HasNecessaryData(doc *goquery.Document) bool {
	return doc.Find("meta").EachWithBreak(func(i int, el *goquery.Selection) bool {
		return hasTwitterData(el)
	}).Length() > 0
}

func (tw TwitterCard) Perform(doc *goquery.Document) map[string]string {
	results := make(map[string]string)

	doc.Find("meta").Each(func(i int, el *goquery.Selection) {
		if hasTwitterData(el) {
			twName, _ := el.Attr("name")
			twData, _ := el.Attr("content")
			twKey := FormatMetaName(RemoveSubstring(twName, TwitterPrefix))

			results[twKey] = twData
		}
	})

	return results
}

func hasTwitterData(element *goquery.Selection) bool {
	propValue, _ := element.Attr("name")

	if strings.HasPrefix(propValue, TwitterPrefix) {
		return true
	}

	return false
}
