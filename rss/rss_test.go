package rss

import "testing"

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

func TestGetDocument(t *testing.T){
	err := getDocument(url, false)
	if err !=nil{
		t.Errorf("Error getting XML Document: %v", err.Error())
	}
}
