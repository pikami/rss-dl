package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	. "./fileio"
	. "./helpers"
	. "./structs"
	"github.com/mmcdole/gofeed"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: rss-dl [FEED_URL]")
		return
	}

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(args[0])

	outputDir := ToCleanString(feed.Title)
	InitOutputDirectory(outputDir)

	WriteToFile(outputDir+"/feed_details.json", GrabFeedDetailsJSON(feed))
	for _, item := range feed.Items {
		itemOutputFilename := ToCleanString(
			item.PublishedParsed.Format("20060102") + "_" + item.Title)
		itemOutputDir := outputDir + "/" + itemOutputFilename

		if CheckIfExists(itemOutputDir) {
			continue
		}

		InitOutputDirectory(itemOutputDir)
		WriteToFile(
			itemOutputDir+"/details.json",
			GrabFeedItemJSON(item))
		DownloadFile(
			itemOutputDir+"/image"+filepath.Ext(item.Image.URL),
			item.Image.URL)
		for _, enclosure := range item.Enclosures {
			DownloadFile(
				itemOutputDir+"/"+filepath.Base(enclosure.URL),
				enclosure.URL)
		}
	}
}

// GrabFeedDetailsJSON - Returns a feed summary json sring
func GrabFeedDetailsJSON(feed *gofeed.Feed) string {
	feedParsed := &FeedDetails{
		Title:       feed.Title,
		Description: feed.Description,
		Categories:  feed.Categories,
		Language:    feed.Language,
		Link:        feed.Link,
		Updated:     feed.Updated,
	}

	feedDetails, _ := json.Marshal(feedParsed)
	return string(feedDetails)
}

// GrabFeedItemJSON - Returns a feed summary json sring
func GrabFeedItemJSON(item *gofeed.Item) string {
	itemParsed := &ItemDetails{
		Title:       item.Title,
		Description: item.Description,
		Content:     item.Content,
		Link:        item.Link,
		Updated:     item.Updated,
		Published:   item.Published,
		GUID:        item.GUID,
		Categories:  item.Categories,
	}

	itemDetails, _ := json.Marshal(itemParsed)
	return string(itemDetails)
}
