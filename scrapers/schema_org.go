package scrapers

import (
	"github.com/PuerkitoBio/goquery"
)

type SchemaOrg struct{}

func (sch SchemaOrg) HasNecessaryData(doc *goquery.Document) bool {
	return doc.Find("[itemscope]").Length() > 0
}

func (sch SchemaOrg) Perform(doc *goquery.Document) map[string]string {
	results := make(map[string]string)

	doc.Find("meta").Each(func(index int, el *goquery.Selection) {
		if hasSchemaData(el) {
			propName, _ := el.Attr("itemprop")
			schemaData, _ := el.Attr("content")

			results[propName] = schemaData
		}
	})

	return results
}

func hasSchemaData(element *goquery.Selection) bool {
	_, exists := element.Attr("itemprop")

	return exists
}
