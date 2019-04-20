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
	LogInfo("Downloading " + args[0])
	feed, _ := fp.ParseURL(args[0])

	outputDir := ToCleanString(feed.Title)
	InitOutputDirectory(outputDir)

	feedInfoPath := outputDir + "/feed_details.json"
	LogInfo("Writing feed details as JSON to " + feedInfoPath)
	WriteToFile(feedInfoPath, GrabFeedDetailsJSON(feed))

	for _, item := range feed.Items {
		itemOutputFilename := ToCleanString(
			item.PublishedParsed.Format("20060102") + "_" + item.Title)
		itemOutputDir := outputDir + "/" + itemOutputFilename

		if CheckIfExists(itemOutputDir) {
			fmt.Println("Item '" + item.Title + "' already downloaded, skipping")
			continue
		}

		LogInfo("Downloading feed item '" + item.Title + "' to " + itemOutputDir)
		InitOutputDirectory(itemOutputDir)

		itemDetailsPath := itemOutputDir + "/details.json"
		LogInfo("Writing details to " + itemDetailsPath)
		WriteToFile(
			itemDetailsPath,
			GrabFeedItemJSON(item))

		itemImagePath := itemOutputDir + "/image" + filepath.Ext(item.Image.URL)
		LogInfo("Downloading image to " + itemImagePath)
		DownloadFile(
			itemImagePath,
			item.Image.URL)

		for _, enclosure := range item.Enclosures {
			LogInfo("Downloading attachment '" + filepath.Base(enclosure.URL) + "'")
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
