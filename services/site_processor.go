package services

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	scrapers "github.com/parser/scrapers"
)

func SelectScrapers(doc *goquery.Document) []scrapers.Scraper {
	allScrapers := []scrapers.Scraper{
		scrapers.OpenGraph{},
		scrapers.TwitterCard{},
		scrapers.SchemaOrg{},
	}

	var selectedScrapers []scrapers.Scraper

	for _, scraper := range allScrapers {
		if scraper.HasNecessaryData(doc) {
			selectedScrapers = append(selectedScrapers, scraper)
		}
	}

	return selectedScrapers
}

func PrintResults(results map[string]string) {
	fmt.Printf("\nRESULTS\n")

	for k, v := range results {
		fmt.Printf("%s : %s\n", k, v)
	}
}
