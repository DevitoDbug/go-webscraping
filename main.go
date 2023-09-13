package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)

	c.OnRequest(func(request *colly.Request) {
		//request.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36\n")
		fmt.Println("We are visiting a URL", request.URL)
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Response Code", response.StatusCode)
	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("error", err.Error())
	})

	c.OnHTML(".text", func(element *colly.HTMLElement) {
		fmt.Println("Quote", element.Text)
	})
	c.OnHTML(".author", func(element *colly.HTMLElement) {
		fmt.Println("Author", element.Text)
	})

	err := c.Visit("http://quotes.toscrape.com/random")
	if err != nil {
		return
	}
}
