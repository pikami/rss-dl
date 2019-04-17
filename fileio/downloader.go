package fileio

import (
	"io"
	"net/http"
	"os"
)

// DownloadFile - Download file and store it
func DownloadFile(outputFilename string, url string) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(outputFilename)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	if err != nil {
		panic(err)
	}
}
