/*
Execute a task continually every X time.
*/
package heartbeat

import (
	"fmt"
	"orbita/rss"
	"time"
)

const everyTime= 1 * time.Hour

func Start() {

	rss.ReadFeeds() // execute the first time...later wait for the newTicker execution.

	ticker := time.NewTicker(everyTime)
	go func() {
		for t := range ticker.C {
			rss.ReadFeeds() // Reading Feeds... Here we go!! :-)
			fmt.Println("Hearbeat -->", t)
		}
	}()

}
