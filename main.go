package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	muxer := http.NewServeMux()
	fileServerHtml := http.FileServer(http.Dir("bin"))
	muxer.Handle("/", fileServerHtml)
	log.Println("Server starting at port " + port + "...")
	if err := http.ListenAndServe(":"+port, muxer); err != nil {
		log.Fatal(err)
	}
}
