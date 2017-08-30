package scrapers

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Scraper interface {
	HasNecessaryData(*goquery.Document) bool
	Perform(*goquery.Document) map[string]string
}

func FormatMetaName(name string) string {
	return strings.Replace(name, ":", "_", -1)
}

func RemoveSubstring(str, substr string) string {
	return strings.Replace(str, substr, "", -1)
}

// повертати з Perform готовий об'єкт
// лишити логіку парсингу всередині кожного скрапера
// і після цього з'єднати їх докупи
// зробити парсинг всередині скоупу для schema org
