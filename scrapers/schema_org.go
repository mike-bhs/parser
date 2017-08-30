package scrapers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type SchemaOrg struct{}

const SchemaOrgStr = "http://schema.org"

// searh by itemscope attribute insted of checking for substring
func (sch SchemaOrg) HasNecessaryData(doc *goquery.Document) bool {
	resp, err := http.Get(doc.Url.String())

	if err != nil {
		return false
	}
	defer resp.Body.Close()

	htmlPage, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return false
	}

	htmlString := fmt.Sprintf("%s", htmlPage)

	return strings.Contains(htmlString, SchemaOrgStr)
}

func (sch SchemaOrg) Perform(doc *goquery.Document, results map[string]string) {
	doc.Find("meta").Each(func(index int, el *goquery.Selection) {
		if hasSchemaData(el) {
			propName, _ := el.Attr("itemprop")
			schemaData, _ := el.Attr("content")

			schemaName := fmt.Sprintf("schema_%s", propName)
			results[schemaName] = schemaData
		}
	})
}

func hasSchemaData(element *goquery.Selection) bool {
	_, exists := element.Attr("itemprop")

	return exists
}
