package services

// func ParseSite(url string) {
// 	doc, err := goquery.NewDocument(url)
//
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	parseSchemaObjects(doc)
// 	parseCustomObjects(doc)
// }

func ParseSite(url string) {
	scrapersList := SelectScrapers(url)

	dataMap := make(map[string]string)

	for _, scraper := range scrapersList {
		scraper.Perform(dataMap)
	}

	PrintResults(dataMap)
}

//
// func parseSchemaObjects(doc *goquery.Document) {
// 	schemaMap := make(map[string]string)
//
// 	doc.Find("[itemprop]").Each(func(index int, item *goquery.Selection) {
// 		propValue, _ := item.Attr("itemprop")
// 		iType, _ := item.Attr("itemtype")
// 		// fmt.Println(item.Text())
//
// 		if propValue != "" && item.Find("div").Length() == 0 {
// 			// if propValue != "" {
// 			schemaMap[propValue] = fmt.Sprintf("%s (%s)", strings.TrimSpace(item.Text()), iType)
// 		}
// 	})
//
// 	PrintSearchResults(schemaMap, "SCHEMA OBJECTS")
// }
//
// func parseCustomObjects(doc *goquery.Document) {
// 	customMap := make(map[string]string)
//
// 	doc.Find("td.attrLabels").Each(func(index int, item *goquery.Selection) {
// 		label := strings.TrimSpace(item.Text())
// 		nextContainer := item.Next()
//
// 		if nextContainer.Is("td") && nextContainer.Find("div").Length() == 0 {
// 			key := strings.Replace(label, ":", "", -1)
// 			customMap[key] = strings.TrimSpace(nextContainer.Text())
// 		}
// 	})
//
// 	PrintSearchResults(customMap, "CUSTOM OBJECTS")
// }

//
// func PrintSearchResults(results map[string]string, header string) {
// 	fmt.Printf("\n%s\n", header)
//
// 	for k, v := range results {
// 		fmt.Printf("%s : %s\n", k, v)
// 	}
// }
