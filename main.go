package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rockpaperscs/rps"
	"text/template"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		log.Println(err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func playHandler(w http.ResponseWriter, r *http.Request) {

	c := r.URL.Query().Get("c")

	var playerValue int
	fmt.Sscanf(c, "%d", &playerValue)

	result := rps.PlayRound(playerValue)

	jsObj, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsObj)
}

func main() {
	var port = 8080

	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", playHandler)

	log.Println("Starting web server on port", port)
	http.ListenAndServe(":8080", nil)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(" server error: ", err)
		return
	}
}

// page 33
