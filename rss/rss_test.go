package rss

import (
	"io/ioutil"
	"testing"
)

const (
	url = "https://www.contratos.gov.co/Archivos/RSSFolder/RSSFiles/rssFeed-10000000.xml"
)

func TestReadBody(t *testing.T) {
	body, err := readBody(url, false)
	if err != nil {
		t.Errorf("cannot get body from url %s . %v ", url, err.Error())
	}
	if string(body[0:4]) == "<?xml" {
		t.Errorf("cannot get body from url: %s . It's not a valid XML file.", url)
		t.Log(string(body))
	}
}

func TestGetDocument(t *testing.T) {
	err := getDocument(url, false)
	if err != nil {
		t.Errorf("Error getting XML Document: %v", err.Error())
	}
}

func TestParserXMLtoDocument(t *testing.T) {
	// Openning Example XML...
	xmlFile, err := ioutil.ReadFile("./rss_test.xml")
	if err != nil {
		t.Fatalf("Cannot open test file...%v. %v", xmlFile, err.Error())
	}

	rss := parserXML(xmlFile)
	if rss == nil{
		t.Error("Error parsing HTML")
	}
	if rss.Version != "2.0"{
		t.Error("Invalid XML version. Error parsing file...")
	}
}

func TestCreateDocument(t *testing.T){
	i := new(XMLItem)  // #TODO, #IMPROVEMENT. Create TDDD
	i.Link = "https://aserejejadeje.com"
	i.Description = "Nueva licitación en santa marta"
	i.Title = "Nueva licitacion"
	i.PubDate = "None"

	newDoc := createDocument(i)

	if newDoc.Title != i.Title{
		t.Errorf("Error matching %v VS %v", newDoc.Title, i.Title)
	}

	if newDoc.Link != i.Link{
		t.Errorf("Error matching %v VS %v", newDoc.Link, i.Link)
	}

	if newDoc.Done != false{
		t.Errorf("Error default values expected is: false, but gets: %v", newDoc.Done)
	}
}