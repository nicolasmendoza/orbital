package rss

import (
	"fmt"
	"io/ioutil"
)



func read(url string, checkCache bool)([]byte, error){
	// HTTP GET to the resource, with conditional value.
	resp, err := conditionalGet(url, checkCache)
	if err!=nil{
		return nil, fmt.Errorf("cannot get URL: %s, %v", url, err.Error())
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		return nil, fmt.Errorf("error getting reading body for url: %s. Error: %v", url, err.Error())
	}
	return body, nil


}

/*func getDocument(url string){
	body, err := read(url)
	if err!=nil{
		log.Println(err.Error())
	}

}*/