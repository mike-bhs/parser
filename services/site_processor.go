package services

import (
	"fmt"

	scrapers "github.com/parser/scrapers"
)

func SelectScrapers(url string) []scrapers.Scraper {
	allScrapers := []scrapers.Scraper{
		scrapers.OpenGraph{Url: url},
		scrapers.TwitterCard{Url: url},
		scrapers.SchemaOrg{Url: url},
	}

	var selectedScrapers []scrapers.Scraper

	for _, scraper := range allScrapers {
		if scraper.HasNecessaryData() {
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
