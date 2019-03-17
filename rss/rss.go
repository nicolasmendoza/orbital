package rss

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
)

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
	PubDate  string    `xml:"channel>pubDate"`
	ItemList []XMLItem `xml:"channel>item"`
}

// http://www.w3schools.com/rss/rss_item.asp
// http://stackoverflow.com/questions/7220670/difference-between-description-and-contentencoded-tags-in-rss2
// https://groups.google.com/d/topic/golang-nuts/uBMo1BpaQCM
type XMLItem struct {
	// Required
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Description template.HTML `xml:"description"`
	// Optional
	Content  template.HTML `xml:"encoded"`
	PubDate  string        `xml:"pubDate"`
	Comments string        `xml:"comments"`
}

type Document struct {
	Title       string
	Description template.HTML
	Link        string
	//Author      string
	//Category    string
	PubDate string
	Done    bool
}

func getDocument(url string, checkCache bool) error {
	body, err := readBody(url, checkCache)
	if err != nil {
		log.Println(err.Error())
		return err

	}

	docXML := parserXML(body)
	if docXML != nil {
		//Iterating items in document.
		for _, value := range docXML.ItemList {

			document := new(Document)
			document.Title = value.Link
			document.Link = value.Link
			document.PubDate = value.PubDate
			document.Done = false

		}
	}

	println(body)
	return nil
}

func parserXML(doc []byte) *RSS {
	//Processing XML
	var docXML RSS
	if err := xml.Unmarshal(doc, &docXML); err != nil { // TODO FIXME. STORE THE XML FILE when error ocurred.
		log.Fatal("Cannot marshall XML (RSS) File...")
		return nil
	}

	return &docXML
}

func readBody(url string, checkCache bool) ([]byte, error) {
	// HTTP GET to the resource, with conditional value.
	resp, err := conditionalGet(url, checkCache)
	if err != nil {
		return nil, fmt.Errorf("cannot get URL: %s, %v", url, err.Error())
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error getting reading body for url: %s. Error: %v", url, err.Error())
	}
	return body, nil

}
