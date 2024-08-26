package main

import (
	"fmt"
)

func getAndPrintPage(cfg *config) {

	fmt.Printf("Page number: %d\n", cfg.page+1)

	url := createURL(cfg)
	fmt.Printf("URL: %s\n", url)

	pageInfo := getPageInfo(url)

	for _, page := range pageInfo.Results {
		fmt.Println(page.Name)
	}

}
