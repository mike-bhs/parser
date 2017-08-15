package scrapers

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

const OgRegexp = "og:.*"

type OpenGraph struct {
	Url string
}

func (og OpenGraph) HasNecessaryData() bool {
	doc, err := goquery.NewDocument(og.Url)
	result := false

	if err != nil {
		return false
	}

	doc.Find("meta").Each(func(i int, el *goquery.Selection) {
		result = hasOgData(el) || result
	})

	return result
}

func hasOgData(element *goquery.Selection) bool {
	propValue, _ := element.Attr("property")

	match, _ := regexp.MatchString(OgRegexp, propValue)

	if match {
		return true
	}

	return false
}

func (og OpenGraph) Perform(results map[string]string) {
	doc, err := goquery.NewDocument(og.Url)

	if err != nil {
		return
	}

	doc.Find("meta").Each(func(i int, el *goquery.Selection) {
		if hasOgData(el) {
			ogName, _ := el.Attr("property")
			ogData, _ := el.Attr("content")

			results[FormatMetaName(ogName)] = ogData
		}
	})
}
