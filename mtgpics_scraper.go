package mtgpics_scraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

func Scrape() string {
	c := colly.NewCollector(
	)

	c.OnHTML("div[style*='background']", func(e *colly.HTMLElement) {
		// Extracting the nested information
		cardName := e.ChildText(".Card12 a.und")
		setName := e.ChildText(".Card12 a[href^='art?set=']")
		artistInfo := e.ChildText(".S10 a[href^='art?set=']")
		imageDimensions := e.ChildText(".S10")

		// Printing the information
		fmt.Println("Card Name:", cardName)
		fmt.Println("Set Name:", setName)
		fmt.Println("Artist Info:", artistInfo)
		fmt.Println("Image Dimensions:", imageDimensions)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on mtgpics at the set selected
	c.Visit("https://www.mtgpics.com/art?set=421")

	return "Get Scraped!"
}