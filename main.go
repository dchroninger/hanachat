package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	// Limit endpoint to only GET requests.
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the HTML template to send back to the client.
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	tmpl.Execute(w, nil)
}

func main() {
	flag.Parse()

	http.HandleFunc("/", serveHome)

	server := &http.Server{
		Addr:              *addr,
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
