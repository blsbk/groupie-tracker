package main

import (
	"fmt"
	groupie "groupie/functions"
	"log"
	"net/http"
)

func main() {
	err := groupie.InitArtists()
	if err != nil {
		log.Fatalln(err)
	}

	styles := http.FileServer(http.Dir("./templates/styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", styles))

	http.HandleFunc("/", groupie.MainPage)
	http.HandleFunc("/artist/", groupie.OneArtist)

	fmt.Printf("Starting server got testing... http://127.0.0.1:8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
