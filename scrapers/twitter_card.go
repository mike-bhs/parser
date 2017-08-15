package scrapers

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

const TwitterCardRegexp = "twitter:.*"

type TwitterCard struct {
	Url string
}

func (tw TwitterCard) HasNecessaryData() bool {
	doc, err := goquery.NewDocument(tw.Url)
	result := false

	if err != nil {
		return false
	}

	doc.Find("meta").Each(func(i int, el *goquery.Selection) {
		result = hasTwitterData(el) || result
	})

	return result
}

func (tw TwitterCard) Perform(results map[string]string) {
	doc, err := goquery.NewDocument(tw.Url)

	if err != nil {
		return
	}

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

	match, _ := regexp.MatchString(TwitterCardRegexp, propValue)

	if match {
		return true
	}

	return false
}
