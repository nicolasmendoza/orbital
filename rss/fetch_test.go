package rss

import (
	"github.com/bradfitz/gomemcache/memcache"
	"net/http"
	"testing"
)

func TestSetCacheHeaders(t *testing.T) {
	url := "https://fakedomain2.com"
	fakeReq, _ := http.NewRequest(http.MethodGet, url, nil)

	fakeReq.Header.Set("If-Modified-Since", "Tue, 06 Feb 2018 17:34:11 GMT")
	fakeReq.Header.Set("If-None-Match", "bfc13a64729c4290ef5b2c2730249c88ca92d82d")

	setCacheHeaders(url, fakeReq.Header)

	// Get cache headers
	cacheHeaders, err := getCacheHeaders(url)
	if err != nil {
		t.Fatalf("Error getting cache Headers...")
	}
	if cacheHeaders["If-Modified-Since"] != "Tue, 06 Feb 2018 17:34:11 GMT" {
		t.Fatalf("Error. If-Modified-Since header doesn't coincide")
	}

	if cacheHeaders["If-None-Match"] != "bfc13a64729c4290ef5b2c2730249c88ca92d82d" {
		t.Fatalf("Error. If-None_match header cached doesn't coincide")
	}

	t.Logf("Cache Header. If-None-Match: %v", "bfc13a64729c4290ef5b2c2730249c88ca92d82d")
	t.Logf("Cache Header. If-Modified-Sice: %v", "Tue, 06 Feb 2018 17:34:11 GMT")
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
	url := "https://www.contratos.gov.co/Archivos/RSSFolder/RSSFiles/rssFeed-10000000.xml"
	resp, err := conditionalGet(url, true)

	if err != nil{
		t.Fatalf("Error doing conditional get to: %s. Error: %v", url, err.Error())
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
