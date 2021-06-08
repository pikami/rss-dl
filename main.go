package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/mmcdole/gofeed"

	fileio "github.com/pikami/rss-dl/fileio"
	helpers "github.com/pikami/rss-dl/helpers"
	structs "github.com/pikami/rss-dl/structs"
)

func main() {
	structs.GetConfig()

	fp := gofeed.NewParser()
	helpers.LogInfo("Downloading " + structs.Config.FeedURL)
	feed, _ := fp.ParseURL(structs.Config.FeedURL)

	outputDir := structs.Config.OutputPath + "/" + helpers.ToCleanString(feed.Title)
	fileio.InitOutputDirectory(outputDir)

	feedInfoPath := outputDir + "/feed_details.json"
	helpers.LogInfo("Writing feed details as JSON to " + feedInfoPath)
	fileio.WriteToFile(feedInfoPath, GrabFeedDetailsJSON(feed))

	feedImagePath := outputDir + "/image" + helpers.RemoveGetParams(filepath.Ext(feed.Image.URL))
	fileio.DownloadFile(feedImagePath, feed.Image.URL)

	for _, item := range feed.Items {
		itemOutputFilename := helpers.ToCleanString(
			item.PublishedParsed.Format("20060102") + "_" + item.Title)
		itemOutputDir := outputDir + "/" + itemOutputFilename

		if fileio.CheckIfExists(itemOutputDir) {
			fmt.Println("Item '" + item.Title + "' already downloaded, skipping")
			continue
		}

		helpers.LogInfo("Downloading feed item '" + item.Title + "' to " + itemOutputDir)
		fileio.InitOutputDirectory(itemOutputDir)

		itemDetailsPath := itemOutputDir + "/details.json"
		helpers.LogInfo("Writing details to " + itemDetailsPath)
		fileio.WriteToFile(
			itemDetailsPath,
			GrabFeedItemJSON(item))

		if item.Image != nil {
			itemImagePath := itemOutputDir + "/image" + helpers.RemoveGetParams(filepath.Ext(item.Image.URL))
			helpers.LogInfo("Downloading image to " + itemImagePath)
			fileio.DownloadFile(
				itemImagePath,
				item.Image.URL)
		}

		for _, enclosure := range item.Enclosures {
			filename := helpers.RemoveGetParams(filepath.Base(enclosure.URL))
			helpers.LogInfo("Downloading attachment '" + filename + "'")
			fileio.DownloadFile(
				itemOutputDir+"/"+filename,
				enclosure.URL)
		}
	}
}

// GrabFeedDetailsJSON - Returns a feed summary json sring
func GrabFeedDetailsJSON(feed *gofeed.Feed) string {
	feedParsed := &structs.FeedDetails{
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
	itemParsed := &structs.ItemDetails{
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
