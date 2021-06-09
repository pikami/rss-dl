package htmlparse

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"

	fileio "github.com/pikami/rss-dl/fileio"
	helpers "github.com/pikami/rss-dl/helpers"
)

func HtmlGrab(htmlStr string, itemOutputDir string) {
	rootNode, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return
	}

	// Init download dir
	outputDir := itemOutputDir + "/html"
	fileio.InitOutputDirectory(outputDir)

	// Load the HTML document
	doc := goquery.NewDocumentFromNode(rootNode)

	// Download assets
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		val, exists := s.Attr("src")
		if exists {
			imageName := "#"
			if strings.Contains(val, "base64") {
				imageName = fileio.SaveFromBase64(val, outputDir)
			} else {
				imageName = helpers.RemoveGetParams(filepath.Base(val))
				itemImagePath := outputDir + "/" + imageName
				helpers.LogInfo("Downloading image to " + itemImagePath)
				err = fileio.DownloadFile(
					itemImagePath,
					val)

				if err != nil {
					fmt.Printf("[htmlgrab] %d: failed to get %s\n", i, val)
				} else {
					fmt.Printf("[htmlgrab] %d: %s\n", i, val)
				}
			}

			s.SetAttr("src", imageName)
		}
	})

	newHtml, err := doc.Html()
	if err != nil {
		return
	}

	fileio.WriteToFile(outputDir+"/index.html", newHtml)
}
