package rss

import (
	"github.com/bradfitz/gomemcache/memcache"
	"net/http"
	"testing"
)

func TestSetCacheHeaders(t *testing.T) {
	url := "https://fakedomain2.com"
	fakeReq, _ := http.NewRequest(http.MethodGet, url, nil)

	fakeReq.Header.Set("Last-Modified", "Sat, 16 Mar 2019 04:50:07 GMT")
	fakeReq.Header.Set("ETag", "bfc13a64729c4290ef5b2c2730249c88ca92d82d")

	setCacheHeaders(url, fakeReq.Header)

	// Get cache headers
	cacheHeaders, err := getCacheHeaders(url)
	if err != nil {
		t.Errorf("Error getting cache Headers... %v", err.Error())
	}
	if cacheHeaders[headerLastModified] != "Sat, 16 Mar 2019 04:50:07 GMT" {
		t.Error("Error. If-Modified-Since header doesn't coincide")
	}

	if cacheHeaders[headerETag] != "bfc13a64729c4290ef5b2c2730249c88ca92d82d" {
		t.Error("Error. If-None_match header cached doesn't coincide")
	}

}

func TestGetCacheHeaders(t *testing.T) {
	key := "https://fakedomain.com"
	// Unexisting cache...
	_, err := getCacheHeaders(key)

	if err != memcache.ErrCacheMiss {
		t.Logf("Error getting cache for key:%v. Error:", key)
	}
}

func TestConditionalGet(t *testing.T){
	url := "https://www.contratos.gov.co/Archivos/RSSFolder/RSSFiles/rssFeed-27000000.xml"
	resp, modified, err := conditionalGet(url, false)

	if err != nil{
		t.Errorf("Error doing conditional get to: %s. Error: %v", url, err.Error())
	}
	if modified != true{
		t.Errorf("Expected Modified: true but we received: %v", modified)
	}

	//getting again...will be 304 not modified
	resp, modified, _= conditionalGet(url, true)
	if resp.StatusCode != http.StatusNotModified{
		t.Logf("getConditional is not caching the Headers. Response: %v, url: %s", resp.StatusCode, url)
	}

	if modified != false{
		t.Logf("Expected False. Received %v", modified)
	}
	t.Logf("Status: %v", resp.Header)
	t.Logf("Status: %v", resp.StatusCode)

}

/*

func main() {

	// "Mon, 01/02/06, 03:04PM"
	layout := time.RFC1123
	str := "Tue, 06 Feb 2018 17:34:11 GMT"

	t, err := time.Parse(layout, str)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)

	fmt.Println(t.String())
	fmt.Println(t.Format(time.RFC1123))
}

*/
