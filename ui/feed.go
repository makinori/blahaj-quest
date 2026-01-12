// feed.go - RSS feed for out-of-stock stores
//
// Purpose:
//   1. See all out-of-stock stores at a glance
//   2. Get notified whenever a new store runs out of stock
//
// Each store has a stable GUID (country + name + coordinates), so RSS readers
// only notify on NEW out-of-stock events. Restocked stores disappear silently.
// If a store goes out of stock again later, it triggers a new notification.

package ui

import (
	"encoding/xml"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/makinori/blahaj-quest/config"
	"github.com/makinori/blahaj-quest/data"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type rssChannel struct {
	XMLName       xml.Name  `xml:"channel"`
	Title         string    `xml:"title"`
	Link          string    `xml:"link"`
	Description   string    `xml:"description"`
	LastBuildDate string    `xml:"lastBuildDate"`
	Items         []rssItem `xml:"item"`
}

type rssItem struct {
	Title       string  `xml:"title"`
	Link        string  `xml:"link"`
	GUID        rssGUID `xml:"guid"`
	Description string  `xml:"description"`
}

type rssGUID struct {
	IsPermaLink bool   `xml:"isPermaLink,attr"`
	Value       string `xml:",chardata"`
}

type rssFeed struct {
	XMLName xml.Name   `xml:"rss"`
	Version string     `xml:"version,attr"`
	Channel rssChannel `xml:"channel"`
}

// slugify converts a store name to a URL-safe slug for use in GUIDs
func slugify(s string) string {
	// Normalize unicode (decompose accents)
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, s)

	// Lowercase
	result = strings.ToLower(result)

	// Replace spaces and common separators with hyphens
	result = strings.ReplaceAll(result, " ", "-")
	result = strings.ReplaceAll(result, "_", "-")

	// Remove any character that isn't alphanumeric or hyphen
	reg := regexp.MustCompile(`[^a-z0-9-]`)
	result = reg.ReplaceAllString(result, "")

	// Collapse multiple hyphens
	reg = regexp.MustCompile(`-+`)
	result = reg.ReplaceAllString(result, "-")

	// Trim leading/trailing hyphens
	result = strings.Trim(result, "-")

	return result
}

// RenderFeed generates an RSS feed of all out-of-stock stores
func RenderFeed() (string, error) {
	var items []rssItem

	for _, store := range data.Blahaj.Current {
		if store.Quantity > 0 {
			continue
		}

		storeSlug := slugify(store.Name)
		// Include lat/lng to ensure uniqueness for stores with same name
		guid := "blahaj|" + store.CountryCode + "|" + storeSlug + "|" + store.Lat + "," + store.Lng + "|oos"

		// Build Google Maps link
		mapsLink := "https://www.google.com/maps?q=" + store.Lat + "," + store.Lng

		// Build IKEA store page link
		ikeaLink := "https://www.ikea.com/" + store.CountryCode + "/" + store.LanguageCode + "/stores/"

		// Rich HTML description
		description := `<p><strong>Blahaj is currently out of stock</strong> at IKEA ` + store.Name + `.</p>` +
			`<p>` +
			`<a href="` + mapsLink + `">View on Google Maps</a> · ` +
			`<a href="` + ikeaLink + `">IKEA Store Page</a>` +
			`</p>`

		items = append(items, rssItem{
			Title:       "IKEA " + store.Name + " (" + strings.ToUpper(store.CountryCode) + ") – Blahaj out of stock",
			Link:        mapsLink,
			Description: description,
			GUID: rssGUID{
				IsPermaLink: false,
				Value:       guid,
			},
		})
	}

	feed := rssFeed{
		Version: "2.0",
		Channel: rssChannel{
			Title:         "Blahaj Quest – Out of Stock Alerts",
			Link:          config.URL,
			Description:   "Get notified when IKEA stores run out of Blahaj",
			LastBuildDate: time.Now().Format(time.RFC1123Z),
			Items:         items,
		},
	}

	output, err := xml.MarshalIndent(feed, "", "  ")
	if err != nil {
		return "", err
	}

	return xml.Header + string(output), nil
}
