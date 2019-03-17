package http

import (
	"net/http"
)

// Do a GET HTTP Request. Using Context, and URL Fetch.
func Get(r *http.Request, url string) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
