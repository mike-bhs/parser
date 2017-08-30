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
		ConcatMaps(dataMap, scraper.Perform(doc))
	}

	PrintResults(dataMap)
}

func ConcatMaps(mainMap, childMap map[string]string) {
	for k, v := range childMap {
		_, hasKey := mainMap[k]

		if !hasKey {
			mainMap[k] = v
		}
	}
}
