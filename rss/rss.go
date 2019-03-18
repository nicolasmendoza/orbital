package rss

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	db2 "orbita/db"
	"time"
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

// Store document into Database...
func (d *Document) Insert() {
	db := db2.Get()
	query := "INSERT INTO curiosity_documents VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	stmInsert, err := db.Prepare(query)
	if err != nil {
		log.Panicf("Error preparing Query %v", err.Error())
	}
	defer stmInsert.Close() // DANGER!!!!
	var datetime = time.Now()
	datetime.Format(time.RFC3339)
	if _, err := stmInsert.Exec(nil, datetime, "category", d.PubDate, d.Title, d.Link, d.Description, d.Done); err != nil {
		log.Panicf("Error inserting document in database %v", err.Error())
	}
}

func newDocument(v *XMLItem) *Document {
	document := new(Document)
	document.Title = v.Title
	document.Link = v.Link
	document.PubDate = v.PubDate
	document.Done = false
	return document
}

func getDocument(url string, checkCache bool) error {
	body, modified, err := readBody(url, checkCache)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	if modified == true {
		// TODO add Here LOGGER...
		docXML := parserXML(body)
		if docXML != nil {
			//Iterating items in document. (building document)
			for _, xmlItem := range docXML.ItemList {
				doc := newDocument(&xmlItem)
				doc.Insert()
			}
		}

	}
	return nil
}

func readBody(url string, checkCache bool) ([]byte, bool, error) {
	// HTTP GET to the resource, with conditional value.
	resp, modified, err := conditionalGet(url, checkCache)
	if err != nil {
		return nil, false, fmt.Errorf("cannot get URL: %s, %v", url, err.Error())
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	// If document not was modified...so we don't need process this file..
	if modified == false {
		return nil, false, nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, false, fmt.Errorf("error getting reading body for url: %s. Error: %v", url, err.Error())
	}
	return body, true, nil

}

func parserXML(doc []byte) *RSS {
	//Processing XML
	var docXML RSS
	log.Println("-->", string(doc))
	if err := xml.Unmarshal(doc, &docXML); err != nil { // #TODO #FIXME. STORE THE XML FILE when error ocurred.
		log.Fatalf("Cannot marshall XML (RSS) File...%v", err.Error())
		return nil
	}

	return &docXML
}
