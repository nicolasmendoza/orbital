package rss

import (
	"fmt"
	"time"
)

func StartBeat(){
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for t := range ticker.C {
			readFeeds() // Reading Feeds... Here we go!! :-)
			fmt.Println("Hearbeat-->", t)
		}
	}()

}