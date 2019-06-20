package helpers

import (
	"log"
	"regexp"
	"strings"
)

// ToCleanString - replaces spaces with underscores
func ToCleanString(str string) string {
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(str, "_")
}

// RemoveGetParams - removes http GET params
func RemoveGetParams(str string) string {
	return strings.Split(str, "?")[0]
}
