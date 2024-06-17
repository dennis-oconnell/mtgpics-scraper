package mtgpics_scraper

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly"
)

func ScrapeArtPics() string {
	// Initialize a new Colly collector
	c := colly.NewCollector()

	// Initialize a data structure to hold scraped data
	type cardImage struct {
		imageURL, cardName, artistName, setName string
	}

	// On every HTML element which contains a style attribute with a background
	c.OnHTML("div[style*='background']", func(e *colly.HTMLElement) {

		//Extract the style attribute of the html element
		style := e.Attr("style")

		// Use regex to find the background URL
        re := regexp.MustCompile(`url\((.*?)\)`)
        match := re.FindStringSubmatch(style)
		
		// Disregard the dud images by ignoring the elements with a background image url of length less than 1
		if len(match) > 1 {
			currentImage := cardImage{}

			currentImage.cardName = e.ChildText(".Card12 a.und")
			currentImage.imageURL = match[1]
			currentImage.artistName = e.ChildText(".S10 a[href^='art?set=']")
			currentImage.setName = e.ChildText(".Card12 a[href^='art?set=']")

			fmt.Println(currentImage)
        } 
		
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on mtgpics at the set selected
	c.Visit("https://www.mtgpics.com/art?set=421")
	c.Visit("https://www.mtgpics.com/art?set=421&pointeur=60")
	c.Visit("https://www.mtgpics.com/art?set=421&pointeur=120")
	c.Visit("https://www.mtgpics.com/art?set=421&pointeur=180")
	c.Visit("https://www.mtgpics.com/art?set=421&pointeur=240")

	return "All done scraping!"
}