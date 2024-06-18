package main

import (
	"mtgpics_scraper"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	//fmt.Println(mtgpics_scraper.ScrapeArtPics())

	spew.Dump(mtgpics_scraper.ScrapeArtPics())
}