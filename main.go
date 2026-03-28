package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// //html := `<strong>Hello World ok</strong>`

	// http.ServeFile(w, r, "index.html")
	// w.Header().Set("Content-Type", "text/html")
	// //fmt.Fprintf(w, html)

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

func main() {
	var port = 8080

	http.HandleFunc("/", homePage)

	log.Println("Starting web server on port", port)
	http.ListenAndServe(":8080", nil)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(" server error: ", err)
		return
	}
}
