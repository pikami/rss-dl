package structs

import (
	"flag"
	"fmt"
	"os"
)

// Config - Runtime configuration
var Config struct {
	FeedURL    string
	OutputPath string
	ParseHtml  bool
}

// GetConfig - Returns Config object
func GetConfig() {
	outputPath := flag.String("output", ".", "Output path")
	parseHtml := flag.Bool("parsehtml", false, "Save content as html")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: rss-dl [OPTIONS] FEED_URL")
		os.Exit(2)
	}

	Config.FeedURL = flag.Args()[len(args)-1]
	Config.OutputPath = *outputPath
	Config.ParseHtml = *parseHtml
}
