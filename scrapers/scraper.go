package scrapers

import "strings"

type Scraper interface {
	HasNecessaryData() bool
	Perform(map[string]string)
}

func FormatMetaName(name string) string {
	return strings.Replace(name, ":", "_", -1)
}
