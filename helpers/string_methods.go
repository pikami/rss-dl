package helpers

import (
	"log"
	"regexp"
)

// ToCleanString - replaces spaces with underscores
func ToCleanString(str string) string {
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(str, "_")
}
