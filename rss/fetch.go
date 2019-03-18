package rss

import (
	"fmt"
	"log"
	"net/http"
	"orbita/cache"
	"strings"
	"time"
)

const (
	headerLastModified = "If-Modified-Since"
	headerETag         = "If-None-Match"
)

func getCacheHeaders(key string) (map[string]string, error) {
	it, err := cache.Get(key)
	if err != nil {
		return nil, err
	}
	h := strings.Split(string(it.Value), "|")
	headers := make(map[string]string)
	headers[headerETag] = h[0]
	headers[headerLastModified] = h[1]

	return headers, nil
}

// Store Header values in cache...
func setCacheHeaders(key string, h http.Header) {
	//parsing time. Lastmod value need special parsing...
	// "Tue, 06 Feb 2018 17:34:11 GMT"
	t, err := time.Parse(time.RFC1123, h.Get("Last-Modified"))
	if err != nil {
		panic(err.Error())
	}

	// https://goplay.space/#I6F6AV_0F-s
	lastMod := t.Format(time.RFC1123)
	etag := h.Get("ETag")

	value := fmt.Sprintf("%s|%s", etag, lastMod)

	if err := cache.Set(key, value); err != nil {
		log.Fatalf("Error storing Header in cache: %v", err.Error())
	}

}
func conditionalGet(url string, checkCache bool) (*http.Response, bool, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic("Error creating New Request %v")
	}

	if checkCache == true {
		cacheHeaders, _ := getCacheHeaders(url)
		if cacheHeaders != nil {
			req.Header.Set("If-None-Match", cacheHeaders[headerETag])
			req.Header.Set("If-Modified-Since", cacheHeaders[headerLastModified])
		}
	}

	// Do request...
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error doing request to %s, %v", url, err.Error())
	}

	var modified bool

	if resp.StatusCode == http.StatusOK {
		setCacheHeaders(url, resp.Header)
		modified = true
	}
	if (resp.StatusCode == http.StatusNotModified) {
		modified = false
	}
	return resp, modified, nil
}
