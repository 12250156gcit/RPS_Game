package main

import (
	"fmt"
	"log"
	"net/http"
	"rockpaperscs/rps"
	"text/template"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		log.Panicln(err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	winner, cp, rr := rps.PlayRound(1) // player chose paper
	fmt.Println("winner:", winner)
	fmt.Println("computer choice:", cp)
	fmt.Println("round result:", rr)
}

func main() {
	var port = 8080

	// http.HandleFunc("/", homePage)
	http.HandleFunc("/", playHandler)

	log.Println("Starting web server on port", port)
	http.ListenAndServe(":8080", nil)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(" server error: ", err)
		return
	}
}
