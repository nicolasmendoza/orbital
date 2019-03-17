package rss

import (
	"encoding/xml"
	"html/template"
)

// http://www.w3schools.com/rss/rss_item.asp
// http://stackoverflow.com/questions/7220670/difference-between-description-and-contentencoded-tags-in-rss2
// https://groups.google.com/d/topic/golang-nuts/uBMo1BpaQCM
type Item struct {
	// Required
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Description template.HTML `xml:"description"`
	// Optional
	Content  template.HTML `xml:"encoded"`
	PubDate  string        `xml:"pubDate"`
	Comments string        `xml:"comments"`
}


// http://www.w3schools.com/rss/rss_syntax.asp
// http://www.w3schools.com/rss/rss_channel.asp
// RSS Version 2.
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	// Required
	Title       string `xml:"channel>title"`
	Link        string `xml:"channel>link"`
	Description string `xml:"channel>description"`
	// Optional
	PubDate  string `xml:"channel>pubDate"`
	ItemList []Item `xml:"channel>item"`
}
