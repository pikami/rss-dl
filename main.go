package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	. "./fileio"
	. "./helpers"
	. "./structs"
	"github.com/mmcdole/gofeed"
)

func main() {
	GetConfig()

	fp := gofeed.NewParser()
	LogInfo("Downloading " + Config.FeedURL)
	feed, _ := fp.ParseURL(Config.FeedURL)

	outputDir := Config.OutputPath + "/" + ToCleanString(feed.Title)
	InitOutputDirectory(outputDir)

	feedInfoPath := outputDir + "/feed_details.json"
	LogInfo("Writing feed details as JSON to " + feedInfoPath)
	WriteToFile(feedInfoPath, GrabFeedDetailsJSON(feed))

	feedImagePath := outputDir + "/image" + RemoveGetParams(filepath.Ext(feed.Image.URL))
	DownloadFile(feedImagePath, feed.Image.URL)

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

		if item.Image != nil {
			itemImagePath := itemOutputDir + "/image" + RemoveGetParams(filepath.Ext(item.Image.URL))
			LogInfo("Downloading image to " + itemImagePath)
			DownloadFile(
				itemImagePath,
				item.Image.URL)
		}

		for _, enclosure := range item.Enclosures {
			filename := RemoveGetParams(filepath.Base(enclosure.URL))
			LogInfo("Downloading attachment '" + filename + "'")
			DownloadFile(
				itemOutputDir+"/"+filename,
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
		FeedLink:    feed.FeedLink,
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
