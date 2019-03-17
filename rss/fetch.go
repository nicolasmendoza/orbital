package rss

import (
	"fmt"
	"log"
	"net/http"
	"orbita/cache"
	"strings"
	"time"
)

func getCacheHeaders(key string) (headers map[string]string, err error) {
	it, err := cache.Get(key)
	if err != nil {
		return nil, err
	}
	h := strings.Split(string(it.Value), "|")
	headers = make(map[string]string)
	headers["If-None-Match"] = h[0]
	headers["If-Modified-Since"] = h[1]

	return headers, nil
}

// Store Header values in cache...
func setCacheHeaders(key string, h http.Header) {
	//parsing time. Lastmod value need special parsing...
	// "Tue, 06 Feb 2018 17:34:11 GMT"
	t, err := time.Parse(time.RFC1123, h.Get("If-Modified-Since"))
	if err != nil {
		log.Fatalf("Error parsing If-Modified-Since Header. %v", err.Error())
	}

	// https://goplay.space/#I6F6AV_0F-s
	lastMod := t.Format(time.RFC1123)
	etag := h.Get("If-None-Match")

	value := fmt.Sprintf("%s|%s", etag, lastMod)

	if err := cache.Set(key, value); err != nil {
		log.Fatalf("Error storing Header in cache: %v", err.Error())
	}

}
func conditionalGet(url string, checkCache bool) (resp *http.Response, err error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if checkCache == true {
		cacheHeaders, _ := getCacheHeaders(url)
		if cacheHeaders != nil {
			req.Header.Set("If-None-Match", fmt.Sprintf("%s",cacheHeaders["If-None-Match"]))
			req.Header.Set("If-Modified-Since", cacheHeaders["If-Modified-Since"])
		}
	}

	// Do request...
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	return resp, nil
}
