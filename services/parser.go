package services

import "github.com/PuerkitoBio/goquery"

func ParseSite(url string) {
	doc, err := goquery.NewDocument(url)

	if err != nil {
		panic(err)
	}

	scrapersList := SelectScrapers(doc)

	dataMap := make(map[string]string)

	for _, scraper := range scrapersList {
		scraper.Perform(doc, dataMap)
	}

	PrintResults(dataMap)
}
