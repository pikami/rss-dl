# RSS-DL

[![Go Report Card](https://goreportcard.com/badge/github.com/pikami/rss-dl)](https://goreportcard.com/report/github.com/pikami/rss-dl)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/pikami/rss-dl/rss-dl_CI)

A simple rss feed downloader written in go

## Basic usage
Clone this repository and run `go build` to build the executable.\
You can download feeds by running `./rss-dl [Options] FEED_URL`

## Available options
* `-output some_directory` - Output path (default ".")

## Acknowledgments
This software uses the gofeed parser which can be found here: https://github.com/mmcdole/gofeed
