package main

import (
	"fmt"
	"log"
	"net/http"
	"orbita/cache"
	"orbita/cronjob"
)

func main() {
	cache.FlushAll()
	cronjob.Start() // Starts hearbeat...
	http.HandleFunc("/", handleIndex)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Esaaa!!")
	fmt.Println(r)
}
