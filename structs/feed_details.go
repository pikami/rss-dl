package structs

// FeedDetails - Struct for outputing feed details summary
type FeedDetails struct {
	Title       string
	Description string
	Link        string
	Updated     string
	Language    string
	Categories  []string
}
