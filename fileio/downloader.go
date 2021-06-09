package fileio

import (
	"io"
	"net/http"
	"os"
)

// DownloadFile - Download file and store it
func DownloadFile(outputFilename string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(outputFilename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return err
	}

	return nil
}
