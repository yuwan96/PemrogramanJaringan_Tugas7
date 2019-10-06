package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func StartNonTLSServer() {
	mux := new(http.ServeMux)
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Redirecting to https://localhost/")
		http.Redirect(w, r, "https://localhost/", http.StatusTemporaryRedirect)
	}))

	http.ListenAndServe(":8080", mux)
}
func main() {
	go StartNonTLSServer()

	mux := new(http.ServeMux)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filepath = path.Join("static", "coba.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var data = map[string]interface{}{
			"title": "Join Kuy",
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	log.Println("Server started at :443")
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", mux)
	if err != nil {
		panic(err)
	}
}
