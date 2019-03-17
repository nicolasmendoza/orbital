package rss

import "testing"

func TestReadURL(t *testing.T) {
	url := "https://www.contratos.gov.co/Archivos/RSSFolder/RSSFiles/rssFeed-10000000.xml"
	body, err := read(url, false)
	if err != nil {
		t.Errorf("cannot get body from url %s . %v ", url, err.Error())
	}
	if string(body[0:4]) == "<?xml" {
		t.Errorf("cannot get body from url: %s . It's not a valid XML file.", url)
		t.Log(string(body))
	}
}
