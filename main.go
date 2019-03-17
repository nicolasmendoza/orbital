package main

import (
	"fmt"
	"log"
	"net/http"
	"orbita/rss"
)

func main() {
	rss.StartBeat()
	http.HandleFunc("/", handleIndex)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Esaaa!!")
}
