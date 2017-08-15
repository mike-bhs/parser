package scrapers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type SchemaOrg struct {
	Url string
}

const SchemaOrgStr = "http://schema.org"

func (sch SchemaOrg) HasNecessaryData() bool {
	resp, err := http.Get(sch.Url)

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

func (sch SchemaOrg) Perform(results map[string]string) {
	doc, err := goquery.NewDocument(sch.Url)

	if err != nil {
		return
	}

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
