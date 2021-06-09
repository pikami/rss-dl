package fileio

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	crypto "github.com/pikami/rss-dl/crypto"
)

func SaveFromBase64(imgStr string, basePath string) string {
	sha1 := crypto.ShaStr(imgStr)
	coI := strings.Index(string(imgStr), ",")
	rawImage := string(imgStr)[coI+1:]

	// Encoded Image DataUrl //
	unbased, _ := base64.StdEncoding.DecodeString(string(rawImage))
	res := bytes.NewReader(unbased)

	switch strings.TrimSuffix(imgStr[5:coI], ";base64") {
	case "image/png":
		pngI, err := png.Decode(res)
		if err == nil {
			fileSavePath := basePath + "/" + sha1 + ".png"
			f, _ := os.OpenFile(fileSavePath, os.O_WRONLY|os.O_CREATE, 0777)
			png.Encode(f, pngI)
			fmt.Println("[save base64] Created image: " + fileSavePath)
			f.Close()
		}
		return sha1 + ".png"
	case "image/jpeg":
		jpgI, err := jpeg.Decode(res)
		if err == nil {
			fileSavePath := basePath + "/" + sha1 + ".jpg"
			f, _ := os.OpenFile(fileSavePath, os.O_WRONLY|os.O_CREATE, 0777)
			jpeg.Encode(f, jpgI, &jpeg.Options{Quality: 100})
			fmt.Println("[save base64] Created image: " + fileSavePath)
			f.Close()
		}
		return sha1 + ".jpg"
	}

	return "#"
}
